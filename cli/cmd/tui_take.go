package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	figlet "tomdeneire.github.io/tiro/lib/figlet"
	qtui "tomdeneire.github.io/tiro/tui"
)

var tuiTakeCmd = &cobra.Command{
	Use:     "take",
	Short:   "tiro TUI take",
	Long:    `Start the tiro TUI note taker`,
	Example: `tiro tui take`,
	RunE:    tuiTake}

func init() {
	tuiCmd.AddCommand(tuiTakeCmd)
}

func tuiTake(cmd *cobra.Command, args []string) error {

	fmt.Println(figlet.Flogo)
	qtui.Take()

	return nil
}
