package binary

import (
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"
	"github.com/pkg/errors"
)

// UnmarshalSectionElement parses element section payload
func UnmarshalSectionElement(r io.Reader) (types.Section, error) {
	var s types.SectionElement
	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if s.Elements == nil {
			s.Elements = make([]types.SectionElementElementType, 0, size)
		}

		elm := types.SectionElementElementType{}

		index, err := leb128.ReadUint32(r)
		if err != nil {
			return errors.WithStack(err)
		}
		elm.TableIndex = types.TableIndex(index)

		elm.Expr, err = ReadExpression(r)

		if err != nil {
			return errors.WithStack(err)
		}

		_, err = binrw.ReadVector(r, func(size uint32, r io.Reader) error {
			if elm.FuncIndices == nil {
				elm.FuncIndices = make([]types.FuncIndex, 0, size)
			}

			v, err := leb128.ReadUint32(r)

			if err != nil {
				return errors.WithStack(err)
			}

			elm.FuncIndices = append(elm.FuncIndices, types.FuncIndex(v))

			return nil
		})

		s.Elements = append(s.Elements, elm)

		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &s, nil
}
