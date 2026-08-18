package modals

import "github.com/charmbracelet/lipgloss"

// ClampToScreen applies a safety-net max width/height to an already-styled
// modal box (border, padding, and whatever else is already on the style) so
// content taller or wider than the screen is clipped instead of overflowing
// — the same problem, and the same fix, as panels.RenderBox. Width()/
// Height() are a floor, not a ceiling: a modal whose content can grow (more
// list items, a longer message, ...) can otherwise render larger than the
// screen and, once the frame no longer fits the terminal, push its own top
// border out of view.
//
// width/height are the full screen size a modal receives via SetSize, not
// a pre-shrunk box, so every modal gets this for free just by rendering its
// final box through here — no per-modal reasoning about whether its content
// could ever grow past a comfortable size required. 0 (SetSize never
// called — e.g. in a test rendering a modal directly) leaves the
// corresponding dimension unclamped rather than collapsing it to nothing.
func ClampToScreen(style lipgloss.Style, width, height int) lipgloss.Style {
	if width > 0 {
		style = style.MaxWidth(width)
	}
	if height > 0 {
		style = style.MaxHeight(height)
	}
	return style
}
