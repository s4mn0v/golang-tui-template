package ui

const (
	BreakpointNarrow = 70

	BreakpointWide = 120
)

const (
	sidebarMinWidth  = 20
	sidebarFixedWide = 32
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
