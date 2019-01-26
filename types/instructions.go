package types

// InstructionType represents instruction
type InstructionType byte

// InstructionSimple represents an InstructionSimple
// taking no parameters
type InstructionSimple struct {
	Instruction InstructionType
}

// InstructionLocalIndex represents an instruction
// taking a localindex parameter
type InstructionLocalIndex struct {
	Instruction InstructionType
	Index       LocalIndex
}

// InstructionGlobalIndex represents an instruction
// taking a globalindex parameter
type InstructionGlobalIndex struct {
	Instruction InstructionType
	Index       GlobalIndex
}

// MemoryArgument represents memargs immediate
type MemoryArgument struct {
	Allignment, Offset uint32
}

// InstructionMemArg represents an instruction
// taking a memarg paramter
type InstructionMemArg struct {
	Instruction InstructionType
	MemArg      MemoryArgument
}

const (
	Unreachable InstructionType = 0x00 + iota
	Nop
	Block
	Loop
	If
	Else
)

const (
	End InstructionType = 0x0b + iota
	Branch
	BranchIf
	BranchTable
	Return
)

const (
	Call InstructionType = 0x10 + iota
	CallIndirect
)

const (
	// Drop => drop
	Drop InstructionType = 0x1a + iota
	// Select => select
	Select
)

const (
	LocalGet InstructionType = 0x20 + iota
	LocalSet
	LocalTee
	GlobalGet
	GlobalSet
)

const (
	I32Load InstructionType = 0x28 + iota
	I64Load
	F32Load
	F64Load
	I32Load8S
	I32Load8U
	I32Load16S
	I32Load16U
	I64Load8S
	I64Load8U
	I64Load16S
	I64Load16U
	I64Load32S
	I64Load32U
	I32Store
	I64Store
	F32Store
	F64Store
	I32Store8
	I32Store16
	I64Store8
	I64Store16
	I64Store32
	MemorySize // followed by 0x00
	MemoryGrow // followed by 0x00
)

const (
	I32Const InstructionType = 0x41 + iota
	I64Const
	F32Const
	F64Const
)

const (
	I32Eqz InstructionType = 0x45 + iota
	I32Eq
	I32Ne
	I32LtS
	I32LtU
	I32GtS
	I32GtU
	I32LeS
	I32LeU
	I32GeS
	I32GeU

	I64Eqz // 0x50
	I64Eq
	I64Ne
	I64LtS
	I64LtU
	I64GtS
	I64GtU
	I64LeS
	I64LeU
	I64GeS
	I64GeU

	F32Eq // 0x5b
	F32Ne
	F32Lt
	F32Gt
	F32Le
	F32Ge

	F64Eq // 0x61
	F64Ne
	F64Lt
	F64Gt
	F64Le
	F64Ge

	I32Clz // 0x67
	I32Ctz
	I32Popcnt
	I32Add
	I32Sub
	I32Mul
	I32DivS
	I32DivU
	I32RemS
	I32RemU
	I32And
	I32Or
	I32Xor
	I32Shl
	I32ShrS
	I32ShrU
	I32Rotl
	I32Rotr

	I64Clz // 0x79
	I64Ctz
	I64Popcnt
	I64Add
	I64Sub
	I64Mul
	I64DivS
	I64DivU
	I64RemS
	I64RemU
	I64And
	I64Or
	I64Xor
	I64Shl
	I64ShrS
	I64ShrU
	I64Rotl
	I64Rotr

	F32Abs // 0x8b
	F32Neg
	F32Ceil
	F32Floor
	F32Trunc
	F32Nearest
	F32Sqrt
	F32Add
	F32Sub
	F32Mul
	F32Div
	F32Min
	F32Max
	F32Copysign

	F64Abs // 0x99
	F64Neg
	F64Ceil
	F64Floor
	F64Trunc
	F64Nearest
	F64Sqrt
	F64Add
	F64Sub
	F64Mul
	F64Div
	F64Min
	l64Max
	F64Copysign

	I32WrapI64 // 0xa7
	I32TruncF32S
	I32TruncF32U
	I32TruncF64S
	I32TruncF64U

	I64ExtendI32S
	I64ExtendI32U
	I64TruncF32S
	I64TruncF32U
	I64TruncF64S
	I64TruncF64U

	F32ConvertI32S
	F32ConvertI32U
	F32ConvertI64S
	F32ConvertI64U
	F32DemoteF64

	F64ConvertI32S
	F64ConvertI32U
	F64ConvertI64S
	F64ConvertI64U
	F64PromoteF32

	I32ReinterpretF32
	I64ReinterpretF64
	F32ReinterpretI32
	F64ReinterpretI64 // 0xbf
)
