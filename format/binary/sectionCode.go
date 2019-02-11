package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

// CodeSectionLocal represents a set of local variables in a code element
type CodeSectionLocal struct {
	Size     uint32
	ValTypes types.ValType
}

// CodeSectionElement represents an element of code section
type CodeSectionElement struct {
	Locals []CodeSectionLocal
	Expr   []types.InstructionInterface
}

// SectionEntityCode stores an entity of code section
type SectionEntityCode struct {
	Codes []CodeSectionElement
}

// UnmarshalSectionEntity parses code section payload
func (s *SectionEntityCode) UnmarshalSectionEntity(r io.Reader) error {
	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if s.Codes == nil {
			s.Codes = make([]CodeSectionElement, 0, size)
		}

		elm := CodeSectionElement{}

		size, err := leb128.ReadUint32(r)

		if err != nil {
			return errors.WithStack(err)
		}

		limited := io.LimitReader(r, int64(size))

		_, err = binrw.ReadVector(r, func(size uint32, r io.Reader) error {
			if elm.Locals == nil {
				elm.Locals = make([]CodeSectionLocal, 0, size)
			}

			var local CodeSectionLocal

			size, err := leb128.ReadUint32(r)
			if err != nil {
				return errors.WithStack(err)
			}

			local.Size = size

			v, err := bintypes.ReadValType(r)
			if err != nil {
				return errors.WithStack(err)
			}

			local.ValTypes = v

			elm.Locals = append(elm.Locals, local)

			return nil
		})

		if err != nil {
			return errors.WithStack(err)
		}

		expr, err := ReadExpression(limited)
		if err != nil {
			return errors.WithStack(err)
		}

		elm.Expr = expr

		s.Codes = append(s.Codes, elm)

		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// SectionID returns wasm section id
func (s *SectionEntityCode) SectionID() SectionIDType {
	return SectionCode
}
