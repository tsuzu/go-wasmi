package binary

import (
	"io"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/pkg/errors"
)

// UnmarshalSectionExport parses export section payload
func UnmarshalSectionExport(r io.Reader) (types.Section, error) {
	var s types.SectionExport

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
		return nil, errors.WithStack(err)
	}

	s.Exports = exportTypes

	return &s, nil
}
