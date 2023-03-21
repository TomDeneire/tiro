package database

import (
	"database/sql"
	"fmt"
)

const delNote = `DELETE FROM notes where key = ?;`
const delMeta = `DELETE FROM meta where noteid = ?;`

// Delete a database note (with note identifier as argument)
func Delete(noteid any, notesFile string) error {

	// Access database
	db, err := sql.Open("sqlite", notesFile)
	if err != nil {
		return fmt.Errorf("cannot open database: %v", err)
	}
	defer db.Close()

	// Delete note
	delStmt, err := db.Prepare(delNote)
	if err != nil {
		return fmt.Errorf("cannot prepare delnote statement: %v", err)
	}
	result, err := delStmt.Exec(noteid)
	if err != nil {
		return fmt.Errorf("cannot execute delnote statement: %v", err)
	}
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		fmt.Println(affected)
		return fmt.Errorf("cannot execute delnote statement: %v", err)
	}

	// Delete meta
	metaStmt, err := db.Prepare(delMeta)
	if err != nil {
		return fmt.Errorf("cannot prepare delmeta statement: %v", err)
	}
	result, err = metaStmt.Exec(noteid)
	if err != nil {
		return fmt.Errorf("cannot execute delmeta statement: %v", err)
	}
	affected, err = result.RowsAffected()
	if err != nil || affected == 0 {
		fmt.Println(affected)
		return fmt.Errorf("cannot execute delmeta statement: %v", err)
	}

	return nil
}
