package bintypes

import (
	"io"

	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
)

var (
	// ErrInvalidElemType means your table element type is not funcref
	ErrInvalidElemType = errors.New("invalid element type for table")
)

// ReadTableType parses table type encoded in wasm binary format
func ReadTableType(r io.Reader) (*types.TableType, error) {
	b, err := binrw.ReadByte(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if types.TableElemType(b) != types.TableElemFuncRef {
		return nil, ErrInvalidElemType
	}

	limits, err := ReadLimits(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &types.TableType{
		ElemType: types.TableElemType(b),
		Limits:   *limits,
	}, nil
}
