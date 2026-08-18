package panels

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// ViewportPanel shows long text (logs, results, docs) with scrolling.
type ViewportPanel struct {
	model         viewport.Model
	focused       bool
	width, height int
}

func NewViewportPanel() Panel {
	m := viewport.New(0, 0)
	m.SetContent(strings.Join(demoViewportLines(), "\n"))

	return &ViewportPanel{model: m}
}

func demoViewportLines() []string {
	lines := make([]string, 0, 40)
	for i := 1; i <= 40; i++ {
		lines = append(lines, "sample log line #"+strconv.Itoa(i))
	}
	return lines
}

func (p *ViewportPanel) Init() tea.Cmd { return nil }

func (p *ViewportPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	var cmd tea.Cmd
	p.model, cmd = p.model.Update(msg)
	return p, cmd
}

func (p *ViewportPanel) View() string {
	return RenderBox(p.focused, p.width, p.height, p.model.View())
}

func (p *ViewportPanel) SetSize(w, h int) {
	p.width, p.height = w, h
	p.model.Width = ContentWidth(w)
	p.model.Height = ContentHeight(h)
}

func (p *ViewportPanel) Focus()          { p.focused = true }
func (p *ViewportPanel) Blur()           { p.focused = false }
func (p *ViewportPanel) Focused() bool   { return p.focused }
func (p *ViewportPanel) Focusable() bool { return true }
func (p *ViewportPanel) Title() string   { return "Viewport" }
