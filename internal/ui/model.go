package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"tui-template/internal/ui/modals"
	"tui-template/internal/ui/panels"
)

// Model es la raíz de la app. No conoce tipos concretos de panel —
// solo la interface panels.Panel — así el patrón compuesto funciona:
// cada panel maneja su propio Update(), y este Model solo intercepta
// las teclas globales (cambiar foco, salir, abrir modal) antes de
// delegar el resto al panel que tiene el foco actualmente.
type Model struct {
	panelList    []panels.Panel
	focusedIndex int
	width        int
	height       int
	arrangement  Arrangement

	// activeModal es nil cuando no hay overlay abierto. Mientras no
	// sea nil, TODOS los mensajes van al modal — el body queda
	// congelado detrás hasta que el usuario responda o cancele.
	activeModal modals.Modal
}

// NewModel arma el esqueleto de prueba con 3 paneles: list (sidebar),
// table (main), statusbar (footer, no-focusable). En la versión
// generada, este slice se construye iterando los Panel del schema
// y llamando panels.New(blockType) por cada uno.
func NewModel() Model {
	list, _ := panels.New(panels.BlockList)
	table, _ := panels.New(panels.BlockTable)
	status, _ := panels.New(panels.BlockStatusbar)

	m := Model{
		panelList: []panels.Panel{list, table, status},
	}

	m.focusedIndex = firstFocusable(m.panelList)
	if m.focusedIndex >= 0 {
		m.panelList[m.focusedIndex].Focus()
	}

	return m
}

func firstFocusable(ps []panels.Panel) int {
	for i, p := range ps {
		if p.Focusable() {
			return i
		}
	}
	return -1
}

func (m Model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0, len(m.panelList))
	for _, p := range m.panelList {
		cmds = append(cmds, p.Init())
	}
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// tea.WindowSizeMsg se procesa siempre, incluso con un modal
	// activo, para que el layout de fondo no quede desactualizado.
	if sizeMsg, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = sizeMsg.Width
		m.height = sizeMsg.Height
		m.applyLayout()
		return m, nil
	}

	// Si el resultado de un Confirm llega y ya no hay modal activo
	// (se cerró en el ciclo anterior), decidimos qué hacer con la
	// respuesta acá — este es el punto donde, en la app real, se
	// llamaría al stub correspondiente (ej. eliminar el equipo).
	if result, ok := msg.(modals.ConfirmResultMsg); ok {
		return m.handleConfirmResult(result)
	}

	// Con un modal activo, TODO lo demás se le delega a él — el
	// resto del árbol (paneles, teclas globales) queda pausado.
	if m.activeModal != nil {
		return m.updateModal(msg)
	}

	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch {
		case key.Matches(keyMsg, Keys.Quit):
			return m, tea.Quit
		case key.Matches(keyMsg, Keys.NextPanel):
			m.cycleFocus(1)
			return m, nil
		case key.Matches(keyMsg, Keys.PrevPanel):
			m.cycleFocus(-1)
			return m, nil
		case key.Matches(keyMsg, Keys.OpenConfirmDemo):
			return m.openConfirmDemo()
		case key.Matches(keyMsg, Keys.OpenAlertDemo):
			return m.openAlertDemo()
		}
	}

	// Todo lo demás se delega al panel con foco (patrón compuesto).
	if m.focusedIndex >= 0 {
		updated, cmd := m.panelList[m.focusedIndex].Update(msg)
		m.panelList[m.focusedIndex] = updated
		return m, cmd
	}

	return m, nil
}

// updateModal envía el mensaje al modal activo y lo cierra si ya
// terminó su ciclo (Done() == true).
func (m Model) updateModal(msg tea.Msg) (tea.Model, tea.Cmd) {
	updated, cmd := m.activeModal.Update(msg)
	m.activeModal = updated

	if m.activeModal.Done() {
		m.activeModal = nil
	}

	return m, cmd
}

// openConfirmDemo simula la acción "action:openModal" de un Trigger
// del schema (tecla "d" → abrir Confirm de eliminación).
func (m Model) openConfirmDemo() (tea.Model, tea.Cmd) {
	confirm := modals.NewConfirm(
		"Confirmar acción",
		"¿Deseas eliminar el equipo seleccionado?",
		"Eliminar",
		"Cancelar",
	)
	m.activeModal = confirm
	return m, confirm.Init()
}

func (m Model) openAlertDemo() (tea.Model, tea.Cmd) {
	alert := modals.NewAlert(
		"Información",
		"Este es un alert de ejemplo — variante info.",
		modals.AlertInfo,
	)
	m.activeModal = alert
	return m, alert.Init()
}

// handleConfirmResult es el punto de integración con internal/stubs
// en la versión final: si Confirmed == true, aquí se llamaría al
// CallStub real (ej. stubs.DeleteEquipo(id)) en vez de solo abrir
// un alert de éxito como en este demo.
func (m Model) handleConfirmResult(result modals.ConfirmResultMsg) (tea.Model, tea.Cmd) {
	if !result.Confirmed {
		return m, nil
	}

	alert := modals.NewAlert(
		"Éxito",
		"El equipo fue eliminado (demo — sin lógica real todavía).",
		modals.AlertSuccess,
	)
	m.activeModal = alert
	return m, alert.Init()
}

func (m *Model) cycleFocus(direction int) {
	focusable := focusableIndices(m.panelList)
	if len(focusable) == 0 {
		return
	}

	currentPos := 0
	for i, idx := range focusable {
		if idx == m.focusedIndex {
			currentPos = i
			break
		}
	}

	nextPos := (currentPos + direction + len(focusable)) % len(focusable)

	m.panelList[m.focusedIndex].Blur()
	m.focusedIndex = focusable[nextPos]
	m.panelList[m.focusedIndex].Focus()
}

func focusableIndices(ps []panels.Panel) []int {
	idx := make([]int, 0, len(ps))
	for i, p := range ps {
		if p.Focusable() {
			idx = append(idx, i)
		}
	}
	return idx
}

// applyLayout delega el cálculo a computeLayout (layout.go), que
// resuelve breakpoints y devuelve tamaños ya listos. Acá solo se
// aplican a cada panel y se guarda el Arrangement para que View()
// sepa si debe unir sidebar+main horizontal o apilado.
func (m *Model) applyLayout() {
	if len(m.panelList) < 3 {
		return
	}

	result := computeLayout(m.width, m.height)
	m.arrangement = result.Arrangement

	m.panelList[0].SetSize(result.SidebarW, result.SidebarH)
	m.panelList[1].SetSize(result.MainW, result.MainH)
	m.panelList[2].SetSize(result.StatusW, statusHeight)
}

func (m Model) View() string {
	if len(m.panelList) < 3 {
		return "cargando..."
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
		body = lipgloss.JoinVertical(
			lipgloss.Left,
			m.panelList[0].View(),
			m.panelList[1].View(),
		)
	} else {
		body = lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.panelList[0].View(),
			m.panelList[1].View(),
		)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		body,
		m.panelList[2].View(),
		m.renderCommandBar(),
	)
}

// renderCommandBar es la barra de comandos inferior "de guía" que
// pediste: siempre visible, lista los bindings globales activos.
// Cuando agreguemos triggers por panel, esto se extiende para mostrar
// también los bindings específicos del panel con foco actual.
func (m Model) renderCommandBar() string {
	parts := make([]string, 0, len(Keys.ShortHelp()))
	for _, b := range Keys.ShortHelp() {
		parts = append(parts, commandKeyStyle.Render(b.Help().Key)+" "+b.Help().Desc)
	}
	return commandBarStyle.Width(m.width).Render(strings.Join(parts, "  ·  "))
}
