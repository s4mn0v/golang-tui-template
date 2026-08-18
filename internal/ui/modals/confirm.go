package modals

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ConfirmResultMsg struct {
	Confirmed bool
}

type Confirm struct {
	Title        string
	Message      string
	ConfirmLabel string
	CancelLabel  string

	selected int
	done     bool
}

func NewConfirm(title, message, confirmLabel, cancelLabel string) *Confirm {
	if confirmLabel == "" {
		confirmLabel = "Yes"
	}
	if cancelLabel == "" {
		cancelLabel = "No"
	}
	return &Confirm{
		Title:        title,
		Message:      message,
		ConfirmLabel: confirmLabel,
		CancelLabel:  cancelLabel,
	}
}

func (c *Confirm) Init() tea.Cmd { return nil }

func (c *Confirm) Update(msg tea.Msg) (Modal, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return c, nil
	}

	switch keyMsg.String() {
	case "left", "right", "h", "l", "tab":
		c.selected = (c.selected + 1) % 2
		return c, nil

	case "enter":
		c.done = true
		confirmed := c.selected == 0
		return c, func() tea.Msg { return ConfirmResultMsg{Confirmed: confirmed} }

	case "esc", "q":
		c.done = true
		return c, func() tea.Msg { return ConfirmResultMsg{Confirmed: false} }
	}

	return c, nil
}

func (c *Confirm) Done() bool { return c.done }

func (c *Confirm) View() string {
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("214")).
		Padding(1, 2).
		Width(44)

	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("214"))

	confirmBtn := renderButton(c.ConfirmLabel, c.selected == 0)
	cancelBtn := renderButton(c.CancelLabel, c.selected == 1)
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, confirmBtn, "   ", cancelBtn)

	content := titleStyle.Render(c.Title) + "\n\n" + c.Message + "\n\n" + buttons
	return box.Render(content)
}

func renderButton(label string, active bool) string {
	style := lipgloss.NewStyle().Padding(0, 2)
	if active {
		style = style.Background(lipgloss.Color("214")).Foreground(lipgloss.Color("0")).Bold(true)
	} else {
		style = style.Foreground(lipgloss.Color("245"))
	}
	return style.Render(label)
}
