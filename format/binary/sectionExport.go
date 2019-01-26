package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"

	"github.com/cs3238-tsuzu/go-wasmi/types"
)

// SectionEntityExport stores an entity of types section
type SectionEntityExport struct {
	ExportTypes []*types.ExportType
}

// UnmarshalSectionEntity parses export section payload
func (s *SectionEntityExport) UnmarshalSectionEntity(r io.Reader) error {
	var exportTypes []*types.ExportType
	_, err := binrw.ReadVector(r, func(size uint32, r io.Reader) error {
		if exportTypes == nil {
			exportTypes = make([]*types.ExportType, 0, size)
		}

		exportType, err := bintypes.ReadExportType(r)

		if err != nil {
			return errors.WithStack(err)
		}

		exportTypes = append(exportTypes, exportType)

		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	s.ExportTypes = exportTypes

	return nil
}

// SectionID returns wasm section id
func (s *SectionEntityExport) SectionID() SectionIDType {
	return SectionExport
}
