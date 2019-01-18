package binary

import (
	"errors"
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"
)

type SectionIDType byte

const (
	SectionCustom SectionIDType = iota
	SectionType
	SectionImport
	SectionFunction
	SectionTable
	SectionMemory
	SectionGlobal
	SectionExport
	SectionStart
	SectionElement
	SectionCode
	SectionData
	SectionIDMax
)

var (
	ErrUnknownSection = errors.New("unknown section")
)

type Section struct {
	Kind SectionIDType
	Data []byte
}

func ReadSection(r io.Reader) (*Section, error) {
	kindByte, err := binrw.ReadByte(r)

	if err != nil {
		return nil, err
	}

	kind := SectionIDType(kindByte)

	if kind >= SectionIDMax {
		return nil, ErrUnknownSection
	}

	size, err := leb128.ReadUint32(r)

	buf := make([]byte, size)

	if _, err := io.ReadFull(r, buf); err != nil && err != io.EOF {
		return nil, err
	}

	return &Section{
		Kind: kind,
		Data: buf,
	}, nil
}
