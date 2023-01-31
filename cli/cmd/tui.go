package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	figlet "tomdeneire.github.io/tiro/lib/figlet"
	qtui "tomdeneire.github.io/tiro/tui"
)

var tuiCmd = &cobra.Command{
	Use:     "tui",
	Short:   "tiro TUI",
	Long:    `Start the tiro TUI`,
	Args:    cobra.NoArgs,
	Example: `tiro tui`,
	RunE:    tui,
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}

func tui(cmd *cobra.Command, args []string) error {

	fmt.Println(figlet.Flogo)
	qtui.Views()

	return nil
}
