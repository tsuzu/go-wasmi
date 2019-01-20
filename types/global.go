package types

// GlobalMutability represents whether global is constant or variable
type GlobalMutability byte

const (
	// GlobalConst represents global constant
	GlobalConst GlobalMutability = iota
	// GlobalVariable represents global variable
	GlobalVariable
)

// GlobalType represents global type
type GlobalType struct {
	Type       ValType
	Mutability GlobalMutability
}
