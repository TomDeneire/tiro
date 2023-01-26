package cmd

import (
	"github.com/spf13/cobra"
)

var argCmd = &cobra.Command{
	Use:   "arg",
	Short: "Alternative ways to start iiiftool",
	Long: `A a CLI application *iiiftool* can be started by specifying the the arguments and flags on the command line.
These are not always the most convenient ways. *arg* specifies several alternatives.`,
	Args:    cobra.NoArgs,
	Example: "iiiftool arg",
}

func init() {
	rootCmd.AddCommand(argCmd)
}
