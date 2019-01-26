package types

// ExportDescriptionKind represents a type of export section
type ExportDescriptionKind byte

const (
	// ExportDescriptionFunc represents type index for export section
	ExportDescriptionFunc ExportDescriptionKind = iota
	// ExportDescriptionTable represents table type for export section
	ExportDescriptionTable
	// ExportDescriptionMemory represents memory type for export section
	ExportDescriptionMemory
	// ExportDescriptionGlobal represents global type for export section
	ExportDescriptionGlobal
)

// ExportType represents export section
type ExportType struct {
	Name            string
	DescriptionKind ExportDescriptionKind
	Index           uint32 // funcidx/tableidx/memidx/globalidx
}
