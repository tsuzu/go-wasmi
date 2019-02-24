package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/pkg/errors"
)

// UnmarshalSectionMemory parses memory section payload
func UnmarshalSectionMemory(r io.Reader) (types.Section, error) {
	var s types.SectionMemory

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
		return nil, errors.WithStack(err)
	}

	s.Types = memoryTypes
	return &s, nil
}
