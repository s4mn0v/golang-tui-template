package panels

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var titlebarStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("230")).
	Background(lipgloss.Color("62")).
	Padding(0, 1)

// TitlebarPanel is a header bar, typically holding the app or view title.
type TitlebarPanel struct {
	width int
	title string
}

func NewTitlebarPanel() Panel {
	return &TitlebarPanel{title: "tui-template"}
}

// SetTitle changes the text rendered in the title bar.
func (p *TitlebarPanel) SetTitle(title string) { p.title = title }

func (p *TitlebarPanel) Init() tea.Cmd { return nil }

func (p *TitlebarPanel) Update(msg tea.Msg) (Panel, tea.Cmd) { return p, nil }

func (p *TitlebarPanel) View() string {
	return titlebarStyle.Width(p.width).Render(p.title)
}

func (p *TitlebarPanel) SetSize(w, h int) { p.width = w }

func (p *TitlebarPanel) Focus()          {}
func (p *TitlebarPanel) Blur()           {}
func (p *TitlebarPanel) Focused() bool   { return false }
func (p *TitlebarPanel) Focusable() bool { return false }
func (p *TitlebarPanel) Title() string   { return "Titlebar" }
