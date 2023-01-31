package tui

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/*
This example assumes an existing understanding of commands and messages. If you
haven't already read our tutorials on the basics of Bubble Tea and working
with commands, we recommend reading those first.
Find them at:
https://github.com/charmbracelet/bubbletea/tree/master/tutorials/commands
https://github.com/charmbracelet/bubbletea/tree/master/tutorials/basics
*/

// sessionState is used to track which model is focused
type sessionState uint

const (
	defaultTime              = time.Minute
	takerView   sessionState = iota
	searcherView
)

var (
	modelStyle = lipgloss.NewStyle().
			Width(55).
			Height(20).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.HiddenBorder())
	focusedModelStyle = lipgloss.NewStyle().
				Width(55).
				Height(20).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("69"))
	viewhelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type mainModel struct {
	state    sessionState
	taker    tea.Model
	searcher tea.Model
	index    int
}

func newViewModel(timeout time.Duration) mainModel {
	m := mainModel{state: takerView}
	m.taker = initialModel()
	m.searcher = newModel()
	return m
}

func (m mainModel) Init() tea.Cmd {
	return tea.Batch(m.taker.Init(), m.searcher.Init())
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.state == takerView {
				m.state = searcherView
			} else {
				m.state = takerView
			}
		}
		switch m.state {
		// update whichever model is focused
		case searcherView:
			m.searcher, cmd = m.searcher.Update(msg)
			cmds = append(cmds, cmd)
		default:
			m.taker, cmd = m.taker.Update(msg)
			cmds = append(cmds, cmd)
		}
	}
	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	var s string
	model := m.currentFocusedModel()
	if m.state == takerView {
		s += lipgloss.JoinHorizontal(lipgloss.Top, focusedModelStyle.Render(fmt.Sprintf("%4s", m.taker.View())), modelStyle.Render(m.searcher.View()))
	} else {
		s += lipgloss.JoinHorizontal(lipgloss.Top, modelStyle.Render(fmt.Sprintf("%4s", m.taker.View())), focusedModelStyle.Render(m.searcher.View()))
	}
	s += viewhelpStyle.Render(fmt.Sprintf("\ntab: focus next • n: new %s • q: exit\n", model))
	return s
}

func (m mainModel) currentFocusedModel() string {
	if m.state == takerView {
		return "taker"
	}
	return "searcher"
}

func Views() {
	p := tea.NewProgram(newViewModel(defaultTime))

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
