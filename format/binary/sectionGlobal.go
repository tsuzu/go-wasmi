package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

// GlobalSectionElement represents an element in global section
type GlobalSectionElement struct {
	Type types.GlobalType
	Expr []types.InstructionInterface
}

// SectionEntityGlobal stores an entity of code section
type SectionEntityGlobal struct {
	Globals []GlobalSectionElement
}

// UnmarshalSectionEntity parses global section payload
func (s *SectionEntityGlobal) UnmarshalSectionEntity(r io.Reader) error {
	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if s.Globals == nil {
			s.Globals = make([]GlobalSectionElement, 0, size)
		}

		elm := GlobalSectionElement{}

		t, err := bintypes.ReadGlobalType(r)
		if err != nil {
			return errors.WithStack(err)
		}

		elm.Type = *t

		elm.Expr, err = ReadExpression(r)

		if err != nil {
			return errors.WithStack(err)
		}

		s.Globals = append(s.Globals, elm)

		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// SectionID returns wasm section id
func (s *SectionEntityGlobal) SectionID() SectionIDType {
	return SectionGlobal
}
