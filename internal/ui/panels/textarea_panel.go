package panels

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

// TextareaPanel is a multi-line input field.
type TextareaPanel struct {
	model         textarea.Model
	focused       bool
	width, height int
}

func NewTextareaPanel() Panel {
	m := textarea.New()
	m.Placeholder = "write a description..."
	m.ShowLineNumbers = false

	return &TextareaPanel{model: m}
}

func (p *TextareaPanel) Init() tea.Cmd { return textarea.Blink }

func (p *TextareaPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	var cmd tea.Cmd
	p.model, cmd = p.model.Update(msg)
	return p, cmd
}

func (p *TextareaPanel) View() string {
	return RenderBox(p.focused, p.width, p.height, p.model.View())
}

func (p *TextareaPanel) SetSize(w, h int) {
	p.width, p.height = w, h
	p.model.SetWidth(ContentWidth(w))
	p.model.SetHeight(ContentHeight(h))
}

func (p *TextareaPanel) Focus() {
	p.focused = true
	p.model.Focus()
}

func (p *TextareaPanel) Blur() {
	p.focused = false
	p.model.Blur()
}

func (p *TextareaPanel) Focused() bool   { return p.focused }
func (p *TextareaPanel) Focusable() bool { return true }
func (p *TextareaPanel) Title() string   { return "Textarea" }

// CapturesText reports that, while focused, this panel consumes printable
// keystrokes (including "q") as text rather than as shortcuts.
func (p *TextareaPanel) CapturesText() bool { return true }
