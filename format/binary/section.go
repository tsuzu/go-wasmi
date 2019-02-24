package binary

import (
	"io"
	"io/ioutil"

	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"
)

// SectionIDType is a type for wasm section id
type SectionIDType byte

const (
	//SectionCustomID : custom section
	SectionCustomID SectionIDType = iota
	//SectionTypeID : type section
	SectionTypeID
	//SectionImportID : import section
	SectionImportID
	//SectionFunctionID : function section
	SectionFunctionID
	//SectionTableID : table section
	SectionTableID
	//SectionMemoryID : memory section
	SectionMemoryID
	//SectionGlobalID : global section
	SectionGlobalID
	//SectionExportID : export section
	SectionExportID
	//SectionStartID : start section
	SectionStartID
	//SectionElementID : element section
	SectionElementID
	//SectionCodeID : code section
	SectionCodeID
	//SectionDataID : data section
	SectionDataID
)

var (
	// ErrUnknownSection is an error occurred when section id is unknown
	ErrUnknownSection = errors.New("unknown section")

	// ErrExcessOfBytesInSection is returned when all data are not read in a section
	ErrExcessOfBytesInSection = errors.New("excess of bytes in section")
)

// SectionEntity is an interface for each kind of sections
type SectionEntity interface {
	// UnmarshalSectionEntity parses section payload
	UnmarshalSectionEntity(io.Reader) error
	// SectionID returns wasm section id
	SectionID() SectionIDType
}

// UnmarshalSection parses wasm section
func UnmarshalSection(r io.Reader) (types.Section, error) {
	kindByte, err := binrw.ReadByte(r)

	if err != nil {
		return nil, err
	}

	kind := SectionIDType(kindByte)

	size, err := leb128.ReadUint32(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var parser func(io.Reader) (types.Section, error)
	switch kind {
	case SectionCustomID:
		parser = UnmarshalSectionCustom
	case SectionTypeID:
		parser = UnmarshalSectionType
	case SectionImportID:
		parser = UnmarshalSectionImport
	case SectionFunctionID:
		parser = UnmarshalSectionFunction
	case SectionTableID:
		parser = UnmarshalSectionTable
	case SectionMemoryID:
		parser = UnmarshalSectionMemory
	case SectionGlobalID:
		parser = UnmarshalSectionGlobal
	case SectionExportID:
		parser = UnmarshalSectionExport
	case SectionStartID:
		parser = UnmarshalSectionStart
	case SectionElementID:
		parser = UnmarshalSectionElement
	case SectionCodeID:
		parser = UnmarshalSectionCode
	case SectionDataID:
		parser = UnmarshalSectionData
	default:
		return nil, ErrUnknownSection
	}

	limited := io.LimitReader(r, int64(size))

	s, err := parser(limited)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if b, _ := ioutil.ReadAll(limited); len(b) != 0 {
		return nil, ErrExcessOfBytesInSection
	}

	return s, nil
}
