package cmd

import (
	"encoding/json"
	"errors"

	"github.com/spf13/cobra"
)

var argJSONCmd = &cobra.Command{
	Use:   "json",
	Short: "Start iiiftool with arguments in JSON",
	Long: `Launches iiiftool with the arguments specified in a JSON string.
	
The command works with exactly one argument: a string containing a *JSON array*.

The following applies:
    - The first element should always be *iiiftool*
    - Whitespace is never stripped
    - Empty arguments remain in the argument list
`,
	Args:    cobra.ExactArgs(1),
	Example: `iiiftool arg json '["iiiftool", "about"]'`,
	RunE:    argJSON,
}

func init() {
	argCmd.AddCommand(argJSONCmd)
}

func argJSON(cmd *cobra.Command, args []string) error {
	jarg := args[0]

	if jarg == "" {
		return errors.New("argument is empty")
	}

	argums := make([]string, 0)

	err := json.Unmarshal([]byte(jarg), &argums)

	return err

}
