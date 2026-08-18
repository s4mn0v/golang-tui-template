package panels

import (
	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

// FilepickerPanel lets the user browse and select a file or directory.
type FilepickerPanel struct {
	model         filepicker.Model
	focused       bool
	selectedPath  string
	width, height int
}

func NewFilepickerPanel() Panel {
	m := filepicker.New()
	m.CurrentDirectory = "."
	m.AutoHeight = false

	return &FilepickerPanel{model: m}
}

func (p *FilepickerPanel) Init() tea.Cmd { return p.model.Init() }

func (p *FilepickerPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	var cmd tea.Cmd
	p.model, cmd = p.model.Update(msg)

	if didSelect, path := p.model.DidSelectFile(msg); didSelect {
		p.selectedPath = path
	}

	return p, cmd
}

func (p *FilepickerPanel) View() string {
	content := p.model.View()
	if p.selectedPath != "" {
		content += "\n\nselected: " + p.selectedPath
	}

	return borderStyleFor(p.focused).
		Width(OuterStyleWidth(p.width)).
		Height(OuterStyleHeight(p.height)).
		Render(content)
}

func (p *FilepickerPanel) SetSize(w, h int) {
	p.width, p.height = w, h
	p.model.SetHeight(ContentHeight(h))
}

func (p *FilepickerPanel) Focus()          { p.focused = true }
func (p *FilepickerPanel) Blur()           { p.focused = false }
func (p *FilepickerPanel) Focused() bool   { return p.focused }
func (p *FilepickerPanel) Focusable() bool { return true }
func (p *FilepickerPanel) Title() string   { return "Filepicker" }
