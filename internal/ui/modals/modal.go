package modals

import tea "github.com/charmbracelet/bubbletea"

type Modal interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Modal, tea.Cmd)
	View() string

	Done() bool

	// SetSize tells the modal how much screen space it has to center
	// itself within (the full window, not a pre-shrunk box). Most modals
	// are small and fixed enough to ignore it safely, but any modal whose
	// content can grow with its input (an option list, a long message)
	// should use it to cap itself — otherwise, like a Panel ignoring its
	// allocated height, it can render taller than the screen and push its
	// own top border out of view.
	SetSize(width, height int)
}
