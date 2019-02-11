package binary

import (
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"
	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

// DataSectionElement represents an element in data section
type DataSectionElement struct {
	MemoryIndex types.MemoryIndex
	Expr        []types.InstructionInterface
	Bytes       []byte
}

// SectionEntityData stores an entity of code section
type SectionEntityData struct {
	Elements []DataSectionElement
}

// UnmarshalSectionEntity parses data section payload
func (s *SectionEntityData) UnmarshalSectionEntity(r io.Reader) error {
	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if s.Elements == nil {
			s.Elements = make([]DataSectionElement, 0, size)
		}

		elm := DataSectionElement{}

		index, err := leb128.ReadUint32(r)
		if err != nil {
			return errors.WithStack(err)
		}

		elm.MemoryIndex = types.MemoryIndex(index)

		elm.Expr, err = ReadExpression(r)

		if err != nil {
			return errors.WithStack(err)
		}

		elm.Bytes, err = binrw.ReadVecBytes(r, 1)

		if err != nil {
			return errors.WithStack(err)
		}

		s.Elements = append(s.Elements, elm)

		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// SectionID returns wasm section id
func (s *SectionEntityData) SectionID() SectionIDType {
	return SectionData
}
