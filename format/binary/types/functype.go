package bintypes

import (
	"io"

	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

var (
	// ErrIncorrectTypeIndex occurs when incorrect type index is found
	ErrIncorrectTypeIndex = errors.New("incorrect type index")
)

// ReadFuncType parses function type from reader
func ReadFuncType(r io.Reader) (*types.FuncType, error) {
	typePrefix, err := binrw.ReadByte(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if typePrefix != 0x60 {
		return nil, ErrIncorrectTypeIndex
	}

	params, err := ReadVecValType(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	res, err := ReadVecValType(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &types.FuncType{
		ParameterTypes: params,
		ResultTypes:    res,
	}, nil
}
