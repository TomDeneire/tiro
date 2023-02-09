package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "modernc.org/sqlite"
)

const newNote = "INSERT INTO notes (key, note) Values($1,$2)"
const newMeta = "INSERT INTO meta (key, noteid, time, action, cwd, user) Values($1,$2,$3,$4,$5,$6)"

const overwriteNote = `UPDATE notes SET note = ? WHERE key = ?;`

// Sets a database note; either a new one (without second argument)
// or overwrites an existing one (with note identifier as second argument)
func Set(contents string, identifier any, notesFile string) error {

	// Check if database exists
	_, err := os.Stat(notesFile)
	if os.IsNotExist(err) {
		createErr := Create(notesFile)
		if createErr != nil {
			return fmt.Errorf("cannot create database: %v", createErr)
		}
	}

	// Construct note
	var note Note
	note.Contents = contents
	if identifier != nil {
		note.Key = identifier.(int)
	}

	// Construct meta data
	var meta Meta
	meta.Time = time.Now().Format("2006-01-02T15:04:05")
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("cannot get current work directory: %v", err)
	}
	meta.Cwd = cwd
	meta.User = os.Getenv("USER")

	// Access database
	db, err := sql.Open("sqlite", notesFile)
	if err != nil {
		return fmt.Errorf("cannot open database: %v", err)
	}
	defer db.Close()

	// Set note

	var setErr error
	if identifier == nil {
		setErr = New(&note, &meta, db)
	} else {
		setErr = Overwrite(&note, &meta, db)
	}
	if setErr != nil {
		return fmt.Errorf("cannot set note: %v", setErr)
	}

	return nil
}

func New(note *Note, meta *Meta, db *sql.DB) error {

	newStmt, err := db.Prepare(newNote)
	if err != nil {
		return fmt.Errorf("cannot prepare newnote statement: %v", err)
	}
	result, err := newStmt.Exec(nil, note.Contents)
	if err != nil {
		return fmt.Errorf("cannot execute newnote statement: %v", err)
	}

	noteid, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("cannot obtain new identifier: %v", err)
	}

	metaStmt, err := db.Prepare(newMeta)
	if err != nil {
		return fmt.Errorf("cannot prepare newmeta statement: %v", err)
	}
	_, err = metaStmt.Exec(nil, noteid, meta.Time, "created", meta.Cwd, meta.User)
	if err != nil {
		return fmt.Errorf("cannot execute newmeta statement: %v", err)
	}

	return nil
}

func Overwrite(note *Note, meta *Meta, db *sql.DB) error {

	overwriteStmt, err := db.Prepare(overwriteNote)
	if err != nil {
		return fmt.Errorf("cannot prepare overwritenote statement: %v", err)
	}

	// needs to be string for SQL query!
	key := strconv.Itoa(note.Key)
	result, err := overwriteStmt.Exec(note.Contents, key)
	if err != nil {
		return fmt.Errorf("cannot execute overwritenote statement: %v", err)
	}
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		fmt.Println(affected)
		return fmt.Errorf("cannot execute overwritenote statement: %v", err)
	}

	metaStmt, err := db.Prepare(newMeta)
	if err != nil {
		return fmt.Errorf("cannot prepare overwritemeta statement: %v", err)
	}
	_, err = metaStmt.Exec(nil, key, meta.Time, "modified", meta.Cwd, meta.User)
	if err != nil {
		return fmt.Errorf("cannot execute overwritemeta statement: %v", err)
	}

	return nil
}
