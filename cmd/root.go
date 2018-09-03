package cmd

import "github.com/spf13/cobra"

// RootCmd is the entry point of command-line interface.
var RootCmd = &cobra.Command{Use: "guard", Short: "Guard"}

func init() { RootCmd.AddCommand(completionCmd, runCmd) }
