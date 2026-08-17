package panels

import "github.com/charmbracelet/lipgloss"

var (
	focusedPanelStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62"))

	blurredPanelStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("240"))
)

// borderStyleFor evita repetir el if/else de focused en cada panel.
func borderStyleFor(focused bool) lipgloss.Style {
	if focused {
		return focusedPanelStyle
	}
	return blurredPanelStyle
}
