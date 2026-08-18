package panels

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SpinnerPanel shows a loading indicator.
type SpinnerPanel struct {
	model         spinner.Model
	label         string
	width, height int
}

func NewSpinnerPanel() Panel {
	m := spinner.New()
	m.Spinner = spinner.Dot
	m.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("62"))

	return &SpinnerPanel{model: m, label: "loading..."}
}

func (p *SpinnerPanel) Init() tea.Cmd { return p.model.Tick }

func (p *SpinnerPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	var cmd tea.Cmd
	p.model, cmd = p.model.Update(msg)
	return p, cmd
}

func (p *SpinnerPanel) View() string {
	content := p.model.View() + " " + p.label

	return RenderBox(false, p.width, p.height, content)
}

func (p *SpinnerPanel) SetSize(w, h int) { p.width, p.height = w, h }

func (p *SpinnerPanel) Focus()          {}
func (p *SpinnerPanel) Blur()           {}
func (p *SpinnerPanel) Focused() bool   { return false }
func (p *SpinnerPanel) Focusable() bool { return false }
func (p *SpinnerPanel) Title() string   { return "Spinner" }
