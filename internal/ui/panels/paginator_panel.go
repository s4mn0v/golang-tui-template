package panels

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// PaginatorPanel demos a pagination indicator/control over a set of items.
type PaginatorPanel struct {
	model         paginator.Model
	items         []string
	focused       bool
	width, height int
}

func NewPaginatorPanel() Panel {
	items := make([]string, 0, 42)
	for i := 1; i <= 42; i++ {
		items = append(items, fmt.Sprintf("item %02d", i))
	}

	m := paginator.New()
	m.Type = paginator.Dots
	m.PerPage = 7
	m.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.Color("62")).Render("●")
	m.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("○")
	m.SetTotalPages(len(items))

	return &PaginatorPanel{model: m, items: items}
}

func (p *PaginatorPanel) Init() tea.Cmd { return nil }

func (p *PaginatorPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "left", "h":
			p.model.PrevPage()
			return p, nil
		case "right", "l":
			p.model.NextPage()
			return p, nil
		}
	}

	var cmd tea.Cmd
	p.model, cmd = p.model.Update(msg)
	return p, cmd
}

func (p *PaginatorPanel) View() string {
	start, end := p.model.GetSliceBounds(len(p.items))

	var b strings.Builder
	for _, item := range p.items[start:end] {
		b.WriteString(item)
		b.WriteString("\n")
	}
	b.WriteString("\n")
	b.WriteString(p.model.View())

	return RenderBox(p.focused, p.width, p.height, b.String())
}

func (p *PaginatorPanel) SetSize(w, h int) {
	p.width, p.height = w, h

	// PerPage was fixed at 7 regardless of the box's actual height, so on a
	// short terminal the item list plus the separator and indicator lines
	// could render taller than the allocated box — lipgloss's Height() is a
	// minimum, not a clamp, so the extra rows pushed everything below (and
	// the top border, once the whole frame no longer fit the screen) out of
	// view. Size PerPage to what's actually available instead.
	perPage := ContentHeight(h) - 2 // 1 blank separator + 1 indicator line
	if perPage < 1 {
		perPage = 1
	}
	p.model.PerPage = perPage
	p.model.SetTotalPages(len(p.items))

	if p.model.Page > p.model.TotalPages-1 {
		p.model.Page = p.model.TotalPages - 1
	}
	if p.model.Page < 0 {
		p.model.Page = 0
	}
}

func (p *PaginatorPanel) Focus()          { p.focused = true }
func (p *PaginatorPanel) Blur()           { p.focused = false }
func (p *PaginatorPanel) Focused() bool   { return p.focused }
func (p *PaginatorPanel) Focusable() bool { return true }
func (p *PaginatorPanel) Title() string   { return "Paginator" }
