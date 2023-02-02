package test

import (
	"os"
	"path/filepath"
)

// Name of the notes database file
var NotesTFile = filepath.Join(os.Getenv("HOME"), ".tiro", "test.sqlite")
