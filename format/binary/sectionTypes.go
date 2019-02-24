package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/pkg/errors"
)

// UnmarshalSectionType parses custom section payload
func UnmarshalSectionType(r io.Reader) (types.Section, error) {
	var s types.SectionType

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
		return nil, errors.WithStack(err)
	}

	s.FuncTypes = funcTypes

	return &s, nil
}
