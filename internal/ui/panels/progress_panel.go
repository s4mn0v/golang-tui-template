package panels

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

// progressTickMsg drives the demo progress bar forward.
type progressTickMsg struct{}

func progressTick() tea.Cmd {
	return tea.Tick(500*time.Millisecond, func(time.Time) tea.Msg {
		return progressTickMsg{}
	})
}

// ProgressPanel shows a progress bar.
type ProgressPanel struct {
	model         progress.Model
	percent       float64
	width, height int
}

func NewProgressPanel() Panel {
	m := progress.New(progress.WithDefaultGradient())
	return &ProgressPanel{model: m}
}

func (p *ProgressPanel) Init() tea.Cmd { return progressTick() }

func (p *ProgressPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	switch msg.(type) {
	case progressTickMsg:
		p.percent += 0.05
		if p.percent > 1 {
			p.percent = 0
		}
		cmd := p.model.SetPercent(p.percent)
		return p, tea.Batch(cmd, progressTick())
	}

	updated, cmd := p.model.Update(msg)
	if m, ok := updated.(progress.Model); ok {
		p.model = m
	}
	return p, cmd
}

func (p *ProgressPanel) View() string {
	return borderStyleFor(false).
		Width(OuterStyleWidth(p.width)).
		Height(OuterStyleHeight(p.height)).
		Render(p.model.View())
}

func (p *ProgressPanel) SetSize(w, h int) {
	p.width, p.height = w, h
	p.model.Width = ContentWidth(w)
}

func (p *ProgressPanel) Focus()          {}
func (p *ProgressPanel) Blur()           {}
func (p *ProgressPanel) Focused() bool   { return false }
func (p *ProgressPanel) Focusable() bool { return false }
func (p *ProgressPanel) Title() string   { return "Progress" }
