package ui

import "github.com/charmbracelet/bubbles/key"

// KeyMap agrupa los bindings globales — los que el Model raíz
// intercepta antes de delegar al panel con foco. Cuando conectemos
// el generador, esta struct se arma dinámicamente desde los nodos
// tipo Trigger del schema; por ahora está a mano para validar el
// esqueleto de orquestación y el sistema de modales.
type KeyMap struct {
	NextPanel key.Binding
	PrevPanel key.Binding
	Quit      key.Binding

	// Demo — se usan mientras no exista el registry de nodos Trigger
	// generado desde el schema. Sirven para validar que el sistema
	// de modales completo (abrir → responder → encadenar) funciona.
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

// ShortHelp es lo que renderiza la commandbar inferior — la lista
// de comandos "de guía" que pediste, siempre visible.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.NextPanel, k.PrevPanel, k.OpenConfirmDemo, k.OpenAlertDemo, k.Quit}
}
