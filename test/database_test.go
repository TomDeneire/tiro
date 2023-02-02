package test

import (
	"fmt"
	"os"
	"testing"

	_ "modernc.org/sqlite"
	"tomdeneire.github.io/tiro/lib/database"
	"tomdeneire.github.io/tiro/lib/util"
)

func TestCreate(t *testing.T) {
	_, err := os.Stat(NotesTFile)
	if !os.IsNotExist(err) {
		os.Remove(NotesTFile)
	}
	err = database.Create(NotesTFile)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nError creating database: \n%v\n", err))
	}
}

func Test1Set(t *testing.T) {
	err := database.Set("hello world", nil, NotesTFile)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nError setting row (1): \n%v\n", err))
	}
}

func Test2Set(t *testing.T) {
	err := database.Set("bonjour tout le monde", nil, NotesTFile)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nError setting row (2): \n%v\n", err))
	}
}

func TestOverwrite(t *testing.T) {
	err := database.Set("goodbye world", 1, NotesTFile)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nError overwriting row: \n%v\n", err))
	}
}

func TestRead(t *testing.T) {
	records := []int{1, 2}
	expected := []string{"goodbye world", "bonjour tout le monde"}
	for i, rec := range records {
		contents, err := database.Get(rec, NotesTFile)
		if err != nil {
			t.Errorf(fmt.Sprintf("\nError reading row: \n%v\n", err))
		}
		util.Check(contents, expected[i], t)
	}
}
