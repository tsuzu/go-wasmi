package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

// SectionEntityType stores an entity of types section
type SectionEntityType struct {
	FuncTypes []*types.FuncType
}

// UnmarshalSectionEntity parses custom section payload
func (s *SectionEntityType) UnmarshalSectionEntity(r io.Reader) error {
	var funcTypes []*types.FuncType
	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if funcTypes == nil {
			funcTypes = make([]*types.FuncType, 0, size)
		}

		funcType, err := bintypes.ReadFuncType(r)

		if err != nil {
			return errors.WithStack(err)
		}

		funcTypes = append(funcTypes, funcType)

		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	s.FuncTypes = funcTypes

	return nil
}

// SectionID returns wasm section id
func (s *SectionEntityType) SectionID() SectionIDType {
	return SectionType
}
