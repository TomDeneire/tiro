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
				takeModel := initialTakeModel(nil)
				return takeModel, nil
			}
		}
		if msg.String() == "e" {
			if m.list.FilterState() != Filtering {
				item := m.list.SelectedItem()
				noteid, err := itemToNoteid(item)
				if err != nil {
					log.Fatalf("Error getting noteid: %v", err)
				}
				takeModel := initialTakeModel(noteid)
				return takeModel, nil
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

func itemToNoteid(item list.Item) (int, error) {
	searchItem := item.(database.SearchItem)
	key := strings.Split(searchItem.ItemTitle, " ")[0]
	noteid, err := strconv.Atoi(key)
	if err != nil {
		return 0, err
	}
	return noteid, nil
}

func initialListModel() (model, error) {

	var m model

	items, err := database.GetSearchList(NotesFile)
	if err != nil {
		return m, err
	}

	m = model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}

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

	return m, nil
}

func Search() {

	m, err := initialListModel()
	if err != nil {
		log.Fatalf("Error making list model: %v", err)
	}

	p := tea.NewProgram(m)

	_, err = p.Run()

	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
