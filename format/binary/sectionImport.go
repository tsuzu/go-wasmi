package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/pkg/errors"
)

// UnmarshalSectionImport parses import section payload
func UnmarshalSectionImport(r io.Reader) (types.Section, error) {
	var s types.SectionImport

	var importTypes []*types.ImportType
	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if importTypes == nil {
			importTypes = make([]*types.ImportType, 0, size)
		}

		importType, err := bintypes.ReadImportType(r)

		if err != nil {
			return errors.WithStack(err)
		}

		importTypes = append(importTypes, importType)

		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	s.ImportTypes = importTypes

	return &s, nil
}
