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

func (s *Selector) View() string {
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2)

	return box.Render(s.model.View())
}
