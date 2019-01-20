package types

// TableElemType is only for FuncRef in wasm 1.0
type TableElemType byte

// TableElemFuncRef represents all types of functions
const TableElemFuncRef TableElemType = 0x70

// TableType represents function table
type TableType struct {
	ElemType TableElemType // Only funcref
	Limits
}
