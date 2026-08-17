package panels

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type TablePanel struct {
	model         table.Model
	focused       bool
	width, height int
}

func NewTablePanel() Panel {
	columns := []table.Column{
		{Title: "ID", Width: 6},
		{Title: "Nombre", Width: 22},
		{Title: "Estado", Width: 14},
	}
	rows := []table.Row{
		{"1", "compresor-a", "activo"},
		{"2", "bomba-b", "mantenimiento"},
		{"3", "motor-c", "inactivo"},
	}

	m := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
	)

	return &TablePanel{model: m}
}

func (p *TablePanel) Init() tea.Cmd { return nil }

func (p *TablePanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	var cmd tea.Cmd
	p.model, cmd = p.model.Update(msg)
	return p, cmd
}

func (p *TablePanel) View() string {
	return borderStyleFor(p.focused).
		Width(max(0, p.width-2)).
		Height(max(0, p.height-2)).
		Render(p.model.View())
}

func (p *TablePanel) SetSize(w, h int) {
	p.width, p.height = w, h
	p.model.SetWidth(max(0, w-2))
	p.model.SetHeight(max(0, h-2))
}

func (p *TablePanel) Focus() {
	p.focused = true
	p.model.Focus()
}

func (p *TablePanel) Blur() {
	p.focused = false
	p.model.Blur()
}

func (p *TablePanel) Focused() bool   { return p.focused }
func (p *TablePanel) Focusable() bool { return true }
func (p *TablePanel) Title() string   { return "Main (table)" }
