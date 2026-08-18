package panels

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// treeNode is a hierarchical entry (e.g. filesystem path, git ref, schema).
type treeNode struct {
	label    string
	expanded bool
	children []*treeNode
}

var (
	treeCursorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("230")).
			Background(lipgloss.Color("62"))

	treeDimStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
)

// TreePanel is a hierarchical tree (schemas, filesystem, git) with
// up/down navigation and expand/collapse.
type TreePanel struct {
	root          *treeNode
	cursor        int
	focused       bool
	width, height int
}

func NewTreePanel() Panel {
	root := &treeNode{
		label:    "project",
		expanded: true,
		children: []*treeNode{
			{label: "internal", expanded: true, children: []*treeNode{
				{label: "ui", expanded: false, children: []*treeNode{
					{label: "panels"},
					{label: "modals"},
				}},
			}},
			{label: "cmd", children: []*treeNode{
				{label: "app"},
			}},
			{label: "go.mod"},
		},
	}

	return &TreePanel{root: root}
}

// flatNode is a visible row: the node plus its depth for indentation.
type flatNode struct {
	node  *treeNode
	depth int
}

func (p *TreePanel) flatten() []flatNode {
	var rows []flatNode
	var walk func(n *treeNode, depth int)
	walk = func(n *treeNode, depth int) {
		rows = append(rows, flatNode{node: n, depth: depth})
		if n.expanded {
			for _, c := range n.children {
				walk(c, depth+1)
			}
		}
	}
	walk(p.root, 0)
	return rows
}

func (p *TreePanel) Init() tea.Cmd { return nil }

func (p *TreePanel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return p, nil
	}

	rows := p.flatten()
	if len(rows) == 0 {
		return p, nil
	}

	switch keyMsg.String() {
	case "up", "k":
		if p.cursor > 0 {
			p.cursor--
		}
	case "down", "j":
		if p.cursor < len(rows)-1 {
			p.cursor++
		}
	case "enter", " ", "l", "h":
		node := rows[p.cursor].node
		if len(node.children) > 0 {
			node.expanded = !node.expanded
		}
	}

	return p, nil
}

func (p *TreePanel) View() string {
	rows := p.flatten()

	var b strings.Builder
	for i, row := range rows {
		prefix := strings.Repeat("  ", row.depth)

		marker := "  "
		if len(row.node.children) > 0 {
			if row.node.expanded {
				marker = "▾ "
			} else {
				marker = "▸ "
			}
		}

		line := prefix + marker + row.node.label
		if i == p.cursor {
			line = treeCursorStyle.Render(line)
		} else if len(row.node.children) == 0 {
			line = treeDimStyle.Render(line)
		}

		b.WriteString(line)
		if i < len(rows)-1 {
			b.WriteString("\n")
		}
	}

	return borderStyleFor(p.focused).
		Width(OuterStyleWidth(p.width)).
		Height(OuterStyleHeight(p.height)).
		Render(b.String())
}

func (p *TreePanel) SetSize(w, h int) { p.width, p.height = w, h }

func (p *TreePanel) Focus()          { p.focused = true }
func (p *TreePanel) Blur()           { p.focused = false }
func (p *TreePanel) Focused() bool   { return p.focused }
func (p *TreePanel) Focusable() bool { return true }
func (p *TreePanel) Title() string   { return "Tree" }
