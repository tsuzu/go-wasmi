package binary

import (
	"io"

	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
)

const (
	magicNumber   uint32 = 0x00 | 0x61<<8 | 0x73<<16 | 0x6d<<24
	versionNumber uint32 = 0x01
)

// ErrInvalidFormat is an error occurred when data contain invalid format
var ErrInvalidFormat = errors.New("invalid format")

// ParseBinaryFormat parses wasm binary and returns sections
func ParseBinaryFormat(r io.Reader) ([]types.Section, error) {
	magic, err := binrw.ReadLEUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if magic != magicNumber {
		return nil, ErrInvalidFormat
	}

	version, err := binrw.ReadLEUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if version != versionNumber {
		return nil, ErrInvalidFormat
	}

	sections := make([]types.Section, 0, 16)
	for {
		section, err := UnmarshalSection(r)

		if errors.Cause(err) == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		sections = append(sections, section)
	}

	return sections, nil
}
