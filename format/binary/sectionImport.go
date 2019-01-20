package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

// SectionEntityImport stores an entity of types section
type SectionEntityImport struct {
	ImportTypes []*types.ImportType
}

// UnmarshalSectionEntity parses custom section payload
func (s *SectionEntityImport) UnmarshalSectionEntity(r io.Reader) error {
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
		return errors.WithStack(err)
	}

	s.ImportTypes = importTypes

	return nil
}

// SectionID returns wasm section id
func (s *SectionEntityImport) SectionID() SectionIDType {
	return SectionImport
}
