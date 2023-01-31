package cmd

import (
	"github.com/spf13/cobra"
	database "tomdeneire.github.io/tiro/lib/database"
)

var takeCmd = &cobra.Command{
	Use:   "take",
	Short: "Take a note",
	Long:  `Save a note; either a new note or overwrite an existing one`,
	Args:  cobra.RangeArgs(1, 2),
	Example: `tiro take "hello world"
tiro take "hello world" 1234`,
	RunE: take,
}

func init() {
	rootCmd.AddCommand(takeCmd)
}

func take(cmd *cobra.Command, args []string) error {

	var noteid any
	if len(args) == 2 {
		noteid = args[1]
	}

	err := database.Set(args[0], noteid)
	if err != nil {
		panic(err)
	}

	return nil
}
