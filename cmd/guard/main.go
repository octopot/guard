package main

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/kamilsk/guard/pkg/cmd"
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

const (
	success = 0
	failure = 1
)

type cli func(executor commander, output io.Writer, shutdown func(code int))

type commander interface {
	AddCommand(...*cobra.Command)
	Execute() error
}
