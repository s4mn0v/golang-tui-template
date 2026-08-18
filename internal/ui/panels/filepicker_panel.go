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

	return RenderBox(p.focused, p.width, p.height, content)
}

func (p *FilepickerPanel) SetSize(w, h int) {
	p.width, p.height = w, h

	// The picker pads its own View() out to one *more* than the height
	// it's given (bubbles/filepicker's own convention), and View() below
	// then appends the "selected: <path>" line after that — so without
	// reserving room up front, that's always 3 rows more than the box
	// actually has. RenderBox's scrollbar safety net keeps that from ever
	// breaking a border, but there's no reason to make it earn its keep
	// on every ordinary-sized terminal: reserve the space so the listing
	// and the selection line actually fit together.
	const (
		pickerHeightQuirk = 1 // bubbles/filepicker pads to m.Height+1, not m.Height
		selectedLines     = 2 // blank separator + the "selected: ..." line
	)

	pickerHeight := ContentHeight(h) - pickerHeightQuirk - selectedLines
	if pickerHeight < 1 {
		pickerHeight = 1
	}
	p.model.SetHeight(pickerHeight)
}

func (p *FilepickerPanel) Focus()          { p.focused = true }
func (p *FilepickerPanel) Blur()           { p.focused = false }
func (p *FilepickerPanel) Focused() bool   { return p.focused }
func (p *FilepickerPanel) Focusable() bool { return true }
func (p *FilepickerPanel) Title() string   { return "Filepicker" }
