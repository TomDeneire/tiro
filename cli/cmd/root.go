package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	figlet "tomdeneire.github.io/tiro/lib/figlet"
)

// Name of the notes database file
var NotesFile = filepath.Join(os.Getenv("HOME"), ".tiro", "tiro.sqlite")

// BuildTime defined by compilation
var BuildTime = ""

// GoVersion defined by compilation
var GoVersion = ""

// BuildHost defined by compilation
var BuildHost = ""

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(buildTime string, goVersion string, buildHost string, args []string) {
	BuildTime = buildTime
	BuildHost = buildHost
	GoVersion = goVersion
	rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:           "tiro",
	Short:         "tiro - CLI application for note taking",
	SilenceUsage:  true,
	SilenceErrors: true,
	Long:          figlet.Flogo + "\ntiro is a CLI application for note taking",
}
