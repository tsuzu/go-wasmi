package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

// UnmarshalSectionCode parses code section payload
func UnmarshalSectionCode(r io.Reader) (types.Section, error) {
	var s types.SectionCode

	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if s.Codes == nil {
			s.Codes = make([]types.SectionCodeElementType, 0, size)
		}

		elm := types.SectionCodeElementType{}

		l, err := leb128.ReadUint32(r)

		if err != nil {
			return errors.WithStack(err)
		}

		limited := io.LimitReader(r, int64(l))

		index := 0
		_, err = binrw.ReadVector(limited, func(size uint32, r io.Reader) error {
			if elm.Locals == nil {
				elm.Locals = make([]types.SectionCodeLocalElement, 0, size)
			}

			var local types.SectionCodeLocalElement

			l, err := leb128.ReadUint32(r)
			if err != nil {
				return errors.Wrapf(err, "reading %dth code in code section", index)
			}

			local.Size = l

			v, err := bintypes.ReadValType(r)
			if err != nil {
				return errors.Wrapf(err, "reading %dth code in code section", index)
			}

			local.ValType = v

			elm.Locals = append(elm.Locals, local)

			index++
			return nil
		})

		if err != nil {
			return errors.WithStack(err)
		}

		expr, err := ReadExpression(limited)
		if err != nil {
			return errors.WithStack(err)
		}

		elm.Expr = expr

		buf := make([]byte, 16)
		if l, _ := limited.Read(buf); l != 0 {
			return errors.New("All data in code section is not read")
		}

		s.Codes = append(s.Codes, elm)

		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return s, nil
}

/*// SectionID returns wasm section id
func (s *SectionEntityCode) SectionID() SectionIDType {
	return SectionCodeID
}*/
