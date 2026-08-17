package ui

// Breakpoints en columnas de terminal (no píxeles). Ajustables según
// qué tan angosto/ancho quieras soportar.
const (
	// BreakpointNarrow: por debajo de este ancho, el sidebar deja de
	// tener sentido al lado del main (columnas muy angostas) — se
	// apilan verticalmente en su lugar.
	BreakpointNarrow = 70

	// BreakpointWide: por encima de este ancho, el sidebar deja de
	// crecer proporcionalmente (30% de 200 columnas es demasiado) y
	// se fija a un ancho cómodo.
	BreakpointWide = 120
)

const (
	sidebarMinWidth  = 20
	sidebarFixedWide = 32
	statusHeight     = 1
	commandBarHeight = 1
)

// Arrangement indica cómo se combinan sidebar+main en View().
type Arrangement int

const (
	ArrangeHorizontal Arrangement = iota // lado a lado (terminal normal/ancha)
	ArrangeStacked                       // uno encima del otro (terminal angosta)
)

// LayoutResult son los tamaños ya resueltos para cada región. model.go
// solo aplica estos números — todo el cálculo/breakpoints vive acá,
// separado a propósito para poder testear con go test sin un
// tea.Program real corriendo.
type LayoutResult struct {
	Arrangement Arrangement

	SidebarW, SidebarH int
	MainW, MainH       int
	StatusW            int
}

func computeLayout(width, height int) LayoutResult {
	bodyHeight := height - statusHeight - commandBarHeight
	if bodyHeight < 0 {
		bodyHeight = 0
	}

	if width < BreakpointNarrow {
		return computeStackedLayout(width, bodyHeight)
	}
	return computeHorizontalLayout(width, bodyHeight)
}

// computeStackedLayout: sidebar arriba (menos alto, es navegación),
// main abajo (más alto, es donde se trabaja).
func computeStackedLayout(width, bodyHeight int) LayoutResult {
	sidebarH := bodyHeight / 3
	mainH := bodyHeight - sidebarH

	return LayoutResult{
		Arrangement: ArrangeStacked,
		SidebarW:    width,
		SidebarH:    sidebarH,
		MainW:       width,
		MainH:       mainH,
		StatusW:     width,
	}
}

// computeHorizontalLayout: sidebar proporcional (30%) entre los
// breakpoints, con piso mínimo y techo fijo en los extremos.
func computeHorizontalLayout(width, bodyHeight int) LayoutResult {
	var sidebarW int

	switch {
	case width > BreakpointWide:
		sidebarW = sidebarFixedWide
	default:
		sidebarW = width * 3 / 10
		if sidebarW < sidebarMinWidth {
			sidebarW = sidebarMinWidth
		}
	}

	mainW := width - sidebarW
	if mainW < 0 {
		mainW = 0
	}

	return LayoutResult{
		Arrangement: ArrangeHorizontal,
		SidebarW:    sidebarW,
		SidebarH:    bodyHeight,
		MainW:       mainW,
		MainH:       bodyHeight,
		StatusW:     width,
	}
}
