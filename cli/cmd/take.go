package cmd

import (
	"github.com/spf13/cobra"
)

var takeCmd = &cobra.Command{
	Use:     "take",
	Short:   "Take note",
	Long:    `Take down a note`,
	Args:    cobra.ExactArgs(1),
	Example: `tiro take "hello world"`,
	RunE:    take,
}

func init() {
	rootCmd.AddCommand(takeCmd)
}

func take(cmd *cobra.Command, args []string) error {

	return nil
}
