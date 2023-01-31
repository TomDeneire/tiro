package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	figlet "tomdeneire.github.io/tiro/lib/figlet"
	qtui "tomdeneire.github.io/tiro/tui"
)

var tuiSearchCmd = &cobra.Command{
	Use:     "search",
	Short:   "tiro TUI search",
	Long:    `Start the tiro TUI search`,
	Example: `tiro tui search`,
	RunE:    tuiSearch}

func init() {
	tuiCmd.AddCommand(tuiSearchCmd)
}

func tuiSearch(cmd *cobra.Command, args []string) error {

	fmt.Println(figlet.Flogo)
	qtui.Search()

	return nil
}
