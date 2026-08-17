package ui

// FocusChangedMsg se emite cuando el panel con foco cambia — útil
// para que un panel reaccione a perder/ganar foco sin acoplarse
// directamente al Model raíz.
type FocusChangedMsg struct {
	PanelIndex int
}

// Los mensajes de modales (ModalOpenedMsg, ModalClosedMsg,
// ConfirmResultMsg, etc.) se agregan aquí cuando construyamos
// internal/ui/modals/ — el esqueleto actual todavía no los necesita.
