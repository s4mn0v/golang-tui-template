package panels

type BlockType string

const (
	BlockList      BlockType = "list"
	BlockTable     BlockType = "table"
	BlockViewport  BlockType = "viewport"
	BlockPaginator BlockType = "paginator"

	BlockTextinput  BlockType = "textinput"
	BlockTextarea   BlockType = "textarea"
	BlockFilepicker BlockType = "filepicker"

	BlockSpinner  BlockType = "spinner"
	BlockProgress BlockType = "progress"
	BlockHelp     BlockType = "help"
	BlockKey      BlockType = "key"

	BlockStatusbar BlockType = "statusbar"
	BlockTitlebar  BlockType = "titlebar"
	BlockTabs      BlockType = "tabs"
	BlockTree      BlockType = "tree"

	BlockCustom BlockType = "custom"
)

type Constructor func() Panel

var registry = map[BlockType]Constructor{
	BlockList:      NewListPanel,
	BlockTable:     NewTablePanel,
	BlockViewport:  NewViewportPanel,
	BlockPaginator: NewPaginatorPanel,

	BlockTextinput:  NewTextinputPanel,
	BlockTextarea:   NewTextareaPanel,
	BlockFilepicker: NewFilepickerPanel,

	BlockSpinner:  NewSpinnerPanel,
	BlockProgress: NewProgressPanel,
	BlockHelp:     NewHelpPanel,
	BlockKey:      NewKeyPanel,

	BlockStatusbar: NewStatusbarPanel,
	BlockTitlebar:  NewTitlebarPanel,
	BlockTabs:      NewTabsPanel,
	BlockTree:      NewTreePanel,

	BlockCustom: NewCustomPanel,
}

func New(blockType BlockType) (Panel, bool) {
	ctor, ok := registry[blockType]
	if !ok {
		return nil, false
	}
	return ctor(), true
}

func Registered() []BlockType {
	types := make([]BlockType, 0, len(registry))
	for t := range registry {
		types = append(types, t)
	}
	return types
}
