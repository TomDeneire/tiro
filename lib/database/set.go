package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

const newNote = "INSERT INTO notes (key, note) Values($1,$2)"
const newAdmin = "INSERT INTO admin (key, noteid, time, action) Values($1,$2,$3,$4)"

const overwriteNote = `UPDATE notes SET note = "?" WHERE key = ?;`

// Sets a database note; either a new one (without second argument)
// or overwrites an existing one (with note identifier as second argument)
func Set(contents string, identifier any) error {

	// Check if database exists
	_, err := os.Stat(NotesFile)
	if os.IsNotExist(err) {
		createErr := Create()
		if createErr != nil {
			return fmt.Errorf("cannot create database: %v", createErr)
		}
	}

	// Construct note

	var note Note
	note.Contents = contents

	// Access database
	db, err := sql.Open("sqlite", NotesFile)
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	defer db.Close()

	// Set note

	var setErr error
	if identifier == nil {
		setErr = New(&note, db)
	} else {
		setErr = Overwrite(&note, db, identifier)
	}
	if setErr != nil {
		return fmt.Errorf("cannot set note: %v", setErr)
	}

	return nil
}

func New(note *Note, db *sql.DB) error {
	mtime := time.Now().Unix() // must be int in sqlar specification!

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

	adminStmt, err := db.Prepare(newAdmin)
	if err != nil {
		return fmt.Errorf("cannot prepare newadmin statement: %v", err)
	}
	_, err = adminStmt.Exec(nil, noteid, mtime, "created")
	if err != nil {
		return fmt.Errorf("cannot execute newadmin statement: %v", err)
	}

	return nil
}

func Overwrite(note *Note, db *sql.DB, noteid any) error {
	mtime := time.Now().Unix() // must be int in sqlar specification!

	overwriteStmt, err := db.Prepare(overwriteNote)
	if err != nil {
		return fmt.Errorf("cannot prepare overwritenote statement: %v", err)
	}

	result, err := overwriteStmt.Exec(noteid, note.Contents)
	if err != nil {
		return fmt.Errorf("cannot execute overwritenote statement: %v", err)
	}
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return fmt.Errorf("cannot execute overwritenote statement: %v", err)
	}

	adminStmt, err := db.Prepare(newAdmin)
	if err != nil {
		return fmt.Errorf("cannot prepare overwriteadmin statement: %v", err)
	}
	_, err = adminStmt.Exec(nil, noteid, mtime, "modified")
	if err != nil {
		return fmt.Errorf("cannot execute overwritadmin statement: %v", err)
	}

	return nil
}

// mtime := time.Now().Unix() // must be int in sqlar specification!
// props, _ := basefs.Properties("nakedfile")
// mode := int64(props.PERM)
// sz := int64(len(data))
// _, err = stmt1.Exec(name, mode, mtime, sz, data)
// if err != nil {
//     return fmt.Errorf("cannot exec stmt1: %v", err)
// }
// _, err = stmt3.Exec(nil, docman, name)
// if err != nil {
//     return fmt.Errorf("cannot exec stmt3: %v", err)
// }
// stmt1, err := db.Prepare("INSERT INTO sqlar (name, mode, mtime, sz, data) Values($1,$2,$3,$4,$5)")
// if err != nil {
//     return fmt.Errorf("cannot prepare insert1: %v", err)
// }
// defer stmt1.Close()
//
// stmt2, err := db.Prepare(insertAdmin)
// if err != nil {
//     return fmt.Errorf("cannot prepare insert2: %v", err)
// }
// defer stmt2.Close()
//
// stmt3, err := db.Prepare("INSERT INTO files (key, docman, name) Values($1,$2,$3)")
// if err != nil {
//     return fmt.Errorf("cannot prepare insert3: %v", err)
// }
// defer stmt3.Close()
//
// stmt4, err := db.Prepare(insertMeta)
// if err != nil {
//     return fmt.Errorf("cannot prepare insert4: %v", err)
// }
// defer stmt4.Close()
//
// // Insert into "admin", "files"
//
// content, err := json.Marshal(iiifMeta.Manifest)
// manifest := string(content)
// if err != nil {
//     return fmt.Errorf("json error on stmt4: %v", err)
// }
