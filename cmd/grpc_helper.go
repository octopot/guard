package cmd

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	pb "github.com/kamilsk/guard/pkg/transport/grpc"

	"github.com/kamilsk/go-kit/pkg/strings"
	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/transport/grpc/middleware"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gopkg.in/yaml.v2"
)

var entities factory

func init() {
	entities = factory{
		createLicense: {"License": func() pb.Proxy { return pb.CreateLicenseRequestProxy{} }},
		readLicense:   {"License": func() pb.Proxy { return pb.ReadLicenseRequestProxy{} }},
		updateLicense: {"License": func() pb.Proxy { return pb.UpdateLicenseRequestProxy{} }},
		deleteLicense: {"License": func() pb.Proxy { return pb.DeleteLicenseRequestProxy{} }},

		// ---

		registerLicense: {"License": func() pb.Proxy { return pb.RegisterLicenseRequestProxy{} }},
	}
}

func communicate(cmd *cobra.Command, _ []string) error {
	entity, err := entities.proxy(cmd)
	if err != nil {
		return err
	}
	if dry, _ := cmd.Flags().GetBool("dry-run"); dry {
		cmd.Printf("%T would be sent with data: ", entity)
		if cmd.Flag("output").Value.String() == jsonFormat {
			return json.NewEncoder(cmd.OutOrStdout()).Encode(entity)
		}
		return yaml.NewEncoder(cmd.OutOrStdout()).Encode(entity)
	}
	response, err := call(cnf.Union.GRPCConfig, entity)
	if err != nil {
		return err
	}
	if cmd.Flag("output").Value.String() == jsonFormat {
		return json.NewEncoder(cmd.OutOrStdout()).Encode(response)
	}
	return yaml.NewEncoder(cmd.OutOrStdout()).Encode(response)
}

type builder func() pb.Proxy
type kind string
type field string
type value interface{}

type schema struct {
	Kind    kind            `yaml:"kind"`
	Payload map[field]value `yaml:"payload"`
}

type factory map[*cobra.Command]map[kind]builder

func (factory) scheme(filename string) (schema, error) {
	var (
		err error
		out schema
		raw []byte
		src io.Reader = os.Stdin
	)
	if filename != "" {
		if src, err = os.Open(filename); err != nil {
			return out, errors.Wrapf(err, "trying to open file %q", filename)
		}
	} else {
		filename = "/dev/stdin"
	}
	if raw, err = ioutil.ReadAll(src); err != nil {
		return out, errors.Wrapf(err, "trying to read file %q", filename)
	}
	err = yaml.Unmarshal(raw, &out)
	return out, errors.Wrapf(err, "trying to decode file %q as YAML", filename)
}

func (f factory) proxy(cmd *cobra.Command) (pb.Proxy, error) {
	scheme, err := f.scheme(cmd.Flag("filename").Value.String())
	if err != nil {
		return nil, err
	}
	build, ok := f[cmd][scheme.Kind]
	if !ok {
		return nil, errors.Errorf("unknown payload type %q", scheme.Kind)
	}
	entity := build()
	if err = mapstructure.Decode(scheme.Payload, &entity); err != nil {
		return nil, errors.Wrapf(err, "trying to decode payload to %#v", entity)
	}
	return entity, nil
}

func call(cnf config.GRPCConfig, entity pb.Proxy) (interface{}, error) {
	deadline, cancel := context.WithTimeout(context.Background(), cnf.Timeout)
	conn, err := grpc.DialContext(deadline, cnf.Interface, grpc.WithInsecure())
	cancel()
	if err != nil {
		return nil, errors.Wrapf(err, "trying to connect to the gRPC server at %q", cnf.Interface)
	}
	defer conn.Close()
	client := pb.NewLicenseClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx,
		middleware.AuthHeader,
		strings.Concat(middleware.AuthScheme, " ", cnf.Token.String()))
	switch request := entity.Convert().(type) {
	case *pb.CreateLicenseRequest:
		return client.Create(ctx, request)
	case *pb.ReadLicenseRequest:
		return client.Read(ctx, request)
	case *pb.UpdateLicenseRequest:
		return client.Update(ctx, request)
	case *pb.DeleteLicenseRequest:
		return client.Delete(ctx, request)
	case *pb.RegisterLicenseRequest:
		return client.Register(ctx, request)
	default:
		return nil, errors.Errorf("unknown request type %T", request)
	}
}
