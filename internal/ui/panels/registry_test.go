package panels

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TestBlockLifecycle exercises every registered block through the standard
// Panel lifecycle (Init -> SetSize -> Focus/Blur -> Update -> View) to catch
// panics or nil-pointer bugs introduced when adding a new block type.
func TestBlockLifecycle(t *testing.T) {
	for _, blockType := range Registered() {
		t.Run(string(blockType), func(t *testing.T) {
			p, ok := New(blockType)
			if !ok {
				t.Fatalf("New(%s) returned ok=false", blockType)
			}

			p.Init()

			p.SetSize(80, 24)
			_ = p.View()

			if p.Focusable() {
				p.Focus()
				if !p.Focused() {
					t.Errorf("%s: Focus() did not set Focused() to true", blockType)
				}
				p.Blur()
				if p.Focused() {
					t.Errorf("%s: Blur() did not set Focused() to false", blockType)
				}
			}

			if _, cmd := p.Update(tea.KeyMsg{Type: tea.KeyDown}); cmd != nil {
				cmd()
			}
			_ = p.View()

			// Degenerate sizes shouldn't panic either.
			p.SetSize(0, 0)
			_ = p.View()

			if p.Title() == "" {
				t.Errorf("%s: Title() is empty", blockType)
			}
		})
	}
}

// TestViewNeverExceedsAllocatedSize renders every registered block into a
// deliberately small box and checks the result never has more lines than
// the height it was given, nor a line wider than the width — regardless of
// how much content the block itself wants to show. This is the guarantee
// RenderBox (see style.go) exists to provide universally, instead of each
// block having to size its own content exactly right; this test is what
// keeps that guarantee from silently regressing for any block, present or
// future.
func TestViewNeverExceedsAllocatedSize(t *testing.T) {
	const width, height = 22, 6

	for _, blockType := range Registered() {
		t.Run(string(blockType), func(t *testing.T) {
			p, ok := New(blockType)
			if !ok {
				t.Fatalf("New(%s) returned ok=false", blockType)
			}

			p.Init()
			p.SetSize(width, height)

			lines := strings.Split(p.View(), "\n")
			if len(lines) > height {
				t.Errorf("%s: View() produced %d lines for height=%d", blockType, len(lines), height)
			}
			for i, line := range lines {
				if w := lipgloss.Width(line); w > width {
					t.Errorf("%s: line %d is %d cells wide for width=%d", blockType, i, w, width)
				}
			}
		})
	}
}

func TestRegisteredCoversAllBlockTypes(t *testing.T) {
	want := []BlockType{
		BlockList, BlockTable, BlockViewport, BlockPaginator,
		BlockTextinput, BlockTextarea, BlockFilepicker,
		BlockSpinner, BlockProgress, BlockHelp, BlockKey,
		BlockStatusbar, BlockTitlebar, BlockTabs, BlockTree,
		BlockCustom,
	}

	got := map[BlockType]bool{}
	for _, bt := range Registered() {
		got[bt] = true
	}

	for _, bt := range want {
		if !got[bt] {
			t.Errorf("block type %s is not registered", bt)
		}
	}
}
