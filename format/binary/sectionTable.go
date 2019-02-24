package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/pkg/errors"
)

// UnmarshalSectionTable parses table section payload
func UnmarshalSectionTable(r io.Reader) (types.Section, error) {
	var s types.SectionTable

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
		return nil, errors.WithStack(err)
	}

	s.Types = tableTypes

	return &s, nil
}
