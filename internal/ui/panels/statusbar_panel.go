package panels

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var statusbarStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("245")).
	Padding(0, 1)

// StatusbarPanel es el ejemplo de referencia para todos los bloques
// "Estructura": no implementan Focus real (Focusable() = false), así
// el ciclo de foco (Tab/Shift+Tab) los salta automáticamente.
type StatusbarPanel struct {
	width int
	text  string
}

func NewStatusbarPanel() Panel {
	return &StatusbarPanel{text: "listo"}
}

func (p *StatusbarPanel) Init() tea.Cmd { return nil }

func (p *StatusbarPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	// Sin estado propio que reaccione a mensajes por ahora.
	// El dev puede escuchar un tea.Msg custom aquí (ej. StatusChangedMsg)
	// para actualizar p.text desde internal/stubs.
	return p, nil
}

func (p *StatusbarPanel) View() string {
	return statusbarStyle.Width(p.width).Render(p.text)
}

func (p *StatusbarPanel) SetSize(w, h int) { p.width = w }
func (p *StatusbarPanel) Focus()           {}
func (p *StatusbarPanel) Blur()            {}
func (p *StatusbarPanel) Focused() bool    { return false }
func (p *StatusbarPanel) Focusable() bool  { return false }
func (p *StatusbarPanel) Title() string    { return "Statusbar" }
