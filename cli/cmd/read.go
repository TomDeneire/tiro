package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	database "tomdeneire.github.io/tiro/lib/database"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a note",
	Long:  `Read a note; either the last note (no arguments) or a specific one (one argument)`,
	Args:  cobra.RangeArgs(0, 1),
	Example: `tiro read
tiro read 1234`,
	RunE: read,
}

func init() {
	rootCmd.AddCommand(readCmd)
}

func read(cmd *cobra.Command, args []string) error {

	var noteid any
	if len(args) == 1 {
		noteid = args[0]
	}

	result, err := database.Get(noteid)
	if err != nil {
		log.Fatalf("read error: %v", err)
	}
	fmt.Println(result)

	return nil
}
