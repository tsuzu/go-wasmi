package bintypes

import (
	"io"

	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

var (
	// ErrInvalidValType occurs when invalid value type index is passed
	ErrInvalidValType = errors.New("invalid value type")
)

// ReadValType reads a value type from reader
func ReadValType(r io.Reader) (types.ValType, error) {
	b, err := binrw.ReadByte(r)

	if err != nil {
		return 0, err
	}

	for i := range types.ValTypes {
		if byte(types.ValTypes[i]) == b {
			return types.ValType(b), nil
		}
	}

	return 0, ErrInvalidValType
}

// ReadVecValType reads a vector of value types from reader
func ReadVecValType(r io.Reader) ([]types.ValType, error) {
	vec, err := binrw.ReadVecBytes(r, 1)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	valTypes := make([]types.ValType, 0, len(vec))

	for i := range vec {
		flag := false
		for j := range types.ValTypes {
			if byte(types.ValTypes[j]) == vec[i] {
				flag = true

				break
			}
		}

		if !flag {
			return nil, ErrInvalidValType
		}

		valTypes = append(valTypes, types.ValType(vec[i]))
	}

	return valTypes, nil
}
