package panels

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	tabActiveStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("230")).
			Background(lipgloss.Color("62")).
			Padding(0, 2)

	tabInactiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("245")).
				Padding(0, 2)
)

// TabsPanel is a tabbed navigation block: left/right (or h/l) switch tabs.
type TabsPanel struct {
	labels        []string
	contents      []string
	active        int
	focused       bool
	width, height int
}

func NewTabsPanel() Panel {
	return &TabsPanel{
		labels: []string{"General", "Details", "History"},
		contents: []string{
			"content for the General tab.",
			"content for the Details tab.",
			"content for the History tab.",
		},
	}
}

func (p *TabsPanel) Init() tea.Cmd { return nil }

func (p *TabsPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return p, nil
	}

	switch keyMsg.String() {
	case "left", "h", "shift+tab":
		p.active = (p.active - 1 + len(p.labels)) % len(p.labels)
	case "right", "l", "tab":
		p.active = (p.active + 1) % len(p.labels)
	}

	return p, nil
}

func (p *TabsPanel) View() string {
	tabs := make([]string, 0, len(p.labels))
	for i, label := range p.labels {
		if i == p.active {
			tabs = append(tabs, tabActiveStyle.Render(label))
		} else {
			tabs = append(tabs, tabInactiveStyle.Render(label))
		}
	}
	header := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)

	var content string
	if p.active < len(p.contents) {
		content = p.contents[p.active]
	}

	body := header + "\n" + strings.Repeat("─", ContentWidth(p.width)) + "\n\n" + content

	return RenderBox(p.focused, p.width, p.height, body)
}

func (p *TabsPanel) SetSize(w, h int) { p.width, p.height = w, h }

func (p *TabsPanel) Focus()          { p.focused = true }
func (p *TabsPanel) Blur()           { p.focused = false }
func (p *TabsPanel) Focused() bool   { return p.focused }
func (p *TabsPanel) Focusable() bool { return true }
func (p *TabsPanel) Title() string   { return "Tabs" }
