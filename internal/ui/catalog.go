package ui

import (
	"tui-template/internal/ui/modals"
	"tui-template/internal/ui/panels"
)

// catalogKind distinguishes a block that can be previewed inline (a Panel)
// from one that only exists as an overlay (a Modal).
type catalogKind int

const (
	catalogPanel catalogKind = iota
	catalogModal
)

// catalogEntry is one row in the sidebar's "Blocks" list: a human label plus
// enough information to build the real block on demand via panels.New or
// modals.New.
type catalogEntry struct {
	label       string
	description string
	kind        catalogKind
	panelType   panels.BlockType
	modalType   modals.ModalType
}

func panelEntry(label, description string, blockType panels.BlockType) catalogEntry {
	return catalogEntry{label: label, description: description, kind: catalogPanel, panelType: blockType}
}

func modalEntry(label, description string, modalType modals.ModalType) catalogEntry {
	return catalogEntry{label: label, description: description, kind: catalogModal, modalType: modalType}
}

// buildCatalog lists every block and modal shipped with the template, grouped
// the same way as the project's block catalog. Selecting a panel entry
// previews it live in the main panel; selecting a modal entry opens it as an
// overlay.
func buildCatalog() []catalogEntry {
	return []catalogEntry{
		// Data / Lists
		panelEntry("list", "navigable list with fuzzy filter", panels.BlockList),
		panelEntry("table", "table with columns/rows", panels.BlockTable),
		panelEntry("viewport", "long text with scrolling (logs, results, docs)", panels.BlockViewport),
		panelEntry("paginator", "pagination indicator/control", panels.BlockPaginator),

		// Input
		panelEntry("textinput", "single-line field", panels.BlockTextinput),
		panelEntry("textarea", "multi-line field", panels.BlockTextarea),
		panelEntry("filepicker", "file/directory picker", panels.BlockFilepicker),

		// Status / Feedback
		panelEntry("spinner", "loading indicator", panels.BlockSpinner),
		panelEntry("progress", "progress bar", panels.BlockProgress),
		panelEntry("help", "help panel with shortcuts", panels.BlockHelp),
		panelEntry("key", "key/shortcut indicator", panels.BlockKey),

		// Structure
		panelEntry("statusbar", "status bar (footer)", panels.BlockStatusbar),
		panelEntry("titlebar", "title bar (header)", panels.BlockTitlebar),
		panelEntry("tabs", "tabbed navigation", panels.BlockTabs),
		panelEntry("tree", "hierarchical tree (schemas, filesystem, git)", panels.BlockTree),

		// Modals (overlay, doesn't occupy grid space)
		modalEntry("modal:alert", "info / error / success, one button", modals.ModalAlert),
		modalEntry("modal:confirm", "yes/no", modals.ModalConfirm),
		modalEntry("modal:form", "form (huh)", modals.ModalForm),
		modalEntry("modal:selector", "picker with fuzzy search", modals.ModalSelector),

		// Generic
		panelEntry("custom", "empty block, no assigned component", panels.BlockCustom),
	}
}

func catalogLabels(catalog []catalogEntry) []string {
	labels := make([]string, len(catalog))
	for i, entry := range catalog {
		labels[i] = entry.label
	}
	return labels
}
