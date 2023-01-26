package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var argURLCmd = &cobra.Command{
	Use:   "url",
	Short: "Start iiiftool with arguments retrieved by URL",
	Long: `Launches iiiftool with the arguments retrieved by URL.
	
If the first non-whitespace character is a *[*, the contents should be a JSON array.

If the input is *NOT* a JSON array, the following restriction apply:
	- Arguments are read line-by-line from the input
	- Whitespace is stripped at the beginning and the end of each line
	- Empty lines are skipped
	- The first line should be *iiiftool*
	
If the input *IS* a JSON array, the following applies:
	- The first element should always be *iiiftool*
	- Whitespace is never stripped
	- Empty arguments remain in the argument list
		`,
	Args:    cobra.ExactArgs(1),
	Example: `iiiftool arg url https://dev.anet.be/about.html`,
	RunE:    argURL,
}

func init() {
	argCmd.AddCommand(argURLCmd)
}

func argURL(cmd *cobra.Command, args []string) error {
	jarg := args[0]

	if jarg == "" {
		return errors.New("srgument is empty")
	}
	return nil

}
