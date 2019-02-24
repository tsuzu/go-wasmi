package binary

import (
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"
	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

// UnmarshalSectionData parses data section payload
func UnmarshalSectionData(r io.Reader) (types.Section, error) {
	var s types.SectionData
	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if s.Data == nil {
			s.Data = make([]types.SectionDataElement, 0, size)
		}

		elm := types.SectionDataElement{}

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

		s.Data = append(s.Data, elm)

		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &s, nil
}
