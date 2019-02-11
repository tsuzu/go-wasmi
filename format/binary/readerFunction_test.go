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
		switch v := sections[i].Entity.(type) {
		case *SectionEntityType:
			s := &SectionEntityType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: make([]types.ValType, 0),
						ResultTypes:    make([]types.ValType, 0),
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *SectionEntityFunction:
			s := &SectionEntityFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *SectionEntityCode:
			s := &SectionEntityCode{
				Codes: []CodeSectionElement{
					CodeSectionElement{
						Locals: []CodeSectionLocal{},
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
		switch v := sections[i].Entity.(type) {
		case *SectionEntityType:
			s := &SectionEntityType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: make([]types.ValType, 0),
						ResultTypes:    make([]types.ValType, 0),
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *SectionEntityFunction:
			s := &SectionEntityFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *SectionEntityCode:
			s := &SectionEntityCode{
				Codes: []CodeSectionElement{
					CodeSectionElement{
						Locals: []CodeSectionLocal{
							CodeSectionLocal{
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
		switch v := sections[i].Entity.(type) {
		case *SectionEntityType:
			s := &SectionEntityType{
				FuncTypes: []*types.FuncType{
					&types.FuncType{
						ParameterTypes: make([]types.ValType, 0),
						ResultTypes:    make([]types.ValType, 0),
					},
				},
			}
			testutil.Assert(t, s, v, "type section does not match")
		case *SectionEntityFunction:
			s := &SectionEntityFunction{
				Types: []types.TypeIndex{
					0,
				},
			}
			testutil.Assert(t, s, v, "function section does not match")
		case *SectionEntityCode:
			s := &SectionEntityCode{
				Codes: []CodeSectionElement{
					CodeSectionElement{
						Locals: []CodeSectionLocal{
							CodeSectionLocal{
								Size:    1,
								ValType: types.I32,
							},
							CodeSectionLocal{
								Size:    1,
								ValType: types.F32,
							},
							CodeSectionLocal{
								Size:    1,
								ValType: types.I64,
							},
							CodeSectionLocal{
								Size:    1,
								ValType: types.I32,
							},
							CodeSectionLocal{
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
