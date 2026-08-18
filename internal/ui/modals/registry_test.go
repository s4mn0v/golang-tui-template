package modals

import "testing"

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
