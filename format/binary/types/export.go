package bintypes

import (
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"

	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

// ReadExportType parses export type from reader
func ReadExportType(r io.Reader) (*types.ExportType, error) {
	name, err := ReadString(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	descType, err := binrw.ReadByte(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	imp := &types.ExportType{
		Name:            name,
		DescriptionKind: types.ExportDescriptionKind(descType),
	}

	index, err := leb128.ReadUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	imp.Index = index

	switch imp.DescriptionKind {
	case types.ExportDescriptionFunc:
	case types.ExportDescriptionTable:
	case types.ExportDescriptionMemory:
	case types.ExportDescriptionGlobal:
	default:
		return nil, ErrIncorrectTypeIndex
	}

	return imp, nil
}
