package types

type Section interface{}

type SectionCustom struct {
	Name string
	Data []byte
}
type SectionType struct {
	FuncTypes []*FuncType
}
type SectionImport struct {
	ImportTypes []*ImportType
}
type SectionFunction struct {
	Types []TypeIndex
}
type SectionTable struct {
	Types []*TableType
}
type SectionMemory struct {
	Types []*MemoryType
}

type SectionGlobalElementType struct {
	Type GlobalType
	Expr []InstructionInterface
}

type SectionGlobal struct {
	Globals []SectionGlobalElementType
}
type SectionExport struct {
	Exports []*ExportType
}
type SectionStart struct {
	FuncIndex FuncIndex
}

type SectionElementElementType struct {
	TableIndex  TableIndex
	Expr        []InstructionInterface
	FuncIndices []FuncIndex
}
type SectionElement struct {
	Elements []SectionElementElementType
}

type SectionCodeLocalElement struct {
	Size    uint32
	ValType ValType
}
type SectionCodeElementType struct {
	Locals []SectionCodeLocalElement
	Expr   []InstructionInterface
}

type SectionCode struct {
	Codes []SectionCodeElementType
}

type SectionDataElement struct {
	MemoryIndex MemoryIndex
	Expr        []InstructionInterface
	Bytes       []byte
}
type SectionData struct {
	Data []SectionDataElement
}
