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
		{Title: "Name", Width: 22},
		{Title: "Status", Width: 14},
	}
	rows := []table.Row{
		{"1", "compressor-a", "active"},
		{"2", "pump-b", "maintenance"},
		{"3", "motor-c", "inactive"},
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
		Width(OuterStyleWidth(p.width)).
		Height(OuterStyleHeight(p.height)).
		Render(p.model.View())
}

func (p *TablePanel) SetSize(w, h int) {
	p.width, p.height = w, h
	p.model.SetWidth(ContentWidth(w))
	p.model.SetHeight(ContentHeight(h))
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
func (p *TablePanel) Title() string   { return "Table" }
