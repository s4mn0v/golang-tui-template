package modals

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AlertVariant string

const (
	AlertInfo    AlertVariant = "info"
	AlertError   AlertVariant = "error"
	AlertSuccess AlertVariant = "success"
)

type Alert struct {
	Title   string
	Message string
	Variant AlertVariant
	done    bool
}

func NewAlert(title, message string, variant AlertVariant) *Alert {
	return &Alert{Title: title, Message: message, Variant: variant}
}

func (a *Alert) Init() tea.Cmd { return nil }

func (a *Alert) Update(msg tea.Msg) (Modal, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter", "esc", "q":
			a.done = true
		}
	}
	return a, nil
}

func (a *Alert) Done() bool { return a.done }

func (a *Alert) View() string {
	color := alertColor(a.Variant)

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(color).
		Padding(1, 2).
		Width(44)

	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(color)
	hint := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render("enter / esc para cerrar")

	content := titleStyle.Render(a.Title) + "\n\n" + a.Message + "\n\n" + hint
	return box.Render(content)
}

func alertColor(variant AlertVariant) lipgloss.Color {
	switch variant {
	case AlertError:
		return lipgloss.Color("196")
	case AlertSuccess:
		return lipgloss.Color("42")
	default: // AlertInfo
		return lipgloss.Color("62")
	}
}
