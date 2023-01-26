package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/spf13/cobra"

	figlet "tomdeneire.github.io/tiro/lib/figlet"
)

var aboutCmd = &cobra.Command{
	Use:     "about",
	Short:   "Information about `tiro`",
	Long:    `Version and build time information about the tiro executable.`,
	Args:    cobra.NoArgs,
	Example: `tiro about`,
	RunE:    about,
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}

func about(cmd *cobra.Command, args []string) error {
	msg := map[string]string{"BuildTime": BuildTime, "BuildHost": BuildHost, "BuildWith": GoVersion}
	host, e := os.Hostname()

	if e == nil {
		msg["uname"] = host
	}
	user, err := user.Current()
	if err == nil {
		msg["user.name"] = user.Name
		msg["user.username"] = user.Username
	}
	b, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(figlet.Flogo)

	fmt.Println(string(b))

	return nil
}
