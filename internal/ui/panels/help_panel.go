package panels

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// helpKeyMap is the demo key map shown by HelpPanel. Swap it out with
// SetKeyMap for a real key.Map that implements help.KeyMap.
type helpKeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	Back   key.Binding
	Quit   key.Binding
}

func (k helpKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Select, k.Quit}
}

func (k helpKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Select, k.Back},
		{k.Quit},
	}
}

var demoHelpKeys = helpKeyMap{
	Up:     key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "up")),
	Down:   key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "down")),
	Select: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
	Back:   key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
	Quit:   key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
}

// HelpPanel shows a help panel with keyboard shortcuts.
type HelpPanel struct {
	model         help.Model
	keys          help.KeyMap
	focused       bool
	width, height int
}

func NewHelpPanel() Panel {
	return &HelpPanel{model: help.New(), keys: demoHelpKeys}
}

// SetKeyMap lets a consumer bind the panel to its own key map.
func (p *HelpPanel) SetKeyMap(keys help.KeyMap) { p.keys = keys }

func (p *HelpPanel) Init() tea.Cmd { return nil }

func (p *HelpPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "?" {
		p.model.ShowAll = !p.model.ShowAll
	}
	return p, nil
}

func (p *HelpPanel) View() string {
	p.model.Width = ContentWidth(p.width)

	return borderStyleFor(p.focused).
		Width(OuterStyleWidth(p.width)).
		Height(OuterStyleHeight(p.height)).
		Render(p.model.View(p.keys))
}

func (p *HelpPanel) SetSize(w, h int) { p.width, p.height = w, h }

func (p *HelpPanel) Focus()          { p.focused = true }
func (p *HelpPanel) Blur()           { p.focused = false }
func (p *HelpPanel) Focused() bool   { return p.focused }
func (p *HelpPanel) Focusable() bool { return true }
func (p *HelpPanel) Title() string   { return "Help" }
