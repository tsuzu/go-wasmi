package bintypes

import (
	"io"

	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

var (
	// ErrInvalidLimitsType occurs when limits type is neither 0x01 nor 0x01
	ErrInvalidLimitsType = errors.New("unknown limits type")
)

// ReadLimits parses limits from reader
func ReadLimits(r io.Reader) (*types.Limits, error) {
	b, err := binrw.ReadByte(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	switch b {
	case 0x00:
		min, err := leb128.ReadUint32(r)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		return &types.Limits{
			Min:       min,
			IgnoreMax: true,
		}, nil
	case 0x01:
		min, err := leb128.ReadUint32(r)

		if err != nil {
			return nil, errors.WithStack(err)
		}
		max, err := leb128.ReadUint32(r)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		return &types.Limits{
			Min:       min,
			Max:       max,
			IgnoreMax: false,
		}, nil
	default:
		return nil, ErrInvalidLimitsType
	}
}
