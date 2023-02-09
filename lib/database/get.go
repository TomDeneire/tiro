package database

import (
	"database/sql"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
)

type SearchItem struct {
	ItemTitle, ItemDesc string
}

func (i SearchItem) Title() string       { return i.ItemTitle }
func (i SearchItem) Description() string { return i.ItemDesc }
func (i SearchItem) FilterValue() string { return i.ItemDesc }

// Gets a database note; either the most recent one (without argument)
// or a specific one (with note identifier as argument)
func Get(noteid any, notesFile string) (string, error) {

	// Access database
	db, err := sql.Open("sqlite", notesFile)
	if err != nil {
		return "", fmt.Errorf("cannot open database: %v", err)
	}
	defer db.Close()

	var note Note

	var row *sql.Row
	if noteid != nil {
		row = db.QueryRow("SELECT * FROM notes WHERE key = ?", noteid)
	} else {
		row = db.QueryRow("SELECT * FROM notes ORDER BY key DESC LIMIT 1")
	}
	err = ReadNoteRow(row, &note)
	if err != nil {
		return "", fmt.Errorf("cannot read file contents from archive: %v", err)
	}

	return note.Contents, nil
}

// Function that reads a single notes sql.Row
func ReadNoteRow(row *sql.Row, note *Note) error {

	err := row.Scan(
		&note.Key,
		&note.Contents)
	if err != nil {
		return err
	}
	return nil
}

func GetSearchList(notesFile string) (searchList []list.Item, err error) {

	// Access database
	db, err := sql.Open("sqlite", notesFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open database: %v", err)
	}
	defer db.Close()

	var rows *sql.Rows
	query := `select notes.key, notes.note, meta.action, meta.time from notes
    join meta on meta.key = notes.key
    order by notes.key desc`
	rows, err = db.Query(query)
	for rows.Next() {
		var item SearchItem
		var key string
		var action string
		var time string
		if err := rows.Scan(&key, &item.ItemDesc, &action, &time); err != nil {
			return nil, fmt.Errorf("cannot read value: %v", err)
		}
		item.ItemTitle = key + " (" +action + " " + time + ")"
		searchList = append(searchList, item)
	}
	if err != nil {
		return nil, fmt.Errorf("cannot read file contents from archive: %v", err)
	}
	return searchList, err
}
