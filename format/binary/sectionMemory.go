package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

// SectionEntityMemory stores an entity of types section
type SectionEntityMemory struct {
	MemoryTypes []*types.MemoryType
}

// UnmarshalSectionEntity parses memory section payload
func (s *SectionEntityMemory) UnmarshalSectionEntity(r io.Reader) error {
	var memoryTypes []*types.MemoryType
	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if memoryTypes == nil {
			memoryTypes = make([]*types.MemoryType, 0, size)
		}

		memoryType, err := bintypes.ReadMemoryType(r)

		if err != nil {
			return errors.WithStack(err)
		}

		memoryTypes = append(memoryTypes, memoryType)

		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	s.MemoryTypes = memoryTypes
	return nil
}

// SectionID returns wasm section id
func (s *SectionEntityMemory) SectionID() SectionIDType {
	return SectionMemory
}
