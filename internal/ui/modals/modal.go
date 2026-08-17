package modals

import tea "github.com/charmbracelet/bubbletea"

// Modal es el contrato que implementan los overlays (alert, confirm,
// form, selector). A diferencia de panels.Panel, un Modal no ocupa
// espacio en el grid — se renderiza flotando encima del body ya
// compuesto, y mientras está activo, captura TODOS los mensajes
// (el Model raíz deja de delegar al panel con foco).
type Modal interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Modal, tea.Cmd)
	View() string

	// Done indica que el modal terminó su ciclo (el usuario ya
	// respondió o canceló) y el Model raíz debe cerrarlo.
	Done() bool
}
