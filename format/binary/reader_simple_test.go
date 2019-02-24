package binary

import (
	"bytes"
	"testing"

	"github.com/cs3238-tsuzu/go-wasmi/testutil"
	"github.com/cs3238-tsuzu/go-wasmi/types"
)

func TestFunctionEmpty(t *testing.T) {
	const wat = `(module
		(func)
	)
	`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse empty function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: make([]types.ValType, 0),
						ResultTypes:    make([]types.ValType, 0),
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionCode:
			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{},
						Expr:   []types.InstructionInterface{},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionSomeParams(t *testing.T) {
	const wat = `(module
		(func (param i32 f64 i64))
	  )
	`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: []types.ValType{
							types.I32,
							types.F64,
							types.I64,
						},
						ResultTypes: make([]types.ValType, 0),
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionCode:
			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{},
						Expr:   []types.InstructionInterface{},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionSomeLocals(t *testing.T) {
	const wat = `(module
		(func (local i32 i32))
	  )	  
	`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: make([]types.ValType, 0),
						ResultTypes:    make([]types.ValType, 0),
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionCode:
			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{
							types.SectionCodeLocalElement{
								Size:    2,
								ValType: types.I32,
							},
						},
						Expr: []types.InstructionInterface{},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionSomeLocals2(t *testing.T) {
	const wat = `(module
		(func (local i32 f32 i64 i32 f64))
	  )	  
	`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: make([]types.ValType, 0),
						ResultTypes:    make([]types.ValType, 0),
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionCode:
			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{
							types.SectionCodeLocalElement{
								Size:    1,
								ValType: types.I32,
							},
							types.SectionCodeLocalElement{
								Size:    1,
								ValType: types.F32,
							},
							types.SectionCodeLocalElement{
								Size:    1,
								ValType: types.I64,
							},
							types.SectionCodeLocalElement{
								Size:    1,
								ValType: types.I32,
							},
							types.SectionCodeLocalElement{
								Size:    1,
								ValType: types.F64,
							},
						},
						Expr: []types.InstructionInterface{},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionResult(t *testing.T) {
	const wat = `(module
		(func 
			(param i32 f32)
			(result i32)
			(local f32)
			(unreachable) (unreachable)
		)
	)	
	`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: []types.ValType{
							types.I32,
							types.F32,
						},
						ResultTypes: []types.ValType{
							types.I32,
						},
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionCode:
			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{
							types.SectionCodeLocalElement{
								Size:    1,
								ValType: types.F32,
							},
						},
						Expr: []types.InstructionInterface{
							&types.InstructionSimple{Instruction: types.Unreachable},
							&types.InstructionSimple{Instruction: types.Unreachable},
						},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionLocals(t *testing.T) {
	const wat = `(module
		(func (export "local-mixed") (result f64)
			(local f32 i32 f64)
			(drop (f32.neg (local.get 0)))
			(drop (i32.eqz (local.get 1)))
			(local.get 2)
		 )
	)	
	`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: []types.ValType{},
						ResultTypes: []types.ValType{
							types.F64,
						},
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionExport:
			s := &types.SectionExport{
				Exports: []*types.ExportType{
					&types.ExportType{
						Name:            "local-mixed",
						DescriptionKind: types.ExportDescriptionFunc,
						Index:           0,
					},
				},
			}
			testutil.Assert(t, s, v, "export section does not match")
		case *types.SectionCode:
			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{
							types.SectionCodeLocalElement{
								Size:    1,
								ValType: types.F32,
							},
							types.SectionCodeLocalElement{
								Size:    1,
								ValType: types.I32,
							},
							types.SectionCodeLocalElement{
								Size:    1,
								ValType: types.F64,
							},
						},
						Expr: []types.InstructionInterface{
							&types.InstructionLocalIndex{Instruction: types.LocalGet, Index: 0},
							&types.InstructionSimple{Instruction: types.F32Neg},
							&types.InstructionSimple{Instruction: types.Drop},

							&types.InstructionLocalIndex{Instruction: types.LocalGet, Index: 1},
							&types.InstructionSimple{Instruction: types.I32Eqz},
							&types.InstructionSimple{Instruction: types.Drop},

							&types.InstructionLocalIndex{Instruction: types.LocalGet, Index: 2},
						},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionI32Const(t *testing.T) {
	const wat = `(module
		(func (export "value-i32") (result i32) (i32.const 77))
	)	
	`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: []types.ValType{},
						ResultTypes: []types.ValType{
							types.I32,
						},
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionExport:
			s := &types.SectionExport{
				Exports: []*types.ExportType{
					&types.ExportType{
						Name:            "value-i32",
						DescriptionKind: types.ExportDescriptionFunc,
						Index:           0,
					},
				},
			}
			testutil.Assert(t, s, v, "export section does not match")
		case *types.SectionCode:
			constInstr := &types.InstructionConst{
				Instruction: types.I32Const,
			}
			constInstr.SetInt32(77)

			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{},
						Expr: []types.InstructionInterface{
							constInstr,
						},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionF32Const(t *testing.T) {
	const wat = `(module
		(func (export "value-f32") (result f32) (f32.const 77.7))
	)	
	`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: []types.ValType{},
						ResultTypes: []types.ValType{
							types.F32,
						},
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionExport:
			s := &types.SectionExport{
				Exports: []*types.ExportType{
					&types.ExportType{
						Name:            "value-f32",
						DescriptionKind: types.ExportDescriptionFunc,
						Index:           0,
					},
				},
			}
			testutil.Assert(t, s, v, "export section does not match")
		case *types.SectionCode:
			constInstr := &types.InstructionConst{
				Instruction: types.F32Const,
			}
			constInstr.SetFloat32(77.7)

			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{},
						Expr: []types.InstructionInterface{
							constInstr,
						},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionBlockAndCall(t *testing.T) {
	const wat = `(module
		(func $dummy)
		(func (export "value-block-void") (block (call $dummy) (call $dummy)))
  	)`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: []types.ValType{},
						ResultTypes:    []types.ValType{},
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionExport:
			s := &types.SectionExport{
				Exports: []*types.ExportType{
					&types.ExportType{
						Name:            "value-block-void",
						DescriptionKind: types.ExportDescriptionFunc,
						Index:           1,
					},
				},
			}
			testutil.Assert(t, s, v, "export section does not match")
		case *types.SectionCode:
			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{}, // dummy
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{},
						Expr: []types.InstructionInterface{
							&types.InstructionBlock{
								Instruction: types.Block,
								BlockType:   types.None,
								Instructions: []types.InstructionInterface{
									&types.InstructionFuncIndex{
										Instruction: types.Call,
										Index:       0,
									},
									&types.InstructionFuncIndex{
										Instruction: types.Call,
										Index:       0,
									},
								},
							},
						},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionBranch(t *testing.T) {
	const wat = `(module
		(func (br 0))
	)`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: []types.ValType{},
						ResultTypes:    []types.ValType{},
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionCode:
			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{},
						Expr: []types.InstructionInterface{
							&types.InstructionLabelIndex{
								Instruction: types.Branch,
								Index:       0,
							},
						},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionBranchIf(t *testing.T) {
	const wat = `(module
		(func (param i32)
			(br_if 0 (local.get 0))
		)
	)`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: []types.ValType{
							types.I32,
						},
						ResultTypes: []types.ValType{},
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionCode:
			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{},
						Expr: []types.InstructionInterface{
							&types.InstructionLocalIndex{
								Instruction: types.LocalGet,
								Index:       0,
							},
							&types.InstructionLabelIndex{
								Instruction: types.BranchIf,
								Index:       0,
							},
						},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionBranchIf2(t *testing.T) {
	const wat = `(module
		(func (param i32) (result i32)
			(drop (br_if 0 (i32.const 50) (local.get 0))) (i32.const 51)
		)
	)`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: []types.ValType{
							types.I32,
						},
						ResultTypes: []types.ValType{
							types.I32,
						},
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionCode:
			constInstr50 := &types.InstructionConst{
				Instruction: types.I32Const,
			}
			constInstr50.SetInt32(50)

			constInstr51 := &types.InstructionConst{
				Instruction: types.I32Const,
			}
			constInstr51.SetInt32(51)

			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{},
						Expr: []types.InstructionInterface{
							constInstr50,
							&types.InstructionLocalIndex{
								Instruction: types.LocalGet,
								Index:       0,
							},
							&types.InstructionLabelIndex{
								Instruction: types.BranchIf,
								Index:       0,
							},
							&types.InstructionSimple{
								Instruction: types.Drop,
							},
							constInstr51,
						},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionBranchTable(t *testing.T) {
	const wat = `(module
		(func (param i32)
			(br_table 0 0 0 (local.get 0))
		)
	)`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: []types.ValType{
							types.I32,
						},
						ResultTypes: []types.ValType{},
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionCode:
			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{},
						Expr: []types.InstructionInterface{
							&types.InstructionLocalIndex{
								Instruction: types.LocalGet,
								Index:       0,
							},
							&types.InstructionBranchTable{
								Instruction: types.BranchTable,
								Indices: []types.LabelIndex{
									0, 0,
								},
								Default: 0,
							},
						},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}

func TestFunctionBranchTable2(t *testing.T) {
	const wat = `(module
		(func (param i32) (result i32)
			(i32.add
				(block (result i32)
					(br_table 0 1 0 (i32.const 50) (local.get 0)) (i32.const 51)
				)
				(i32.const 2)
			)
		)
	)`

	wasm := bytes.NewReader(testutil.MustRunWat2Wasm(wat))

	sections, err := ParseBinaryFormat(wasm)

	testutil.AssertNotError(t, err, "failed to parse function")

	for i := range sections {
		switch v := sections[i].(type) {
		case *types.SectionType:
			s := &types.SectionType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: []types.ValType{
							types.I32,
						},
						ResultTypes: []types.ValType{
							types.I32,
						},
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *types.SectionFunction:
			s := &types.SectionFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *types.SectionCode:
			constInstr50 := &types.InstructionConst{
				Instruction: types.I32Const,
			}
			constInstr50.SetInt32(50)

			constInstr51 := &types.InstructionConst{
				Instruction: types.I32Const,
			}
			constInstr51.SetInt32(51)

			constInstr2 := &types.InstructionConst{
				Instruction: types.I32Const,
			}
			constInstr2.SetInt32(2)

			s := &types.SectionCode{
				Codes: []types.SectionCodeElementType{
					types.SectionCodeElementType{
						Locals: []types.SectionCodeLocalElement{},
						Expr: []types.InstructionInterface{
							&types.InstructionBlock{
								Instruction: types.Block,
								BlockType:   types.I32,
								Instructions: []types.InstructionInterface{
									constInstr50,
									&types.InstructionLocalIndex{
										Instruction: types.LocalGet,
										Index:       0,
									},
									&types.InstructionBranchTable{
										Instruction: types.BranchTable,
										Indices:     []types.LabelIndex{0, 1},
										Default:     0,
									},
									constInstr51,
								},
							},

							constInstr2,

							&types.InstructionSimple{
								Instruction: types.I32Add,
							},
						},
					},
				},
			}
			testutil.Assert(t, s, v, "code section does not match")
		}
	}
}
