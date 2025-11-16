package components

import tea "github.com/charmbracelet/bubbletea"

type IFocus interface {
	tea.Model
	Focus()
	UnFocus()
}

type DefaultIFocus struct {
	IsFocused bool
}

func (m *DefaultIFocus) Focus() {
	m.IsFocused = true
}

func (m *DefaultIFocus) UnFocus() {
	m.IsFocused = false
}
