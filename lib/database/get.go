package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)

const searchQuery = `
select notes.key, notes.note, meta.action, meta.time, max(meta.key),
(select count() from notes) as total from notes
join meta on meta.noteid = notes.key
group by notes.key
order by notes.key desc`

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
	rows, err = db.Query(searchQuery)
	count := 0

	for rows.Next() {
		var item SearchItem
		var key string
		var action string
		var time string
		var max int
		var total int
		err := rows.Scan(&key, &item.ItemDesc, &action, &time, &max, &total)
		if err != nil {
			return nil, fmt.Errorf("cannot read value: %v", err)
		}
		if count == 0 {
			searchList = make([]list.Item, 0, total)
		}
		count++

		time = strings.ReplaceAll(time, "T", ", ")
		item.ItemTitle = key + " (" + action + " " + time + ")"
		searchList = append(searchList, item)
	}
	if err != nil {
		return nil, fmt.Errorf("cannot read file contents from archive: %v", err)
	}
	return searchList, err
}
