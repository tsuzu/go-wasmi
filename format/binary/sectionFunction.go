package binary

import (
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/pkg/errors"
)

// UnmarshalSectionFunction parses function section payload
func UnmarshalSectionFunction(r io.Reader) (types.Section, error) {
	var s types.SectionFunction

	v, err := binrw.ReadVecBytes(r, 1)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	s.Types = make([]types.TypeIndex, 0, len(v))

	for i := range v {
		s.Types = append(s.Types, types.TypeIndex(v[i]))
	}

	return &s, nil
}
