package cmd

import (
	"github.com/spf13/cobra"
)

var argStdinCmd = &cobra.Command{
	Use:   "stdin",
	Short: "Start iiiftool with arguments read from stdin",
	Long: `Launches iiiftool with the arguments as lines on stdin. 
	
If the first non-whitespace character is a *[*, the contents should be a JSON array.

If the input is *NOT* a JSON array, the following restriction apply:
	- Arguments are read line-by-line from *stdin*
	- Whitespace is stripped at the beginning and the end of each line
	- Empty lines are skipped
	- The first line should be *iiiftool*
	
If the input *IS* a JSON array, the following applies:
	- The first element should always be *iiiftool*
	- Whitespace is never stripped
	- Empty arguments remain in the argument list`,
	Args:    cobra.NoArgs,
	Example: `iiiftool arg stdin < commands.txt`,
	RunE:    argStdin,
}

func init() {
	argCmd.AddCommand(argStdinCmd)
}

func argStdin(cmd *cobra.Command, args []string) error {
	return nil
}
