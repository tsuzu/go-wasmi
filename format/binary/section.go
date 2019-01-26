package binary

import (
	"io"
	"io/ioutil"

	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"
)

// SectionIDType is a type for wasm section id
type SectionIDType byte

const (
	//SectionCustom : custom section
	SectionCustom SectionIDType = iota
	//SectionType : type section
	SectionType
	//SectionImport : import section
	SectionImport
	//SectionFunction : function section
	SectionFunction
	//SectionTable : table section
	SectionTable
	//SectionMemory : memory section
	SectionMemory
	//SectionGlobal : global section
	SectionGlobal
	//SectionExport : export section
	SectionExport
	//SectionStart : start section
	SectionStart
	//SectionElement : element section
	SectionElement
	//SectionCode : code section
	SectionCode
	//SectionData : data section
	SectionData
)

var (
	// ErrUnknownSection is an error occurred when section id is unknown
	ErrUnknownSection = errors.New("unknown section")
)

// SectionEntity is an interface for each kind of sections
type SectionEntity interface {
	// UnmarshalSectionEntity parses section payload
	UnmarshalSectionEntity(io.Reader) error
	// SectionID returns wasm section id
	SectionID() SectionIDType
}

// Section stores wasm section
type Section struct {
	Entity SectionEntity
}

// UnmarshalSection parses wasm section
func (s *Section) UnmarshalSection(r io.Reader) error {
	kindByte, err := binrw.ReadByte(r)

	if err != nil {
		return errors.WithStack(err)
	}

	kind := SectionIDType(kindByte)

	size, err := leb128.ReadUint32(r)

	if err != nil {
		return errors.WithStack(err)
	}

	limited := io.LimitReader(r, int64(size))

	var entity SectionEntity
	switch kind {
	case SectionCustom:
		entity = &SectionEntityCustom{}
	case SectionType:
		entity = &SectionEntityType{}
	case SectionImport:
		entity = &SectionEntityImport{}
	case SectionFunction:
		entity = &SectionEntityFunction{}
	case SectionTable:
		entity = &SectionEntityTable{}
	case SectionMemory:
		entity = &SectionEntityMemory{}
	//case SectionGlobal:
	case SectionExport:
		entity = &SectionEntityExport{}

	default:
		// TODO: For debugging, but must be removed
		ioutil.ReadAll(limited)
		return nil
		return ErrUnknownSection
	}

	if err := entity.UnmarshalSectionEntity(limited); err != nil {
		return errors.WithStack(err)
	}

	if b, _ := ioutil.ReadAll(limited); len(b) != 0 {
		panic(errors.Errorf("All data in this section must be loaded %v %d", kind, len(b)))
	}

	s.Entity = entity
	return nil
}
