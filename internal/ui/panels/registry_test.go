package panels

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
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
