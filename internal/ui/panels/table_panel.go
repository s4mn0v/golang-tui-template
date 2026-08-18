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
		{"4", "generator-d", "active"},
		{"5", "compressor-e", "active"},
		{"6", "pump-f", "maintenance"},
		{"7", "motor-g", "inactive"},
		{"8", "generator-h", "active"},
		{"9", "compressor-i", "maintenance"},
		{"10", "pump-j", "active"},
		{"11", "motor-k", "inactive"},
		{"12", "generator-l", "active"},
		{"13", "compressor-m", "maintenance"},
		{"14", "pump-n", "active"},
		{"15", "motor-o", "inactive"},
		{"16", "generator-p", "maintenance"},
		{"17", "compressor-q", "active"},
		{"18", "pump-r", "inactive"},
		{"19", "motor-s", "active"},
		{"20", "generator-t", "maintenance"},
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
	return RenderBox(p.focused, p.width, p.height, p.model.View())
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
