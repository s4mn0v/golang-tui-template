package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	NextPanel key.Binding
	PrevPanel key.Binding
	Activate  key.Binding
	Filter    key.Binding
	Quit      key.Binding
	ForceQuit key.Binding
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
	// Quit only fires outside of text entry/filtering — see
	// Model.capturingText — so a bare "q" types normally when the user is
	// typing somewhere. ForceQuit always works as an escape hatch.
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
	ForceQuit: key.NewBinding(
		key.WithKeys("ctrl+c"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.NextPanel, k.Activate, k.Filter, k.Quit}
}
