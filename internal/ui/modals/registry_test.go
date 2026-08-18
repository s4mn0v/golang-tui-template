package modals

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

// TestModalRegistry exercises every registered modal through New to catch
// panics or nil-pointer bugs introduced when adding a new modal type.
func TestModalRegistry(t *testing.T) {
	for _, modalType := range Registered() {
		t.Run(string(modalType), func(t *testing.T) {
			m, ok := New(modalType)
			if !ok {
				t.Fatalf("New(%s) returned ok=false", modalType)
			}

			m.Init()
			if m.View() == "" {
				t.Errorf("%s: View() returned empty string", modalType)
			}
			if m.Done() {
				t.Errorf("%s: modal reports Done() immediately after construction", modalType)
			}
		})
	}
}

// TestViewNeverExceedsAllocatedSize renders every registered modal into a
// deliberately small screen and checks the result never has more lines than
// the height it was given via SetSize, nor a line wider than the width —
// regardless of how much content the modal itself wants to show. This is
// the guarantee ClampToScreen (see style.go) exists to provide universally,
// instead of each modal having to reason about whether its content could
// ever grow past a comfortable size; this test is what keeps that guarantee
// from silently regressing for any modal, present or future.
func TestViewNeverExceedsAllocatedSize(t *testing.T) {
	const width, height = 22, 6

	for _, modalType := range Registered() {
		t.Run(string(modalType), func(t *testing.T) {
			m, ok := New(modalType)
			if !ok {
				t.Fatalf("New(%s) returned ok=false", modalType)
			}

			m.Init()
			m.SetSize(width, height)

			lines := strings.Split(m.View(), "\n")
			if len(lines) > height {
				t.Errorf("%s: View() produced %d lines for height=%d", modalType, len(lines), height)
			}
			for i, line := range lines {
				if w := lipgloss.Width(line); w > width {
					t.Errorf("%s: line %d is %d cells wide for width=%d", modalType, i, w, width)
				}
			}
		})
	}
}

func TestRegisteredCoversAllModalTypes(t *testing.T) {
	want := []ModalType{ModalAlert, ModalConfirm, ModalForm, ModalSelector}

	got := map[ModalType]bool{}
	for _, mt := range Registered() {
		got[mt] = true
	}

	for _, mt := range want {
		if !got[mt] {
			t.Errorf("modal type %s is not registered", mt)
		}
	}
}
