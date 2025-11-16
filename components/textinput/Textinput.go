package textinput

import (
	"bytes"
	"os"

	"github.com/alireza-karampour/qnq/components"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type Textinput = textinput.Model

type Model struct {
	components.DefaultIFocus
	Textinput
	buf    bytes.Buffer
	logger *log.Logger
}

func InitialModel() Model {
	l := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller: true,
	})
	return Model{
		buf:    bytes.Buffer{},
		logger: l,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.IsFocused {
			switch {
			case msg.String() == "backspace":
				m.buf.Truncate(max(m.buf.Len()-1, 0))
			case len(msg.Runes) == 1:
				_, err := m.buf.WriteRune(msg.Runes[0])
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
	m.SetValue(m.buf.String())
	m.Textinput, _ = m.Textinput.Update(msg)
	return m, nil
}

func (m Model) View() string {
	return ""
}
