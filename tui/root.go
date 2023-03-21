package tui

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
)

// Name of the notes database file
var NotesFile = filepath.Join(os.Getenv("HOME"), ".tiro", "tiro.sqlite")

// Default caption style
var DefaultCaptionStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#99c794")).
	Foreground(lipgloss.Color("230")).
	Padding(0, 1)
