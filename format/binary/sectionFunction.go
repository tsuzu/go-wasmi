package binary

import (
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/types"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/pkg/errors"
)

// SectionEntityFunction stores an entity of function section
type SectionEntityFunction struct {
	Types []types.TypeIndex
}

// UnmarshalSectionEntity parses function section payload
func (s *SectionEntityFunction) UnmarshalSectionEntity(r io.Reader) error {
	v, err := binrw.ReadVecBytes(r, 1)

	if err != nil {
		return errors.WithStack(err)
	}

	s.Types = make([]types.TypeIndex, 0, len(v))

	for i := range v {
		s.Types = append(s.Types, types.TypeIndex(v[i]))
	}

	return nil
}

// SectionID returns wasm section id
func (s *SectionEntityFunction) SectionID() SectionIDType {
	return SectionFunction
}
