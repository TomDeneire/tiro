package tui

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"tomdeneire.github.io/tiro/lib/database"
)

type model struct {
	list list.Model
}

const Filtering list.FilterState = 1

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "n" {
			if m.list.FilterState() != Filtering {
				Take(nil)
				m.Init()
				Search()
				return m, tea.Quit
			}
		}
		if msg.String() == "e" {
			if m.list.FilterState() != Filtering {
				item := m.list.SelectedItem()
				searchItem := item.(database.SearchItem)
				key := strings.Split(searchItem.ItemTitle, " ")[0]
				noteid, _ := strconv.Atoi(key)
				Take(noteid)
				m.Init()
				Search()
				return m, tea.Quit
			}
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

	m.list.Title = "Browse your notes"

	fn := func() []key.Binding {
		var keys []key.Binding
		editkey := key.NewBinding(
			key.WithKeys("edit", "e"),
			key.WithHelp("e", "edit"),
		)
		newkey := key.NewBinding(
			key.WithKeys("new", "n"),
			key.WithHelp("e", "new"),
		)
		keys = append(keys, newkey, editkey)
		return keys
	}
	m.list.AdditionalShortHelpKeys = fn
	m.list.AdditionalFullHelpKeys = fn

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
