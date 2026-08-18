package ui

const (
	BreakpointNarrow = 70

	BreakpointWide = 120
)

const (
	sidebarMinWidth  = 20
	sidebarFixedWide = 32
	titlebarHeight   = 1
	statusHeight     = 1
	commandBarHeight = 1
)

type Arrangement int

const (
	ArrangeHorizontal Arrangement = iota
	ArrangeStacked
)

type LayoutResult struct {
	Arrangement Arrangement

	TitlebarW int

	SidebarW, SidebarH int
	MainW, MainH       int
	StatusW            int
}

// computeLayout reserves one full-width row at the very top for the title
// bar and one at the bottom for both the status bar and the command bar,
// before splitting whatever vertical space remains between the sidebar and
// the main panel. Reserving the title bar row keeps the panels' top border
// from ever landing on row 0: without it, a body one row taller than the
// terminal (an off-by-one during a resize) scrolls the top border out of
// view with no room to recover.
func computeLayout(width, height int) LayoutResult {
	bodyHeight := height - titlebarHeight - statusHeight - commandBarHeight
	if bodyHeight < 0 {
		bodyHeight = 0
	}

	var result LayoutResult
	if width < BreakpointNarrow {
		result = computeStackedLayout(width, bodyHeight)
	} else {
		result = computeHorizontalLayout(width, bodyHeight)
	}

	result.TitlebarW = width
	return result
}

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
