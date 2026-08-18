package panels

import (
	"strings"
	"testing"
)

// TestRenderBoxKeepsBothBordersOnOverflow is the regression test for the
// bug this file exists to prevent: content taller than its box used to
// get truncated from the bottom of the *whole rendered block* (border
// included), which kept the top border intact but cut the bottom one off.
// fitToHeight now clips the content itself, before the border is drawn, so
// both borders should always be present regardless of how much content
// overflows.
func TestRenderBoxKeepsBothBordersOnOverflow(t *testing.T) {
	tall := strings.Repeat("line\n", 50) // way taller than any box below

	// 5 is this chrome's floor: border (2 rows) + padding (2 rows) + 1 row
	// of actual content. Anything shorter can't fit a top border, a
	// bottom border, *and* the padding this style always adds, no matter
	// how little content there is — a box that small isn't a realistic
	// panel size to defend, just an unrenderable one.
	for _, height := range []int{5, 6, 10, 30} {
		out := RenderBox(false, 24, height, tall)
		lines := strings.Split(out, "\n")

		if len(lines) != height {
			t.Errorf("height=%d: got %d lines, want exactly %d", height, len(lines), height)
			continue
		}
		if !strings.Contains(lines[0], "╭") {
			t.Errorf("height=%d: top border missing, first line was %q", height, lines[0])
		}
		if !strings.Contains(lines[len(lines)-1], "╰") {
			t.Errorf("height=%d: bottom border missing, last line was %q", height, lines[len(lines)-1])
		}
	}
}

// TestRenderBoxShowsScrollbarOnOverflow checks that clipped content gets a
// scrollbar painted into it, so the cut reads as "scroll for more" instead
// of looking like broken output.
func TestRenderBoxShowsScrollbarOnOverflow(t *testing.T) {
	tall := strings.Repeat("line\n", 50)
	out := RenderBox(false, 24, 8, tall)

	// scrollThumbChar ("█") is checked rather than scrollTrackChar ("│"):
	// the box's own left/right border already uses "│", so that character
	// alone would be a false positive.
	if !strings.Contains(out, scrollThumbChar) {
		t.Errorf("expected a scrollbar thumb (%q) in overflowed output, got:\n%s", scrollThumbChar, out)
	}
}

// TestRenderBoxLeavesShortContentUntouched checks that content which
// already fits isn't clipped or given a scrollbar it doesn't need.
func TestRenderBoxLeavesShortContentUntouched(t *testing.T) {
	out := RenderBox(false, 24, 10, "hello")

	if strings.Contains(out, scrollThumbChar) {
		t.Errorf("content that fits should not get a scrollbar, got:\n%s", out)
	}
}
