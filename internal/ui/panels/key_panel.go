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

	keyDescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245")).
			PaddingRight(2)
)

// KeyPanel is a compact key/shortcut indicator, useful for surfacing one or
// more bindings inline (e.g. next to a field or an action).
type KeyPanel struct {
	bindings      []key.Binding
	width, height int
}

func NewKeyPanel() Panel {
	return &KeyPanel{
		bindings: []key.Binding{
			key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "save")),
			key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "delete")),
		},
	}
}

// SetBindings replaces the bindings shown by the panel.
func (p *KeyPanel) SetBindings(bindings ...key.Binding) { p.bindings = bindings }

func (p *KeyPanel) Init() tea.Cmd { return nil }

func (p *KeyPanel) Update(msg tea.Msg) (Panel, tea.Cmd) { return p, nil }

func (p *KeyPanel) View() string {
	parts := make([]string, 0, len(p.bindings))
	for _, b := range p.bindings {
		h := b.Help()
		parts = append(parts, keyBadgeStyle.Render(h.Key)+" "+keyDescStyle.Render(h.Desc))
	}
	return lipgloss.NewStyle().Width(p.width).Render(strings.Join(parts, " "))
}

func (p *KeyPanel) SetSize(w, h int) { p.width, p.height = w, h }

func (p *KeyPanel) Focus()          {}
func (p *KeyPanel) Blur()           {}
func (p *KeyPanel) Focused() bool   { return false }
func (p *KeyPanel) Focusable() bool { return false }
func (p *KeyPanel) Title() string   { return "Key" }
