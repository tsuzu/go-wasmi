package types

// ImportDescriptionKind represents a type of import section
type ImportDescriptionKind byte

const (
	// ImportDescriptionType represents type index for import section
	ImportDescriptionType ImportDescriptionKind = iota
	// ImportDescriptionTable represents table type for import section
	ImportDescriptionTable
	// ImportDescriptionMemory represents memory type for import section
	ImportDescriptionMemory
	// ImportDescriptionGlobal represents global type for import section
	ImportDescriptionGlobal
)

// ImportType represents import section
type ImportType struct {
	Module          string
	Name            string
	DescriptionKind ImportDescriptionKind

	// Only one of the following is used
	Type   TypeIndex
	Table  *TableType
	Memory *MemoryType
	Global *GlobalType
}
