package panels

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var customPlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)

// CustomPanel is an empty block with no assigned component — a placeholder
// slot meant to be swapped out for bespoke logic/rendering by the consumer.
type CustomPanel struct {
	focused       bool
	width, height int
}

func NewCustomPanel() Panel {
	return &CustomPanel{}
}

func (p *CustomPanel) Init() tea.Cmd { return nil }

func (p *CustomPanel) Update(msg tea.Msg) (Panel, tea.Cmd) { return p, nil }

func (p *CustomPanel) View() string {
	content := customPlaceholderStyle.Render("custom block — no component assigned")

	return RenderBox(p.focused, p.width, p.height, content)
}

func (p *CustomPanel) SetSize(w, h int) { p.width, p.height = w, h }

func (p *CustomPanel) Focus()          { p.focused = true }
func (p *CustomPanel) Blur()           { p.focused = false }
func (p *CustomPanel) Focused() bool   { return p.focused }
func (p *CustomPanel) Focusable() bool { return true }
func (p *CustomPanel) Title() string   { return "Custom" }
