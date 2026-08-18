package panels

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// TextinputPanel is a single-line input field.
type TextinputPanel struct {
	model         textinput.Model
	focused       bool
	width, height int
}

func NewTextinputPanel() Panel {
	m := textinput.New()
	m.Placeholder = "type something..."
	m.CharLimit = 156
	m.Prompt = "> "

	return &TextinputPanel{model: m}
}

func (p *TextinputPanel) Init() tea.Cmd { return textinput.Blink }

func (p *TextinputPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	var cmd tea.Cmd
	p.model, cmd = p.model.Update(msg)
	return p, cmd
}

func (p *TextinputPanel) View() string {
	return borderStyleFor(p.focused).
		Width(OuterStyleWidth(p.width)).
		Height(OuterStyleHeight(p.height)).
		Render(p.model.View())
}

func (p *TextinputPanel) SetSize(w, h int) {
	p.width, p.height = w, h
	p.model.Width = ContentWidth(w)
}

func (p *TextinputPanel) Focus() {
	p.focused = true
	p.model.Focus()
}

func (p *TextinputPanel) Blur() {
	p.focused = false
	p.model.Blur()
}

func (p *TextinputPanel) Focused() bool   { return p.focused }
func (p *TextinputPanel) Focusable() bool { return true }
func (p *TextinputPanel) Title() string   { return "Textinput" }
