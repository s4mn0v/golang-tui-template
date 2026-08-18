package modals

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SelectorResultMsg is emitted when a Selector modal is completed or
// cancelled.
type SelectorResultMsg struct {
	Cancelled bool
	Value     string
}

type selectorItem string

func (i selectorItem) FilterValue() string { return string(i) }
func (i selectorItem) Title() string       { return string(i) }
func (i selectorItem) Description() string { return "" }

// Selector is a modal picker with fuzzy search over a fixed list of options.
type Selector struct {
	title string
	model list.Model
	done  bool

	width, height int
}

// NewSelector builds a selector modal over the given options.
func NewSelector(title string, options ...string) *Selector {
	items := make([]list.Item, 0, len(options))
	for _, o := range options {
		items = append(items, selectorItem(o))
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.SetSpacing(0)

	m := list.New(items, delegate, 40, 12)
	m.Title = title
	m.SetShowStatusBar(false)
	m.SetShowHelp(false)
	m.SetFilteringEnabled(true)

	return &Selector{title: title, model: m}
}

func (s *Selector) Init() tea.Cmd { return nil }

func (s *Selector) Update(msg tea.Msg) (Modal, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok && s.model.FilterState() != list.Filtering {
		switch keyMsg.String() {
		case "enter":
			s.done = true
			value := ""
			if item, ok := s.model.SelectedItem().(selectorItem); ok {
				value = string(item)
			}
			return s, func() tea.Msg { return SelectorResultMsg{Value: value} }

		case "esc", "q":
			s.done = true
			return s, func() tea.Msg { return SelectorResultMsg{Cancelled: true} }
		}
	}

	var cmd tea.Cmd
	s.model, cmd = s.model.Update(msg)
	return s, cmd
}

func (s *Selector) Done() bool { return s.done }

// SetSize fits the option list within the given screen size. It was
// previously built at a fixed 40x12 regardless of the terminal — on a
// short window that could render taller than the screen, and just like a
// Panel whose content ignores its allocated height, the overflow pushed
// the modal's own top border out of view. The list already clips itself
// exactly to whatever height it's given (see bubbles/list), so this only
// needs to pick a height that both fits the screen and isn't larger than
// the option list actually needs — a compact popup, not a full-screen one.
func (s *Selector) SetSize(width, height int) {
	s.width, s.height = width, height

	const (
		borderRows = 2 // top + bottom border
		borderCols = 2 // left + right border
		vPadding   = 2 // Padding(1, 2): 1 top + 1 bottom
		hPadding   = 4 // Padding(1, 2): 2 left + 2 right
		titleRows  = 2 // the list's own title row + pagination row
	)

	listHeight := height - borderRows - vPadding
	if listHeight < 3 {
		listHeight = 3
	}
	if needed := len(s.model.Items()) + titleRows; listHeight > needed {
		listHeight = needed
	}

	listWidth := width - borderCols - hPadding
	if listWidth > 40 {
		listWidth = 40
	}
	if listWidth < 20 {
		listWidth = 20
	}

	s.model.SetSize(listWidth, listHeight)
}

func (s *Selector) View() string {
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2)
	// The inner list is already sized to fit exactly (see SetSize above);
	// this is the same universal safety net every other modal gets, in
	// case that math is ever off rather than something it depends on.
	box = ClampToScreen(box, s.width, s.height)

	return box.Render(s.model.View())
}
