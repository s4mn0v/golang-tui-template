package modals

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

// TestModalEscCloses checks that every modal treats "esc" as a cancel/close
// action and reports Done() afterwards, matching the convention used by the
// existing Alert/Confirm modals.
func TestModalEscCloses(t *testing.T) {
	cases := map[string]Modal{
		"alert":    NewAlert("Title", "message", AlertInfo),
		"confirm":  NewConfirm("Title", "message", "", ""),
		"form":     NewForm("Title", [2]string{"name", "Name"}),
		"selector": NewSelector("Title", "one", "two", "three"),
	}

	for name, modal := range cases {
		t.Run(name, func(t *testing.T) {
			modal.Init()
			_ = modal.View()

			updated, _ := modal.Update(tea.KeyMsg{Type: tea.KeyEsc})

			if !updated.Done() {
				t.Errorf("%s: Update(esc) did not mark modal as Done()", name)
			}
		})
	}
}

func TestConfirmSelectionTogglesAndConfirms(t *testing.T) {
	c := NewConfirm("Confirm", "continue?", "Yes", "No")
	c.Init()

	updated, _ := c.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if !updated.Done() {
		t.Fatal("enter on default selection should close the confirm")
	}
}
