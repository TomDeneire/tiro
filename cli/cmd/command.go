package cmd

import (
	"github.com/spf13/cobra"
)

var commandCmd = &cobra.Command{
	Use:     "command",
	Short:   "Command functions",
	Long:    `Working with tui commands`,
	Example: "tui command list",
}

func init() {

	rootCmd.AddCommand(commandCmd)
}
