package main

import (
	"io"

	"github.com/spf13/cobra"
)

const (
	success = 0
	failed  = 1
)

type cli func(executor commander, output io.Writer, shutdown func(code int))

type commander interface {
	AddCommand(...*cobra.Command)
	Execute() error
}
