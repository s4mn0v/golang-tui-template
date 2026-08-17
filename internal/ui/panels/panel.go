package panels

import tea "github.com/charmbracelet/bubbletea"

// Panel es el contrato que implementa cada bloque de la paleta
// (list, table, viewport, ..., statusbar, custom). El Model raíz
// (internal/ui/model.go) solo conoce esta interface, nunca los
// tipos concretos — así agregar un bloque nuevo no toca model.go.
type Panel interface {
	// Init dispara comandos iniciales (ej. cargar datos async).
	Init() tea.Cmd

	// Update procesa un tea.Msg y retorna el panel actualizado.
	// Retorna Panel (no *ConcretePanel) porque el struct puede
	// necesitar reemplazarse completo, no solo mutarse.
	Update(msg tea.Msg) (Panel, tea.Cmd)

	// View renderiza el panel ya con su tamaño asignado.
	View() string

	// SetSize se llama en cada tea.WindowSizeMsg o recálculo de layout.
	SetSize(width, height int)

	// Focus/Blur/Focused controlan el estado de foco visual y de input.
	Focus()
	Blur()
	Focused() bool

	// Focusable indica si el panel puede recibir foco/teclado.
	// Bloques estructurales como statusbar/titlebar devuelven false.
	Focusable() bool

	// Title identifica el panel en el help/commandbar/debug.
	Title() string
}
