package tui

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"tomdeneire.github.io/tiro/lib/database"
)

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "e" {
			item := m.list.SelectedItem()
			searchItem := item.(database.SearchItem)
			noteid, _ := strconv.Atoi(searchItem.ItemTitle)
			Take(noteid)
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width-10, msg.Height-10)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}

func Search() {

	items, err := database.GetSearchList(NotesFile)
	if err != nil {
		log.Fatalf("tui error: %v", err)
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "My Notes"

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
