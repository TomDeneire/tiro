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

const createMeta = `
CREATE TABLE meta (
	key INTEGER PRIMARY KEY AUTOINCREMENT,
    noteid INTEGER,
	time TEXT,
	action TEXT,
    cwd TEXT,
    user TEXT
);`

func Create(notesFile string) error {

	// Create directory

	notesDir := filepath.Dir(notesFile)
	err := os.Mkdir(notesDir, 0750)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("cannot create directory: %v", err)
	}

	// Create tables

	db, err := sql.Open("sqlite", notesFile)
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	defer db.Close()

	if _, err = db.Exec(createNotes); err != nil {
		return fmt.Errorf("cannot create table notes: %v", err)
	}

	if _, err = db.Exec(createMeta); err != nil {
		return fmt.Errorf("cannot create table meta: %v", err)
	}

	return nil
}
