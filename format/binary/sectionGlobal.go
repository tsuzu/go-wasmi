package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/pkg/errors"
)

// UnmarshalSectionGlobal parses global section payload
func UnmarshalSectionGlobal(r io.Reader) (types.Section, error) {
	var s types.SectionGlobal

	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if s.Globals == nil {
			s.Globals = make([]types.SectionGlobalElementType, 0, size)
		}

		elm := types.SectionGlobalElementType{}

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
		return nil, errors.WithStack(err)
	}

	return &s, nil
}
