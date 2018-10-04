package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/kamilsk/go-kit/pkg/strings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	yamlFormat = "yaml"
	jsonFormat = "json"
)

var (
	// Install TODO issue#docs
	Install = &cobra.Command{Use: "install", Short: "Install application", RunE: communicate}
	// License TODO issue#docs
	License         = &cobra.Command{Use: "license", Short: "Guard License Management"}
	registerLicense = &cobra.Command{Use: "register", Short: "Register client license", RunE: communicate}
	createLicense   = &cobra.Command{Use: "create", Short: "Create client license", RunE: communicate}
	readLicense     = &cobra.Command{Use: "read", Short: "Read client license", RunE: communicate}
	updateLicense   = &cobra.Command{Use: "update", Short: "Update client license", RunE: communicate}
	deleteLicense   = &cobra.Command{Use: "delete", Short: "Delete client license", RunE: communicate}
	restoreLicense  = &cobra.Command{Use: "restore", Short: "Restore client license", RunE: communicate}
)

func init() {
	v := viper.New()
	fn.Must(
		func() error { return v.BindEnv("bind") },
		func() error { return v.BindEnv("grpc_port") },
		func() error { return v.BindEnv("guard_token") },
		func() error {
			v.SetDefault("bind", defaults["bind"])
			v.SetDefault("grpc_port", defaults["grpc_port"])
			v.SetDefault("grpc_host", strings.Concat(v.GetString("bind"), ":", strconv.Itoa(v.GetInt("grpc_port"))))
			v.SetDefault("guard_token", defaults["guard_token"])
			return nil
		},
		func() error {
			flags := License.PersistentFlags()

			flags.StringVarP(new(string), "filename", "f", "", "entity source (default is standard input)")
			flags.StringVarP(new(string), "output", "o", yamlFormat, fmt.Sprintf(
				"output format, one of: %s|%s",
				jsonFormat, yamlFormat))
			flags.Bool("dry-run", false, "if true, only print the object that would be sent, without sending it")
			flags.StringVarP(&cnf.Union.GRPCConfig.Interface,
				"grpc-host", "", v.GetString("grpc_host"), "gRPC server host")
			flags.DurationVarP(&cnf.Union.GRPCConfig.Timeout,
				"timeout", "t", time.Second, "connection timeout")

			flags.StringVarP((*string)(&cnf.Union.GRPCConfig.Token),
				"token", "", v.GetString("guard_token"), "user access token")
			return nil
		},
		func() error {
			flags := Install.Flags()

			flags.StringVarP(new(string), "filename", "f", "", "entity source (default is standard input)")
			flags.StringVarP(new(string), "output", "o", yamlFormat, fmt.Sprintf(
				"output format, one of: %s|%s",
				jsonFormat, yamlFormat))
			flags.Bool("dry-run", false, "if true, only print the object that would be sent, without sending it")
			flags.StringVarP(&cnf.Union.GRPCConfig.Interface,
				"grpc-host", "", v.GetString("grpc_host"), "gRPC server host")
			flags.DurationVarP(&cnf.Union.GRPCConfig.Timeout,
				"timeout", "t", time.Second, "connection timeout")

			return nil
		},
	)
	License.AddCommand(registerLicense, createLicense, readLicense, updateLicense, deleteLicense, restoreLicense)
}
