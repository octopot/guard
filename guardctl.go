// +build cli ctl

package main

import (
	"io"
	"os"

	"github.com/kamilsk/guard/cmd"
	"github.com/spf13/cobra"
)

var control cli = func(executor commander, output io.Writer, shutdown func(code int)) {
	defer func() {
		if r := recover(); r != nil {
			shutdown(failure)
		}
	}()
	executor.AddCommand(cmd.Completion, cmd.Control, cmd.Version)
	if err := executor.Execute(); err != nil {
		shutdown(failure)
	}
	shutdown(success)
}

func main() { control(&cobra.Command{Use: "guardctl", Short: "Guard Control"}, os.Stderr, os.Exit) }
