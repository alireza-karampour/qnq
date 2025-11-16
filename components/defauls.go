package components

import "github.com/charmbracelet/lipgloss"

var TitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Padding(0, 1)

var SubtitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF")).
	Bold(true).
	MarginTop(1).
	MarginBottom(1)

var CursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#7D56F4")).
	Bold(true)

var SelectedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00FF00")).
	Bold(true)

var ItemStyle = lipgloss.NewStyle().
	PaddingLeft(2)

var HelpStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#626262")).
	MarginTop(2)
