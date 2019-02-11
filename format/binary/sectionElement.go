package binary

import (
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"
	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

type ElementSectionElement struct {
	TableIndex  types.TableIndex
	Expr        []types.InstructionInterface
	FuncIndices []types.FuncIndex
}

// SectionEntityElement stores an entity of code section
type SectionEntityElement struct {
	Elements []ElementSectionElement
}

// UnmarshalSectionEntity parses element section payload
func (s *SectionEntityElement) UnmarshalSectionEntity(r io.Reader) error {
	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if s.Elements == nil {
			s.Elements = make([]ElementSectionElement, 0, size)
		}

		elm := ElementSectionElement{}

		if index, err := leb128.ReadUint32(r); err != nil {
			return errors.WithStack(err)
		} else {
			elm.TableIndex = types.TableIndex(index)
		}

		var err error
		elm.Expr, err = ReadExpression(r)

		if err != nil {
			return errors.WithStack(err)
		}

		_, err = binrw.ReadVector(r, func(size uint32, r io.Reader) error {
			if elm.FuncIndices == nil {
				elm.FuncIndices = make([]types.FuncIndex, 0, size)
			}

			v, err := leb128.ReadUint32(r)

			if err != nil {
				return errors.WithStack(err)
			}

			elm.FuncIndices = append(elm.FuncIndices, types.FuncIndex(v))

			return nil
		})

		s.Elements = append(s.Elements, elm)

		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// SectionID returns wasm section id
func (s *SectionEntityElement) SectionID() SectionIDType {
	return SectionElement
}
