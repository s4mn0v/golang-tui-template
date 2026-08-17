package panels

import "github.com/charmbracelet/lipgloss"

var (
	focusedPanelStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62")).
				Padding(1, 2)

	blurredPanelStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("240")).
				Padding(1, 2)
)

func borderStyleFor(focused bool) lipgloss.Style {
	if focused {
		return focusedPanelStyle
	}
	return blurredPanelStyle
}

const (
	borderSize = 2
	paddingV   = 2
	paddingH   = 4
)

func ContentWidth(totalWidth int) int {
	if w := totalWidth - borderSize - paddingH; w > 0 {
		return w
	}
	return 0
}

func ContentHeight(totalHeight int) int {
	if h := totalHeight - borderSize - paddingV; h > 0 {
		return h
	}
	return 0
}

func OuterStyleWidth(totalWidth int) int {
	if w := totalWidth - borderSize; w > 0 {
		return w
	}
	return 0
}

func OuterStyleHeight(totalHeight int) int {
	if h := totalHeight - borderSize; h > 0 {
		return h
	}
	return 0
}
