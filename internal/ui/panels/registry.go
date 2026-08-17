package panels

// BlockType espeja el tipo homónimo del schema del editor visual.
type BlockType string

const (
	// Datos / listas
	BlockList      BlockType = "list"
	BlockTable     BlockType = "table"
	BlockViewport  BlockType = "viewport"
	BlockPaginator BlockType = "paginator"

	// Entrada
	BlockTextinput  BlockType = "textinput"
	BlockTextarea   BlockType = "textarea"
	BlockFilepicker BlockType = "filepicker"

	// Estado / feedback
	BlockSpinner  BlockType = "spinner"
	BlockProgress BlockType = "progress"
	BlockHelp     BlockType = "help"
	BlockKey      BlockType = "key"

	// Estructura
	BlockStatusbar BlockType = "statusbar"
	BlockTitlebar  BlockType = "titlebar"
	BlockTabs      BlockType = "tabs"
	BlockTree      BlockType = "tree"

	// Genérico
	BlockCustom BlockType = "custom"
)

// Constructor crea una instancia nueva de un Panel concreto.
type Constructor func() Panel

// registry solo tiene 3 bloques implementados por ahora (list, table,
// statusbar) para validar la orquestación del Model raíz antes de
// escribir los 13 restantes. Cada uno se agrega con una sola línea
// aquí — ese es el punto de extender sin tocar model.go.
var registry = map[BlockType]Constructor{
	BlockList:      NewListPanel,
	BlockTable:     NewTablePanel,
	BlockStatusbar: NewStatusbarPanel,

	// TODO: BlockViewport, BlockPaginator, BlockTextinput,
	// BlockTextarea, BlockFilepicker, BlockSpinner, BlockProgress,
	// BlockHelp, BlockKey, BlockTitlebar, BlockTabs, BlockTree,
	// BlockCustom
}

// New instancia un Panel a partir de su BlockType. El segundo valor
// de retorno es false si el bloque aún no tiene implementación.
func New(blockType BlockType) (Panel, bool) {
	ctor, ok := registry[blockType]
	if !ok {
		return nil, false
	}
	return ctor(), true
}

// Registered lista los BlockType con implementación disponible —
// útil para el togglemenu, que no debería ofrecer bloques sin construir.
func Registered() []BlockType {
	types := make([]BlockType, 0, len(registry))
	for t := range registry {
		types = append(types, t)
	}
	return types
}
