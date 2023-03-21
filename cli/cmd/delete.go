package cmd

import (
	"log"

	"github.com/spf13/cobra"
	database "tomdeneire.github.io/tiro/lib/database"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a note",
	Long:    `Delete a note (one argument)`,
	Args:    cobra.ExactArgs(1),
	Example: `tiro delete 1234`,
	RunE:    del,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func del(cmd *cobra.Command, args []string) error {

	noteid := args[0]

	err := database.Delete(noteid, NotesFile)
	if err != nil {
		log.Fatalf("delete error: %v", err)
	}

	return nil
}
