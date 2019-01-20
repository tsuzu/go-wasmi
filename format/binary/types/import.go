package bintypes

import (
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"
	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

// ReadImportType parses import type from reader
func ReadImportType(r io.Reader) (*types.ImportType, error) {
	module, err := ReadString(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	name, err := ReadString(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	descType, err := binrw.ReadByte(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	imp := &types.ImportType{
		Module:          module,
		Name:            name,
		DescriptionKind: types.ImportDescriptionKind(descType),
	}

	switch descType {
	case 0x00:
		v, err := leb128.ReadUint32(r)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		imp.Type = types.TypeIndex(v)

	case 0x01:
		imp.Table, err = ReadTableType(r)

		if err != nil {
			return nil, errors.WithStack(err)
		}
	case 0x02:
		imp.Memory, err = ReadMemory(r)

		if err != nil {
			return nil, errors.WithStack(err)
		}
	case 0x03:
		imp.Global, err = ReadGlobalType(r)

		if err != nil {
			return nil, errors.WithStack(err)
		}
	default:
		return nil, ErrIncorrectTypeIndex
	}

	return imp, nil
}
