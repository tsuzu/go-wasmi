package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/types"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
)

// SectionEntityTable stores an entity of table section
type SectionEntityTable struct {
	Types []*types.TableType
}

// UnmarshalSectionEntity parses table section payload
func (s *SectionEntityTable) UnmarshalSectionEntity(r io.Reader) error {
	var tableTypes []*types.TableType

	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if tableTypes == nil {
			tableTypes = make([]*types.TableType, 0, size)
		}

		t, err := bintypes.ReadTableType(r)

		if err != nil {
			return errors.WithStack(err)
		}

		tableTypes = append(tableTypes, t)
		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	s.Types = tableTypes

	return nil
}

// SectionID returns wasm section id
func (s *SectionEntityTable) SectionID() SectionIDType {
	return SectionTable
}
