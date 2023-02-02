package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	figlet "tomdeneire.github.io/tiro/lib/figlet"
	qtui "tomdeneire.github.io/tiro/tui"
)

var tuiTakeCmd = &cobra.Command{
	Use:   "take",
	Short: "tiro TUI take",
	Long:  `Start the tiro TUI note taker`,
	Args:  cobra.RangeArgs(0, 1),
	Example: `tiro tui take"
tiro tui take 1234`,
	RunE: tuiTake}

func init() {
	tuiCmd.AddCommand(tuiTakeCmd)
}

func tuiTake(cmd *cobra.Command, args []string) error {

	fmt.Println(figlet.Flogo)

	var noteid any
	var err error

	if len(args) == 1 {
		noteid, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid note identifier: %v", err)
		}
	}

	qtui.Take(noteid)

	return nil
}
