package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var commandListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List commands",
	Long:    `Displays a list of all available iiiftool commands`,
	Example: "iiiftool command list",
	RunE:    commandList}

func init() {

	commandCmd.AddCommand(commandListCmd)
}

func commandList(cmd *cobra.Command, args []string) error {
	msg := map[string]map[string]string{}
	for _, command := range rootCmd.Commands() {
		msg[command.Use] = map[string]string{}
		for _, subCommand := range command.Commands() {
			parts := strings.SplitN(subCommand.Use, " ", 2)
			msg[command.Use][parts[0]] = subCommand.Short
		}
	}

	json, _ := json.MarshalIndent(msg, "", "    ")
	fmt.Println(string(json))

	return nil
}
