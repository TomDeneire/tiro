package tui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"tomdeneire.github.io/tiro/lib/database"
)

func Take() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type errMsg error

type takeModel struct {
	textarea textarea.Model
	err      error
}

func initialModel() takeModel {
	ti := textarea.New()
	ti.Placeholder = "..."
	ti.Focus()

	return takeModel{
		textarea: ti,
		err:      nil,
	}
}

func (m takeModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m takeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			err := database.Set(m.textarea.Value(), nil)
			if err != nil {
				log.Fatalf("tui take error: %v", err)
			}
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m takeModel) View() string {
	return fmt.Sprintf(
		"Start taking your note!\n\n%s\n\n%s",
		m.textarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}
