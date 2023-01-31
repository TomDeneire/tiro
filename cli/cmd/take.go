package cmd

import (
	"log"
	"strconv"

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
	var err error

	if len(args) == 2 {
		noteid, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("take error, noteid invalid: %v", err)
		}
	}

	content := args[0]

	err = database.Set(content, noteid)
	if err != nil {
		log.Fatalf("take error: %v", err)
	}

	return nil
}
