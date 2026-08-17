package panels

import tea "github.com/charmbracelet/bubbletea"

type Panel interface {
	Init() tea.Cmd

	Update(msg tea.Msg) (Panel, tea.Cmd)

	View() string

	SetSize(width, height int)

	Focus()
	Blur()
	Focused() bool

	Focusable() bool

	Title() string
}
