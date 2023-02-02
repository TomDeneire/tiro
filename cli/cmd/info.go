package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	database "tomdeneire.github.io/tiro/lib/database"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "note info",
	Long:    `info about a specific note`,
	Args:    cobra.ExactArgs(1),
	Example: `tiro info 1234`,
	RunE:    info,
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func info(cmd *cobra.Command, args []string) error {

	noteid, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("info error: invalid note id: %v", err)
	}
	result, err := database.Info(noteid, NotesFile)
	if err != nil {
		log.Fatalf("info error: %v", err)
	}
	for _, res := range result {
		fmt.Println(res)
	}

	return nil
}
