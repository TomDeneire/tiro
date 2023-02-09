package database

import (
	"database/sql"
	"fmt"
	"strconv"
)

// Get all available info about a specific database note
func Info(noteid int, notesFile string) ([]Meta, error) {

	// Access database
	db, err := sql.Open("sqlite", notesFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open database: %v", err)
	}
	defer db.Close()

	identifier := strconv.Itoa(noteid)
	var rows *sql.Rows
	var result []Meta

	rows, err = db.Query("SELECT * FROM meta where noteid = ?", identifier)
	for rows.Next() {
		var meta Meta
		if err := rows.Scan(
			&meta.Key,
			&meta.Noteid,
			&meta.Time,
			&meta.Action,
			&meta.Cwd,
			&meta.User); err != nil {
			return nil, fmt.Errorf("cannot read value: %v", err)
		}
		result = append(result, meta)
	}
	if err != nil {
		return nil, fmt.Errorf("cannot read file contents from archive: %v", err)
	}

	return result, nil

}
