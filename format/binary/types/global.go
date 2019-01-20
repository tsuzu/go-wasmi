package bintypes

import (
	"io"

	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

var (
	// ErrInvalidGlobalMutability means invalid global mutability is passed
	ErrInvalidGlobalMutability = errors.New("invalid global mutability")
)

// ReadGlobalMutability parses
func ReadGlobalMutability(r io.Reader) (types.GlobalMutability, error) {
	m, err := binrw.ReadByte(r)

	if err != nil {
		return 0, errors.WithStack(err)
	}

	mut := types.GlobalMutability(m)

	if mut != types.GlobalConst && mut != types.GlobalVariable {
		return 0, errors.WithStack(ErrInvalidGlobalMutability)
	}

	return mut, nil
}

// ReadGlobalType parses global type from reader
func ReadGlobalType(r io.Reader) (*types.GlobalType, error) {
	v, err := ReadValType(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	mut, err := ReadGlobalMutability(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &types.GlobalType{
		Type:       v,
		Mutability: mut,
	}, nil
}
