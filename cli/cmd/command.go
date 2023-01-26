package cmd

import (
	"github.com/spf13/cobra"
)

var commandCmd = &cobra.Command{
	Use:     "command",
	Short:   "Command functions",
	Long:    `Working with iiiftool commands`,
	Example: "iiiftool command list",
}

func init() {

	rootCmd.AddCommand(commandCmd)
}
