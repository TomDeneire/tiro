package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var argFileCmd = &cobra.Command{
	Use:   "file",
	Short: "Start iiiftool with arguments in a file",
	Long: `Launches iiiftool with the arguments specified as lines in a file. 
If the first non-whitespace character in the file is a *[*, the contents should be a *JSON array*.

If the file is *NOT* a JSON array, the following restriction apply:
    - Arguments are read line-by-line from the named file.
    - Whitespace is stripped at the beginning and the end of each line
    - Empty lines are skipped
    - The first line should be *iiiftool*
	
If the file *IS* a JSON array, the following applies:
    - The first element should always be *iiiftool*
    - Whitespace is never stripped
    - Empty arguments remain in the argument list
`,

	Args:    cobra.ExactArgs(1),
	Example: `iiiftool arg file myargs.txt`,
	RunE:    argFile,
}

func init() {
	argCmd.AddCommand(argFileCmd)
}

func argFile(cmd *cobra.Command, args []string) error {
	jarg := args[0]

	if jarg == "" {
		return errors.New("argument is empty")
	}
	_, err := os.ReadFile(jarg)

	return err
}
