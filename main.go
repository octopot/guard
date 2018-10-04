// +build !ctl

package main

import (
	"io"
	"os"

	"github.com/kamilsk/guard/cmd"
	"github.com/spf13/cobra"
)

var service cli = func(executor commander, output io.Writer, shutdown func(code int)) {
	defer func() {
		if r := recover(); r != nil {
			shutdown(failure)
		}
	}()
	executor.AddCommand(cmd.Completion, cmd.Migrate, cmd.Run, cmd.Version)
	if err := executor.Execute(); err != nil {
		shutdown(failure)
	}
	shutdown(success)
}

func main() {
	service(&cobra.Command{Use: "guard", Short: "Guard Service"}, os.Stderr, os.Exit)
}
