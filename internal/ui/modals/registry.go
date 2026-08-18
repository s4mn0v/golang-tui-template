package modals

// ModalType identifies one of the overlay blocks available through New.
type ModalType string

const (
	ModalAlert    ModalType = "alert"
	ModalConfirm  ModalType = "confirm"
	ModalForm     ModalType = "form"
	ModalSelector ModalType = "selector"
)

// Constructor builds a ready-to-use demo instance of a modal.
type Constructor func() Modal

var registry = map[ModalType]Constructor{
	ModalAlert: func() Modal {
		return NewAlert("Info", "This is an info alert.", AlertInfo)
	},
	ModalConfirm: func() Modal {
		return NewConfirm("Confirm action", "Are you sure you want to continue?", "", "")
	},
	ModalForm: func() Modal {
		return NewForm("New item", [2]string{"name", "Name"}, [2]string{"type", "Type"})
	},
	ModalSelector: func() Modal {
		return NewSelector("Choose an item", "compressor-a", "pump-b", "motor-c")
	},
}

// New builds a demo instance of the given modal type.
func New(modalType ModalType) (Modal, bool) {
	ctor, ok := registry[modalType]
	if !ok {
		return nil, false
	}
	return ctor(), true
}

// Registered returns every modal type available through New.
func Registered() []ModalType {
	types := make([]ModalType, 0, len(registry))
	for t := range registry {
		types = append(types, t)
	}
	return types
}
