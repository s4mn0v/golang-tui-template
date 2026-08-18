package panels

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	keyBadgeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("0")).
			Background(lipgloss.Color("214")).
			Padding(0, 1)

	keyBadgeActiveStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("0")).
				Background(lipgloss.Color("42")).
				Padding(0, 1)

	keyDescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245")).
			PaddingRight(2)

	keyHintStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
)

// KeyPanel is a compact key/shortcut indicator: a row of key badges that,
// once focused, actually react when pressed — proof the binding underneath
// them fires, not just a static label.
type KeyPanel struct {
	bindings      []key.Binding
	triggered     int // index into bindings of the last match, -1 if none yet
	focused       bool
	width, height int
}

func NewKeyPanel() Panel {
	return &KeyPanel{
		bindings: []key.Binding{
			key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "save")),
			key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "delete")),
		},
		triggered: -1,
	}
}

// SetBindings replaces the bindings shown by the panel.
func (p *KeyPanel) SetBindings(bindings ...key.Binding) {
	p.bindings = bindings
	p.triggered = -1
}

func (p *KeyPanel) Init() tea.Cmd { return nil }

func (p *KeyPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return p, nil
	}

	for i, b := range p.bindings {
		if key.Matches(keyMsg, b) {
			p.triggered = i
			break
		}
	}

	return p, nil
}

func (p *KeyPanel) View() string {
	badges := make([]string, 0, len(p.bindings))
	for i, b := range p.bindings {
		h := b.Help()
		badgeStyle := keyBadgeStyle
		if i == p.triggered {
			badgeStyle = keyBadgeActiveStyle
		}
		badges = append(badges, badgeStyle.Render(h.Key)+" "+keyDescStyle.Render(h.Desc))
	}

	content := strings.Join(badges, " ") + "\n\n"
	if p.triggered >= 0 {
		h := p.bindings[p.triggered].Help()
		content += keyHintStyle.Render("triggered: " + h.Key + " → " + h.Desc)
	} else {
		content += keyHintStyle.Render("focus this panel and press a key above to try it")
	}

	return RenderBox(p.focused, p.width, p.height, content)
}

func (p *KeyPanel) SetSize(w, h int) { p.width, p.height = w, h }

func (p *KeyPanel) Focus()          { p.focused = true }
func (p *KeyPanel) Blur()           { p.focused = false }
func (p *KeyPanel) Focused() bool   { return p.focused }
func (p *KeyPanel) Focusable() bool { return true }
func (p *KeyPanel) Title() string   { return "Key" }
