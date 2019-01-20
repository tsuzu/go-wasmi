package types

// ValType is a value type for wasm
type ValType uint32

const (
	// I32 is 32-bit integer type
	I32 ValType = 0x7f
	// I64 is 64-bit integer type
	I64 ValType = 0x7e
	// F32 is 32-bit float type
	F32 ValType = 0x7d
	// F64 is 64-bit float type
	F64 ValType = 0x7c
)

// ValTypes contains all value types
var ValTypes = []ValType{I32, I64, F32, F64}
