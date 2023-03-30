package cmd

import (
	"io"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	database "tomdeneire.github.io/tiro/lib/database"
)

var takeCmd = &cobra.Command{
	Use:   "take",
	Short: "Take a note",
	Long: `Take a note; either a new note or overwrite an existing one.
Can also be used by piping from stdin`,
	Args: cobra.RangeArgs(0, 2),
	Example: `tiro take "hello world"
tiro take "hello world" 1234
cat myfile.txt | tiro take`,
	RunE: take,
}

func init() {
	rootCmd.AddCommand(takeCmd)
}

func take(cmd *cobra.Command, args []string) error {

	var noteid any
	var err error
	var content string

	if len(args) == 2 {
		noteid, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("take error, noteid invalid: %v", err)
		}
	}

	if len(args) == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("take error: %v", err)
		}
		content = string(data)
	} else {
		content = args[0]
	}

	err = database.Set(content, noteid, NotesFile)
	if err != nil {
		log.Fatalf("take error: %v", err)
	}

	return nil
}
