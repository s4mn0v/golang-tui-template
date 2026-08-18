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

// TextCapturing is implemented by panels that, while focused, consume
// arbitrary printable keystrokes as text (single-line/multi-line inputs,
// filterable lists, etc). A host app should check this before firing
// single-key global shortcuts — such as a bare "q" to quit — so typing the
// letter doesn't get swallowed by the shortcut instead of the field.
type TextCapturing interface {
	CapturesText() bool
}
