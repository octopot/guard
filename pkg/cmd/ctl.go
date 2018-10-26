package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/kamilsk/go-kit/pkg/strings"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

	// TODO issue#draft {

	employee        = &cobra.Command{Use: "employee", Short: "Works with client employees (draft)"}
	addEmployee     = &cobra.Command{Use: "add", Short: "Add employee to a client license", RunE: communicate}
	deleteEmployee  = &cobra.Command{Use: "delete", Short: "Delete employee from a client license", RunE: communicate}
	workplace       = &cobra.Command{Use: "workplace", Short: "Works with client workplaces (draft)"}
	addWorkplace    = &cobra.Command{Use: "add", Short: "Add workplace to a client license", RunE: communicate}
	deleteWorkplace = &cobra.Command{Use: "delete", Short: "Delete workplace from a client license", RunE: communicate}
	pushWorkplace   = &cobra.Command{Use: "push", Short: "Push workplace of a client license up", RunE: communicate}

	// issue#draft }
)

func init() {
	v := viper.New()

	configure := func(flags *pflag.FlagSet) *pflag.FlagSet {
		flags.StringVarP(new(string), "filename", "f", "", "entity source (default is standard input)")
		flags.StringVarP(new(string), "output", "o", yamlFormat, fmt.Sprintf(
			"output format, one of: %s|%s",
			jsonFormat, yamlFormat))
		flags.Bool("dry-run", false, "if true, only print the object that would be sent, without sending it")
		flags.StringVarP(&cnf.Union.GRPCConfig.RPC.Interface,
			"grpc-host", "", v.GetString("grpc_host"), "gRPC server host")
		flags.DurationVarP(&cnf.Union.GRPCConfig.RPC.Timeout,
			"timeout", "t", time.Second, "connection timeout")
		return flags
	}

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
			configure(Install.Flags())
			configure(License.PersistentFlags()).
				StringVarP((*string)(&cnf.Union.GRPCConfig.RPC.Token),
					"token", "", v.GetString("guard_token"), "user access token")
			return nil
		},
	)
	License.AddCommand(registerLicense, createLicense, readLicense, updateLicense, deleteLicense, restoreLicense)

	// TODO issue#draft {

	License.AddCommand(employee, workplace)
	employee.AddCommand(addEmployee, deleteEmployee)
	workplace.AddCommand(addWorkplace, deleteWorkplace, pushWorkplace)

	// issue#draft }

}
