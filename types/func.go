package types

// FuncType is a definition type for wasm function
type FuncType struct {
	ParameterTypes []ValType
	ResultTypes    []ValType
}
