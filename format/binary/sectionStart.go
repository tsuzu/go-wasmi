package binary

import (
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"

	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/pkg/errors"
)

// UnmarshalSectionStart parses start section payload
func UnmarshalSectionStart(r io.Reader) (types.Section, error) {
	var s types.SectionStart
	index, err := leb128.ReadUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	s.FuncIndex = types.FuncIndex(index)
	return &s, nil
}
