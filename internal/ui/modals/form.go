package modals

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// FormResultMsg is emitted when a Form modal is completed or aborted.
type FormResultMsg struct {
	Aborted bool
	Values  map[string]string
}

// Form is a modal wrapping a huh.Form for structured data entry.
type Form struct {
	title string
	form  *huh.Form

	keys []string
	vals []*string

	done bool
}

// NewForm builds a form modal from a set of (key, label) fields, each
// rendered as a single-line text input.
func NewForm(title string, fields ...[2]string) *Form {
	keys := make([]string, 0, len(fields))
	vals := make([]*string, 0, len(fields))
	groupFields := make([]huh.Field, 0, len(fields))

	for _, f := range fields {
		key, label := f[0], f[1]
		val := new(string)

		keys = append(keys, key)
		vals = append(vals, val)
		groupFields = append(groupFields, huh.NewInput().Title(label).Value(val))
	}

	form := huh.NewForm(huh.NewGroup(groupFields...)).
		WithShowHelp(true).
		WithWidth(44)

	return &Form{title: title, form: form, keys: keys, vals: vals}
}

func (f *Form) values() map[string]string {
	out := make(map[string]string, len(f.keys))
	for i, key := range f.keys {
		out[key] = *f.vals[i]
	}
	return out
}

func (f *Form) Init() tea.Cmd { return f.form.Init() }

func (f *Form) Update(msg tea.Msg) (Modal, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "esc" {
		f.done = true
		return f, func() tea.Msg { return FormResultMsg{Aborted: true} }
	}

	updated, cmd := f.form.Update(msg)
	if form, ok := updated.(*huh.Form); ok {
		f.form = form
	}

	switch f.form.State {
	case huh.StateCompleted:
		f.done = true
		return f, tea.Batch(cmd, func() tea.Msg {
			return FormResultMsg{Values: f.values()}
		})
	case huh.StateAborted:
		f.done = true
		return f, tea.Batch(cmd, func() tea.Msg {
			return FormResultMsg{Aborted: true}
		})
	}

	return f, cmd
}

func (f *Form) Done() bool { return f.done }

func (f *Form) View() string {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("62")).MarginBottom(1)

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2)

	return box.Render(titleStyle.Render(f.title) + "\n" + f.form.View())
}
