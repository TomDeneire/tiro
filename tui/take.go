package tui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"tomdeneire.github.io/tiro/lib/database"
)

type Caption struct {
	Style   lipgloss.Style
	Message string
}

func (c Caption) View() string {
	var style lipgloss.Style
	return style.Render(c.Style.Render(fmt.Sprintf(c.Message)))
}

func Take(noteid any) {
	m := initialTakeModel(noteid)
	p := tea.NewProgram(m)

	_, err := p.Run()

	if err != nil {
		log.Fatalf("TUI error: %v", err)
	}
}

type errMsg error

type takeModel struct {
	err      error
	noteid   any
	textarea textarea.Model
}

func initialTakeModel(noteid any) takeModel {
	ti := textarea.New()
	ti.Placeholder = "..."
	ti.CharLimit = 0
	ti.FocusedStyle.CursorLine = ti.BlurredStyle.CursorLine
	ti.FocusedStyle.Text = ti.BlurredStyle.Text
	ti.EndOfBufferCharacter = ' '
	ti.SetWidth(100)
	ti.CursorEnd()
	ti.Focus()

	if noteid != nil {
		contents, err := database.Get(noteid, NotesFile)
		if err != nil {
			log.Fatalf("cannot open note: %v", err)
		}
		ti.SetValue(contents)
	}

	return takeModel{
		err:      nil,
		noteid:   noteid,
		textarea: ti,
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
			contents := m.textarea.Value()
			if contents != "" {
				err := database.Set(contents, m.noteid, NotesFile)
				if err != nil {
					log.Fatalf("tui take error: %v", err)
				}
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
	var c Caption
	c.Message = "Start taking your note!"
	c.Style = DefaultCaptionStyle
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		c.View(),
		m.textarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}
