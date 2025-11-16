package login

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width    int
	height   int
	username textinput.Model
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Username"
	return Model{
		username: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.username.Width = m.width / 3
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			m.username.SetValue(m.username.Value() + msg.String())
		}
	}
	m.username, _ = m.username.Update(msg)

	return m, nil
}

func (m Model) View() string {
	if m.width == 0 {
		return "Initializing..."
	}
	usernameV := m.username.View()
	content := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, usernameV)
	return content
}
