package cmd

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	pb "github.com/kamilsk/guard/pkg/transport/grpc"

	"github.com/ghodss/yaml"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kamilsk/go-kit/pkg/strings"
	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/transport/grpc/middleware"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const nl = 10 // \n

var entities factory

func init() {
	entities = factory{
		registerLicense: func() interface{} { return &pb.RegisterLicenseRequest{} },
		createLicense:   func() interface{} { return &pb.CreateLicenseRequest{} },
		readLicense:     func() interface{} { return &pb.ReadLicenseRequest{} },
		updateLicense:   func() interface{} { return &pb.UpdateLicenseRequest{} },
		deleteLicense:   func() interface{} { return &pb.DeleteLicenseRequest{} },
		restoreLicense:  func() interface{} { return &pb.RestoreLicenseRequest{} },
	}
}

func communicate(cmd *cobra.Command, _ []string) error {
	request, resolveErr := entities.request(cmd)
	if resolveErr != nil {
		return resolveErr
	}

	encoder := &runtime.JSONPb{OrigName: true}
	writer := func(w io.Writer) writerFunc {
		return func(p []byte) error {
			_, err := w.Write(p)
			return err
		}
	}(cmd.OutOrStdout())

	if dry, _ := cmd.Flags().GetBool("dry-run"); dry {
		cmd.Printf("%T would be sent with data: ", request)
		output, err := encoder.Marshal(request)
		if err != nil {
			return err
		}
		if cmd.Flag("output").Value.String() == jsonFormat {
			return writer.Write(append(output, nl))
		}
		output, err = yaml.JSONToYAML(output)
		if err != nil {
			return err
		}
		return writer.Write(append([]byte{nl}, output...))
	}

	response, err := call(cnf.Union.GRPCConfig, request)
	if err != nil {
		return writer.Write(append([]byte(err.Error()), nl))
	}
	output, err := encoder.Marshal(response)
	if cmd.Flag("output").Value.String() == jsonFormat {
		return writer.Write(append(output, nl))
	}
	output, err = yaml.JSONToYAML(output)
	if err != nil {
		return err
	}
	return writer.Write(output)
}

type writerFunc func([]byte) error

func (fn writerFunc) Write(p []byte) error {
	return fn(p)
}

type builder func() interface{}

type factory map[*cobra.Command]builder

func (f factory) request(cmd *cobra.Command) (interface{}, error) {
	data, err := f.data(cmd.Flag("filename").Value.String())
	if err != nil {
		return nil, err
	}

	var raw json.RawMessage
	if decodeErr := yaml.Unmarshal(data, &raw); decodeErr != nil {
		return nil, errors.Wrap(decodeErr, "trying to decode YAML")
	}

	encoder, request := &runtime.JSONPb{OrigName: true}, f[cmd]()
	encoder.Unmarshal(raw, request)
	return request, nil
}

func (factory) data(filename string) ([]byte, error) {
	var (
		err error
		raw []byte
		src io.Reader = os.Stdin
	)
	if filename != "" {
		if src, err = os.Open(filename); err != nil {
			return nil, errors.Wrapf(err, "opening the file %q", filename)
		}
	} else {
		filename = "/dev/stdin"
	}
	if raw, err = ioutil.ReadAll(src); err != nil {
		return nil, errors.Wrapf(err, "reading the file %q", filename)
	}
	return raw, nil
}

func call(cnf config.GRPCConfig, request interface{}) (interface{}, error) {
	deadline, cancel := context.WithTimeout(context.Background(), cnf.Timeout)
	conn, err := grpc.DialContext(deadline, cnf.Interface, grpc.WithInsecure())
	cancel()
	if err != nil {
		return nil, errors.Wrapf(err, "connecting to the gRPC server at %q", cnf.Interface)
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx,
		middleware.AuthHeader,
		strings.Concat(middleware.AuthScheme, " ", string(cnf.Token)))

	client := pb.NewLicenseClient(conn)
	switch in := request.(type) {
	case *pb.CreateLicenseRequest:
		return client.Create(ctx, in)
	case *pb.ReadLicenseRequest:
		return client.Read(ctx, in)
	case *pb.UpdateLicenseRequest:
		return client.Update(ctx, in)
	case *pb.DeleteLicenseRequest:
		return client.Delete(ctx, in)
	case *pb.RestoreLicenseRequest:
		return client.Restore(ctx, in)
	case *pb.RegisterLicenseRequest:
		return client.Register(ctx, in)
	default:
		return nil, errors.Errorf("unknown request type %T", request)
	}
}
