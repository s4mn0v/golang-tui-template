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

// borderStyleFor evita repetir el if/else de focused en cada panel.
func borderStyleFor(focused bool) lipgloss.Style {
	if focused {
		return focusedPanelStyle
	}
	return blurredPanelStyle
}

// Presupuesto de espacio que cada panel debe restar del ancho/alto
// que le asigna el layout, para que el contenido interno no quede
// pegado al borde. Centralizado acá para no repetir constantes
// mágicas (-2, -4, -6...) sueltas en cada archivo de panel.
const (
	borderSize = 2 // 1 col/fila por lado (RoundedBorder)
	paddingV   = 2 // debe coincidir con Padding(1, 2): 1 arriba + 1 abajo
	paddingH   = 4 // debe coincidir con Padding(1, 2): 2 izq + 2 der
)

// ContentWidth/ContentHeight: lo que le queda disponible al
// componente interno (list/table/etc.) después de restar borde+padding.
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

// OuterStyleWidth/Height: lo que se pasa a Style.Width()/Height().
// En lipgloss, Width/Height del Style incluyen el padding pero NO
// el borde — por eso acá solo se resta borderSize, no el padding.
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
