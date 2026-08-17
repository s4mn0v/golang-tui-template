package modals

import tea "github.com/charmbracelet/bubbletea"

type Modal interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Modal, tea.Cmd)
	View() string

	Done() bool
}
