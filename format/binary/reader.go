package binary

import (
	"errors"
	"io"
	"log"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
)

const (
	magicNumber   uint32 = 0x00 | 0x61<<8 | 0x73<<16 | 0x6d<<24
	versionNumber uint32 = 0x01
)

var ErrInvalidFormat = errors.New("invalid format")

type Parser struct {
	r io.Reader
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		r: r,
	}
}

func (p *Parser) Parse() error {
	magic, err := binrw.ReadLEUint32(p.r)

	if err != nil {
		return err
	}

	if magic != magicNumber {
		return ErrInvalidFormat
	}

	version, err := binrw.ReadLEUint32(p.r)

	if err != nil {
		return err
	}

	if version != versionNumber {
		return ErrInvalidFormat
	}

	sections := make([]Section, 0, 16)
	for {
		section, err := ReadSection(p.r)

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		sections = append(sections, section)
	}
	return sections, nil
}
