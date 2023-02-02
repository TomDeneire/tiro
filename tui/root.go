package tui

import (
	"os"
	"path/filepath"
)

// Name of the notes database file
var NotesFile = filepath.Join(os.Getenv("HOME"), ".tiro", "tiro.sqlite")
