package cmd

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.octolab.org/strings"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/transport/grpc/middleware"
	"github.com/kamilsk/guard/pkg/transport/grpc/protobuf"
)

const nl = 10 // \n

var entities factory

func init() {
	entities = factory{
		Install:         func() interface{} { return &protobuf.InstallRequest{} },
		registerLicense: func() interface{} { return &protobuf.RegisterLicenseRequest{} },
		createLicense:   func() interface{} { return &protobuf.CreateLicenseRequest{} },
		readLicense:     func() interface{} { return &protobuf.ReadLicenseRequest{} },
		updateLicense:   func() interface{} { return &protobuf.UpdateLicenseRequest{} },
		deleteLicense:   func() interface{} { return &protobuf.DeleteLicenseRequest{} },
		restoreLicense:  func() interface{} { return &protobuf.RestoreLicenseRequest{} },

		// TODO issue#draft {

		addEmployee:     func() interface{} { return &protobuf.AddEmployeeRequest{} },
		deleteEmployee:  func() interface{} { return &protobuf.DeleteEmployeeRequest{} },
		addWorkplace:    func() interface{} { return &protobuf.AddWorkplaceRequest{} },
		deleteWorkplace: func() interface{} { return &protobuf.DeleteWorkplaceRequest{} },
		pushWorkplace:   func() interface{} { return &protobuf.PushWorkplaceRequest{} },

		// issue#draft }
	}
}

func communicate(cmd *cobra.Command, _ []string) error {
	request, resolveErr := entities.request(cmd)
	if resolveErr != nil {
		return resolveErr
	}

	encoder := &runtime.JSONPb{OrigName: true}
	wrap := func(w io.Writer) writerFunc {
		return func(p []byte) error {
			_, err := w.Write(p)
			return err
		}
	}

	if dry, _ := cmd.Flags().GetBool("dry-run"); dry {
		cmd.Printf("%T would be sent with data: ", request)
		output, err := encoder.Marshal(request)
		if err != nil {
			return err
		}
		if cmd.Flag("output").Value.String() == jsonFormat {
			return wrap(cmd.OutOrStderr()).Write(append(output, nl))
		}
		output, err = yaml.JSONToYAML(output)
		if err != nil {
			return err
		}
		return wrap(cmd.OutOrStderr()).Write(append([]byte{nl}, output...))
	}

	response, err := call(cnf.Union.GRPCConfig, request)
	if err != nil {
		return wrap(cmd.OutOrStderr()).Write(append([]byte(err.Error()), nl))
	}
	output, err := encoder.Marshal(response)
	if err != nil {
		return err
	}
	if cmd.Flag("output").Value.String() == jsonFormat {
		return wrap(cmd.OutOrStdout()).Write(append(output, nl))
	}
	output, err = yaml.JSONToYAML(output)
	if err != nil {
		return err
	}
	return wrap(cmd.OutOrStdout()).Write(output)
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
	deadline, cancel := context.WithTimeout(context.Background(), cnf.RPC.Timeout)
	conn, err := grpc.DialContext(deadline, cnf.RPC.Interface, grpc.WithInsecure())
	cancel()
	if err != nil {
		return nil, errors.Wrapf(err, "connecting to the gRPC server at %q", cnf.RPC.Interface)
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx,
		middleware.AuthHeader,
		strings.Concat(middleware.AuthScheme, " ", string(cnf.RPC.Token)))

	switch in := request.(type) {
	case *protobuf.InstallRequest:
		return protobuf.NewMaintenanceClient(conn).Install(ctx, in)
	case *protobuf.CreateLicenseRequest:
		return protobuf.NewLicenseClient(conn).Create(ctx, in)
	case *protobuf.ReadLicenseRequest:
		return protobuf.NewLicenseClient(conn).Read(ctx, in)
	case *protobuf.UpdateLicenseRequest:
		return protobuf.NewLicenseClient(conn).Update(ctx, in)
	case *protobuf.DeleteLicenseRequest:
		return protobuf.NewLicenseClient(conn).Delete(ctx, in)
	case *protobuf.RestoreLicenseRequest:
		return protobuf.NewLicenseClient(conn).Restore(ctx, in)
	case *protobuf.RegisterLicenseRequest:
		return protobuf.NewLicenseClient(conn).Register(ctx, in)

	// TODO issue#draft {

	case *protobuf.AddEmployeeRequest:
		return protobuf.NewLicenseClient(conn).AddEmployee(ctx, in)
	case *protobuf.DeleteEmployeeRequest:
		return protobuf.NewLicenseClient(conn).DeleteEmployee(ctx, in)
	case *protobuf.AddWorkplaceRequest:
		return protobuf.NewLicenseClient(conn).AddWorkplace(ctx, in)
	case *protobuf.DeleteWorkplaceRequest:
		return protobuf.NewLicenseClient(conn).DeleteWorkplace(ctx, in)
	case *protobuf.PushWorkplaceRequest:
		return protobuf.NewLicenseClient(conn).PushWorkplace(ctx, in)

	// issue#draft }

	default:
		return nil, errors.Errorf("unknown request type %T", request)
	}
}
