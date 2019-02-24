package wasmi

import "github.com/cs3238-tsuzu/go-wasmi/types"

type SectionCustom struct {
	Name string
	Data []byte
}
type SectionType struct {
	FuncTypes []*types.FuncType
}
type SectionImport struct {
	ImportTypes []*types.ImportType
}
type SectionFunction struct {
	Types []types.TypeIndex
}
type SectionTable struct {
	Types []*types.TableType
}
type SectionMemory struct {
	Types []*types.MemoryType
}

type SectionGlobalElementType struct {
	Type types.GlobalType
	Expr []types.InstructionInterface
}

type SectionGlobal struct {
	Globals []SectionGlobalElementType
}
type SectionExport struct {
	Exports []*types.ExportType
}
type SectionStart struct {
	FuncIndex types.FuncIndex
}

type SectionElementElementType struct {
	TableIndex  types.TableIndex
	Expr        []types.InstructionInterface
	FuncIndices []types.FuncIndex
}
type SectionElement struct {
	Elements []SectionElementElementType
}

type SectionCodeLocalElement struct {
	Size    uint32
	ValType types.ValType
}
type SectionCodeElementType struct {
	Locals []SectionCodeLocalElement
	Expr   []types.InstructionInterface
}

type SectionCode struct {
	Codes []SectionCodeElementType
}

type SectionDataElement struct {
	MemoryIndex types.MemoryIndex
	Expr        []types.InstructionInterface
	Bytes       []byte
}
type SectionData struct {
	Data []SectionDataElement
}
