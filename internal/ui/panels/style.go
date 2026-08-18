package panels

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

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

// RenderBox renders content inside the standard panel chrome — border plus
// padding, sized to exactly fill width x height — the way every bordered
// panel in this package does.
//
// Content taller than the box is fitted rather than left to overflow: a
// panel whose content can grow past its allocated height — more rows than
// fit, a status line appended after content that already filled its box,
// ... — would otherwise render a block taller than height, which pushes
// its own bottom border down past where it belongs (and, once the whole
// frame no longer fits the terminal, scrolls the top border out of view
// too). fitToHeight below clips to the box before the border is drawn, so
// both borders always land exactly where the given height says they
// should, and paints a scrollbar in the last column when it had to cut
// anything — the overflow reads as "scroll for more", not a broken box.
//
// Every panel gets this for free just by rendering through here, instead
// of each one having to reason about its own worst case.
func RenderBox(focused bool, width, height int, content string) string {
	content = fitToHeight(content, ContentWidth(width), ContentHeight(height))

	return borderStyleFor(focused).
		Width(OuterStyleWidth(width)).
		Height(OuterStyleHeight(height)).
		MaxWidth(width).
		MaxHeight(height).
		Render(content)
}

var (
	scrollTrackStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	scrollThumbStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
)

const (
	scrollTrackChar = "│"
	scrollThumbChar = "█"
)

// fitToHeight clips content to at most innerH lines, keeping the tail —
// the newest or most relevant part is usually at the end (a "selected: "
// line appended after a listing, the latest entry in a log) — and, when it
// had to cut anything, paints a scrollbar in the last column so what's
// missing is visibly "scroll up for more" rather than a border that just
// looks wrong. Width is only used to line the scrollbar up in a straight
// column; over-wide lines are left to the caller's own MaxWidth safety net.
func fitToHeight(content string, innerW, innerH int) string {
	if innerH <= 0 {
		return ""
	}

	lines := strings.Split(content, "\n")
	if len(lines) <= innerH {
		return content
	}

	visible := lines[len(lines)-innerH:]
	track := scrollbarColumn(len(lines), innerH)

	textWidth := innerW - 1 // reserve the last column for the scrollbar
	if textWidth < 0 {
		textWidth = 0
	}
	padStyle := lipgloss.NewStyle().Width(textWidth)

	for i, line := range visible {
		visible[i] = padStyle.Render(line) + track[i]
	}

	return strings.Join(visible, "\n")
}

// scrollbarColumn builds a `visible`-row scrollbar for content that's
// `total` rows tall when only the tail is shown. The thumb sits at the
// bottom of the track, sized proportionally to how much is visible, to
// signal "you're at the end; there's more above."
func scrollbarColumn(total, visible int) []string {
	col := make([]string, visible)

	thumbSize := visible * visible / total
	if thumbSize < 1 {
		thumbSize = 1
	}
	if thumbSize > visible {
		thumbSize = visible
	}
	thumbStart := visible - thumbSize

	for i := range col {
		if i >= thumbStart {
			col[i] = scrollThumbStyle.Render(scrollThumbChar)
		} else {
			col[i] = scrollTrackStyle.Render(scrollTrackChar)
		}
	}

	return col
}
