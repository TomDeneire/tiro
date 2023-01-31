package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// CONSTANTS

const createNotes = `
CREATE TABLE notes (
	key INTEGER PRIMARY KEY AUTOINCREMENT,
	note TEXT
);`

const createAdmin = `
CREATE TABLE admin (
	key INTEGER PRIMARY KEY AUTOINCREMENT,
    noteid INTEGER,
	time TEXT,
	action TEXT
);`

var notesDir = filepath.Join(os.Getenv("HOME"), ".tiro")
var NotesFile = filepath.Join(notesDir, "tiro.sqlite")

func Create() error {

	// Create directory

	err := os.Mkdir(notesDir, 0750)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("cannot create directory: %v", err)
	}

	// Create tables

	db, err := sql.Open("sqlite", NotesFile)
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	defer db.Close()

	if _, err = db.Exec(createNotes); err != nil {
		return fmt.Errorf("cannot create table notes: %v", err)
	}

	if _, err = db.Exec(createAdmin); err != nil {
		return fmt.Errorf("cannot create table admin: %v", err)
	}

	return nil
}
