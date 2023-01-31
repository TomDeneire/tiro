package cmd

import (
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:     "tui",
	Short:   "tiro TUI",
	Long:    `Start the tiro TUI`,
	Example: `tiro tui search`,
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
