package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	NextPanel key.Binding
	PrevPanel key.Binding
	Activate  key.Binding
	Filter    key.Binding
	Quit      key.Binding
}

var Keys = KeyMap{
	NextPanel: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next panel"),
	),
	PrevPanel: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "previous panel"),
	),
	Activate: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "open block/modal"),
	),
	Filter: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "filter blocks"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.NextPanel, k.Activate, k.Filter, k.Quit}
}
