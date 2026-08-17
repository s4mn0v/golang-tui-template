package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	NextPanel key.Binding
	PrevPanel key.Binding
	Quit      key.Binding

	OpenConfirmDemo key.Binding
	OpenAlertDemo   key.Binding
}

var Keys = KeyMap{
	NextPanel: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "siguiente panel"),
	),
	PrevPanel: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "panel anterior"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "salir"),
	),

	OpenConfirmDemo: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "eliminar (demo confirm)"),
	),
	OpenAlertDemo: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "alerta (demo)"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.NextPanel, k.PrevPanel, k.OpenConfirmDemo, k.OpenAlertDemo, k.Quit}
}
