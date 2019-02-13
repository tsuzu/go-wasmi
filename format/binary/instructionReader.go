package binary

import (
	"io"
	"math"

	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"
	"github.com/pkg/errors"
)

var (
	// ErrElseWithoutIf is returned when 'else' opcode is used after not 'if' opcode
	ErrElseWithoutIf = errors.New("'else' must follow 'if'")

	// ErrNotConstExpr is returned when not-constant opcode is used in constant expressions
	ErrNotConstExpr = errors.New("not constant expression")
)

func readInstructionSimple(r io.Reader, t types.InstructionType) (*types.InstructionSimple, error) {
	return &types.InstructionSimple{
		Instruction: t,
	}, nil
}

func readInstructionSimpleWith0x00(r io.Reader, t types.InstructionType) (*types.InstructionSimple, error) {
	instr, err := readInstructionSimple(r, t)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	b, err := binrw.ReadByte(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if b != 0x00 {
		return nil, ErrInvalidFormat
	}

	return instr, nil
}

func readInstructionBranchTable(r io.Reader, t types.InstructionType) (*types.InstructionBranchTable, error) {
	var indices []types.LabelIndex
	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if indices == nil {
			indices = make([]types.LabelIndex, 0, size)
		}

		index, err := leb128.ReadUint32(r)

		if err != nil {
			return errors.WithStack(err)
		}

		indices = append(indices, types.LabelIndex(index))

		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	d, err := leb128.ReadUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &types.InstructionBranchTable{
		Instruction: t,
		Indices:     indices,
		Default:     types.LabelIndex(d),
	}, nil
}

func readInstructionLocalIndex(r io.Reader, t types.InstructionType) (*types.InstructionLocalIndex, error) {
	res := &types.InstructionLocalIndex{
		Instruction: t,
	}

	index, err := leb128.ReadUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	res.Index = types.LocalIndex(index)

	return res, nil
}

func readInstructionGlobalIndex(r io.Reader, t types.InstructionType) (*types.InstructionGlobalIndex, error) {
	res := &types.InstructionGlobalIndex{
		Instruction: t,
	}

	index, err := leb128.ReadUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	res.Index = types.GlobalIndex(index)

	return res, nil
}

func readInstructionFuncIndex(r io.Reader, t types.InstructionType) (*types.InstructionFuncIndex, error) {
	res := &types.InstructionFuncIndex{
		Instruction: t,
	}

	index, err := leb128.ReadUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	res.Index = types.FuncIndex(index)

	return res, nil
}

func readInstructionLabelIndex(r io.Reader, t types.InstructionType) (*types.InstructionLabelIndex, error) {
	res := &types.InstructionLabelIndex{
		Instruction: t,
	}

	index, err := leb128.ReadUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	res.Index = types.LabelIndex(index)

	return res, nil
}

func readInstructionCallIndirect(r io.Reader, t types.InstructionType) (*types.InstructionTypeIndex, error) {
	res := &types.InstructionTypeIndex{
		Instruction: t,
	}

	index, err := leb128.ReadUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	res.Index = types.TypeIndex(index)

	b, err := binrw.ReadByte(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if b != 0x00 {
		return nil, ErrInvalidFormat
	}

	return res, nil
}

func readInstructionBlock(r io.Reader, t types.InstructionType) (*types.InstructionBlock, error) {
	res := &types.InstructionBlock{
		Instruction: t,
	}

	b, err := binrw.ReadByte(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	instr, err := ReadExpression(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	res.BlockType = types.ValType(b)
	res.Instructions = instr

	return res, nil
}

func readInstructionBlockIfElse(r io.Reader, block *types.InstructionBlock) (*types.InstructionBlockIfElse, error) {
	res := &types.InstructionBlockIfElse{
		InstructionBlock: block,
	}

	instr, err := ReadExpression(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	res.ElseInstructions = instr

	return res, nil
}

func readInstructionMemArg(r io.Reader, t types.InstructionType) (*types.InstructionMemArg, error) {
	res := &types.InstructionMemArg{
		Instruction: t,
	}
	arg := types.MemoryArgument{}

	var err error

	arg.Allignment, err = leb128.ReadUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	arg.Offset, err = leb128.ReadUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	res.MemArg = arg

	return res, nil
}

func readInstructionConst(r io.Reader, t types.InstructionType) (*types.InstructionConst, error) {
	res := &types.InstructionConst{
		Instruction: t,
	}

	switch t {
	case types.I32Const:
		v, err := leb128.ReadInt32(r)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		res.SetInt32(v)
	case types.I64Const:
		v, err := leb128.ReadInt64(r)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		res.SetInt64(v)
	case types.F32Const:
		b, err := binrw.ReadLEUint32(r)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		res.SetFloat32(math.Float32frombits(b))
	case types.F64Const:
		b, err := binrw.ReadLEUint64(r)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		res.SetFloat64(math.Float64frombits(b))
	default:
		panic("unknown instruction")
	}

	return res, nil
}

func readInstruction(r io.Reader) (types.InstructionInterface, error) {
	b, err := binrw.ReadByte(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	t := types.InstructionType(b)
	switch t {
	case types.End:
		return nil, io.EOF
	case types.Else:
		return nil, ErrElseWithoutIf
	case types.Branch, types.BranchIf:
		return readInstructionLabelIndex(r, t)
	case types.BranchTable:
		return readInstructionBranchTable(r, t)
	case types.Block, types.Loop:
		return readInstructionBlock(r, t)
	case types.If:
		instr, err := readInstructionBlock(r, t)

		if errors.Cause(err) != ErrElseWithoutIf {
			return instr, errors.WithStack(err)
		}

		return readInstructionBlockIfElse(r, instr)
	case types.Call:
		return readInstructionFuncIndex(r, t)
	case types.CallIndirect:
		return readInstructionCallIndirect(r, t)
	case types.LocalGet, types.LocalSet, types.LocalTee:
		return readInstructionLocalIndex(r, t)
	case types.GlobalGet, types.GlobalSet:
		return readInstructionGlobalIndex(r, t)
	case types.MemorySize, types.MemoryGrow:
		return readInstructionSimpleWith0x00(r, t)
	case types.I32Const, types.I64Const, types.F32Const, types.F64Const:
		return readInstructionConst(r, t)
	}

	if t >= types.I32Load && t <= types.I64Store32 {
		return readInstructionMemArg(r, t)
	}

	return readInstructionSimple(r, t)
}

// ReadExpression parses instructions and end.
func ReadExpression(r io.Reader) ([]types.InstructionInterface, error) {
	res := make([]types.InstructionInterface, 0, 10)

	for {
		r, err := readInstruction(r)

		if err != nil {
			if errors.Cause(err) == io.EOF {
				err = nil
			}

			return res, errors.WithStack(err)
		}

		res = append(res, r)
	}
}
