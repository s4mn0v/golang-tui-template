package panels

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// listItem implementa list.Item. En la app real, el dev reemplaza
// esto por su propio tipo (ej. un registro de equipo, un archivo, etc.)
// vía internal/stubs — este es solo el bloque estructural.
type listItem string

func (i listItem) FilterValue() string { return string(i) }
func (i listItem) Title() string       { return string(i) }
func (i listItem) Description() string { return "" }

type ListPanel struct {
	model         list.Model
	focused       bool
	width, height int
}

func NewListPanel() Panel {
	items := []list.Item{
		listItem("equipo-compresor-a"),
		listItem("equipo-bomba-b"),
		listItem("equipo-motor-c"),
	}
	delegate := list.NewDefaultDelegate()
	m := list.New(items, delegate, 0, 0)
	m.Title = "Sidebar"
	m.SetShowStatusBar(false)
	m.SetShowHelp(false)

	return &ListPanel{model: m}
}

func (p *ListPanel) Init() tea.Cmd { return nil }

func (p *ListPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	var cmd tea.Cmd
	p.model, cmd = p.model.Update(msg)
	return p, cmd
}

func (p *ListPanel) View() string {
	return borderStyleFor(p.focused).
		Width(max(0, p.width-2)).
		Height(max(0, p.height-2)).
		Render(p.model.View())
}

func (p *ListPanel) SetSize(w, h int) {
	p.width, p.height = w, h
	// -2 en cada eje para descontar el borde que agrega el estilo.
	p.model.SetSize(max(0, w-2), max(0, h-2))
}

func (p *ListPanel) Focus()          { p.focused = true }
func (p *ListPanel) Blur()           { p.focused = false }
func (p *ListPanel) Focused() bool   { return p.focused }
func (p *ListPanel) Focusable() bool { return true }
func (p *ListPanel) Title() string   { return "Sidebar (list)" }
