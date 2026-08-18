package panels

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type listItem string

func (i listItem) FilterValue() string { return string(i) }
func (i listItem) Title() string       { return string(i) }
func (i listItem) Description() string { return "" }

type ListPanel struct {
	model         list.Model
	focused       bool
	width, height int
}

func NewListPanel() Panel {
	items := []list.Item{
		listItem("compressor-a"),
		listItem("pump-b"),
		listItem("motor-c"),
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.SetSpacing(0)

	m := list.New(items, delegate, 0, 0)
	m.Title = "List"
	m.SetShowStatusBar(false)
	m.SetShowHelp(false)

	return &ListPanel{model: m}
}

// SetItems replaces the list's items.
func (p *ListPanel) SetItems(items []string) {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = listItem(item)
	}
	p.model.SetItems(listItems)
}

// SetTitle changes the heading shown above the list.
func (p *ListPanel) SetTitle(title string) { p.model.Title = title }

// Index returns the cursor position of the currently highlighted item.
func (p *ListPanel) Index() int { return p.model.Index() }

// Filtering reports whether the list is currently capturing text for its
// fuzzy filter (and therefore shouldn't have its keystrokes intercepted).
func (p *ListPanel) Filtering() bool { return p.model.FilterState() == list.Filtering }

// CapturesText reports that, while the fuzzy filter is being typed into, this
// panel consumes printable keystrokes as text rather than as shortcuts.
func (p *ListPanel) CapturesText() bool { return p.Filtering() }

func (p *ListPanel) Init() tea.Cmd { return nil }

func (p *ListPanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	var cmd tea.Cmd
	p.model, cmd = p.model.Update(msg)
	return p, cmd
}

func (p *ListPanel) View() string {
	return RenderBox(p.focused, p.width, p.height, p.model.View())
}

func (p *ListPanel) SetSize(w, h int) {
	p.width, p.height = w, h

	p.model.SetSize(
		ContentWidth(w),
		ContentHeight(p.height),
	)
}

func (p *ListPanel) Focus()          { p.focused = true }
func (p *ListPanel) Blur()           { p.focused = false }
func (p *ListPanel) Focused() bool   { return p.focused }
func (p *ListPanel) Focusable() bool { return true }
func (p *ListPanel) Title() string   { return "List" }
