package types

import (
	"testing"
)

func checkInstructionOpcode(t *testing.T, actual, expected InstructionType) {
	if expected != actual {
		t.Errorf("The expected value is 0x%02x, but got 0x%02x", expected, actual)
	}
}

func TestInstructionOpcode(t *testing.T) {
	checkInstructionOpcode(t, Unreachable, 0x00)
	checkInstructionOpcode(t, Else, 0x05)
	checkInstructionOpcode(t, End, 0x0b)
	checkInstructionOpcode(t, CallIndirect, 0x11)

	checkInstructionOpcode(t, Drop, 0x1a)
	checkInstructionOpcode(t, Select, 0x1b)

	checkInstructionOpcode(t, LocalGet, 0x20)
	checkInstructionOpcode(t, GlobalSet, 0x24)

	checkInstructionOpcode(t, I32Load, 0x28)
	checkInstructionOpcode(t, MemoryGrow, 0x40)

	checkInstructionOpcode(t, I32Const, 0x41)
	checkInstructionOpcode(t, F64Const, 0x44)

	checkInstructionOpcode(t, I32Eqz, 0x45)
	checkInstructionOpcode(t, I64Eqz, 0x50)
	checkInstructionOpcode(t, F64Eq, 0x61)
	checkInstructionOpcode(t, I32Clz, 0x67)
	checkInstructionOpcode(t, I64Clz, 0x79)
	checkInstructionOpcode(t, F32Abs, 0x8b)
	checkInstructionOpcode(t, F64Abs, 0x99)
	checkInstructionOpcode(t, I32WrapI64, 0xa7)
	checkInstructionOpcode(t, F64ReinterpretI64, 0xbf)
}
