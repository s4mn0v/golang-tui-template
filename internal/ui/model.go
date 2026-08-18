package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"tui-template/internal/ui/modals"
	"tui-template/internal/ui/panels"
)

// Model is a showcase of every block in the template: the sidebar lists the
// full catalog (built in catalog.go), the main panel previews whichever
// block is selected, and picking a modal entry opens it as an overlay. It
// doubles as a demo app and as a reference for how the pieces fit together.
type Model struct {
	titlebar  panels.Panel
	sidebar   *panels.ListPanel
	preview   panels.Panel
	statusbar *panels.StatusbarPanel

	catalog        []catalogEntry
	selectedIdx    int
	focusOnPreview bool

	width, height        int
	arrangement          Arrangement
	lastMainW, lastMainH int

	activeModal modals.Modal
}

func NewModel() Model {
	catalog := buildCatalog()

	titlebar, _ := panels.New(panels.BlockTitlebar)
	titlebar.(*panels.TitlebarPanel).SetTitle("tui-template — block showcase")

	sidebarPanel, _ := panels.New(panels.BlockList)
	sidebar := sidebarPanel.(*panels.ListPanel)
	sidebar.SetTitle("Blocks")
	sidebar.SetItems(catalogLabels(catalog))
	sidebar.Focus()

	statusbarPanel, _ := panels.New(panels.BlockStatusbar)
	statusbar := statusbarPanel.(*panels.StatusbarPanel)

	m := Model{
		titlebar:  titlebar,
		sidebar:   sidebar,
		statusbar: statusbar,
		catalog:   catalog,
	}

	m.setPreview(0)

	return m
}

// setPreview updates the status line to describe the highlighted catalog
// entry and, if it's a panel block, rebuilds the main preview from it via
// panels.New.
func (m *Model) setPreview(idx int) {
	entry := m.catalog[idx]
	m.statusbar.SetText(entry.label + " — " + entry.description)

	if entry.kind != catalogPanel {
		return
	}

	p, _ := panels.New(entry.panelType)
	p.SetSize(m.lastMainW, m.lastMainH)
	m.preview = p
}

func (m Model) Init() tea.Cmd {
	cmds := []tea.Cmd{m.titlebar.Init(), m.sidebar.Init(), m.statusbar.Init()}
	if m.preview != nil {
		cmds = append(cmds, m.preview.Init())
	}
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if sizeMsg, ok := msg.(tea.WindowSizeMsg); ok {
		m.width, m.height = sizeMsg.Width, sizeMsg.Height
		m.applyLayout()
		return m, nil
	}

	switch result := msg.(type) {
	case modals.ConfirmResultMsg:
		return m.handleConfirmResult(result)
	case modals.FormResultMsg:
		return m.handleFormResult(result)
	case modals.SelectorResultMsg:
		return m.handleSelectorResult(result)
	}

	if m.activeModal != nil {
		return m.updateModal(msg)
	}

	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch {
		case key.Matches(keyMsg, Keys.Quit):
			return m, tea.Quit
		case key.Matches(keyMsg, Keys.NextPanel), key.Matches(keyMsg, Keys.PrevPanel):
			m.toggleFocus()
			return m, nil
		}

		if !m.focusOnPreview {
			return m.updateSidebar(keyMsg)
		}
	}

	if m.focusOnPreview && m.preview != nil {
		updated, cmd := m.preview.Update(msg)
		m.preview = updated
		return m, cmd
	}

	return m, nil
}

// updateSidebar routes a key either to the "activate" action (open the
// highlighted block/modal) or to the underlying list, then checks whether
// the cursor moved so the preview can follow it reactively.
func (m Model) updateSidebar(keyMsg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if key.Matches(keyMsg, Keys.Activate) && !m.sidebar.Filtering() {
		return m.activateSelection()
	}

	updated, cmd := m.sidebar.Update(keyMsg)
	m.sidebar = updated.(*panels.ListPanel)

	if idx := m.sidebar.Index(); idx != m.selectedIdx {
		m.selectedIdx = idx
		m.setPreview(idx)
	}

	return m, cmd
}

// activateSelection opens the highlighted catalog entry: a modal is shown as
// an overlay via modals.New, a panel hands keyboard focus over to the
// preview so it can be interacted with directly.
func (m Model) activateSelection() (tea.Model, tea.Cmd) {
	entry := m.catalog[m.selectedIdx]

	if entry.kind == catalogModal {
		modal, _ := modals.New(entry.modalType)
		m.activeModal = modal
		return m, modal.Init()
	}

	if m.preview != nil && m.preview.Focusable() {
		m.sidebar.Blur()
		m.preview.Focus()
		m.focusOnPreview = true
	}

	return m, nil
}

func (m Model) updateModal(msg tea.Msg) (tea.Model, tea.Cmd) {
	updated, cmd := m.activeModal.Update(msg)
	m.activeModal = updated

	if m.activeModal.Done() {
		m.activeModal = nil
	}

	return m, cmd
}

func (m Model) handleConfirmResult(result modals.ConfirmResultMsg) (tea.Model, tea.Cmd) {
	m.activeModal = nil
	if result.Confirmed {
		m.statusbar.SetText("modal:confirm — confirmed")
	} else {
		m.statusbar.SetText("modal:confirm — cancelled")
	}
	return m, nil
}

func (m Model) handleFormResult(result modals.FormResultMsg) (tea.Model, tea.Cmd) {
	m.activeModal = nil
	if result.Aborted {
		m.statusbar.SetText("modal:form — cancelled")
		return m, nil
	}
	m.statusbar.SetText(fmt.Sprintf("modal:form — submitted %v", result.Values))
	return m, nil
}

func (m Model) handleSelectorResult(result modals.SelectorResultMsg) (tea.Model, tea.Cmd) {
	m.activeModal = nil
	if result.Cancelled {
		m.statusbar.SetText("modal:selector — cancelled")
		return m, nil
	}
	m.statusbar.SetText("modal:selector — picked " + result.Value)
	return m, nil
}

// toggleFocus swaps keyboard focus between the sidebar catalog and the main
// preview. There are only ever two focus targets, and the preview is one of
// them only when the currently previewed block is itself focusable.
func (m *Model) toggleFocus() {
	if m.preview == nil || !m.preview.Focusable() {
		return
	}

	if m.focusOnPreview {
		m.preview.Blur()
		m.sidebar.Focus()
	} else {
		m.sidebar.Blur()
		m.preview.Focus()
	}
	m.focusOnPreview = !m.focusOnPreview
}

func (m *Model) applyLayout() {
	result := computeLayout(m.width, m.height)
	m.arrangement = result.Arrangement

	m.titlebar.SetSize(result.TitlebarW, titlebarHeight)
	m.sidebar.SetSize(result.SidebarW, result.SidebarH)

	m.lastMainW, m.lastMainH = result.MainW, result.MainH
	if m.preview != nil {
		m.preview.SetSize(result.MainW, result.MainH)
	}

	m.statusbar.SetSize(result.StatusW, statusHeight)
}

func (m Model) View() string {
	if m.width == 0 {
		return "loading..."
	}

	if m.activeModal != nil {
		return lipgloss.Place(
			m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			m.activeModal.View(),
		)
	}

	var body string
	if m.arrangement == ArrangeStacked {
		body = lipgloss.JoinVertical(lipgloss.Left, m.sidebar.View(), m.preview.View())
	} else {
		body = lipgloss.JoinHorizontal(lipgloss.Top, m.sidebar.View(), m.preview.View())
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.titlebar.View(),
		body,
		m.statusbar.View(),
		m.renderCommandBar(),
	)
}

func (m Model) renderCommandBar() string {
	parts := make([]string, 0, len(Keys.ShortHelp()))
	for _, b := range Keys.ShortHelp() {
		parts = append(parts, commandKeyStyle.Render(b.Help().Key)+" "+b.Help().Desc)
	}
	return commandBarStyle.Width(m.width).Render(strings.Join(parts, "  ·  "))
}
