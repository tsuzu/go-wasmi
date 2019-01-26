package binary

import (
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"

	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/pkg/errors"
)

// SectionEntityStart stores an entity of start section
type SectionEntityStart struct {
	Index types.FuncIndex
}

// UnmarshalSectionEntity parses start section payload
func (s *SectionEntityStart) UnmarshalSectionEntity(r io.Reader) error {
	index, err := leb128.ReadUint32(r)

	if err != nil {
		return errors.WithStack(err)
	}

	s.Index = types.FuncIndex(index)
	return nil
}

// SectionID returns wasm section id
func (s *SectionEntityStart) SectionID() SectionIDType {
	return SectionStart
}
