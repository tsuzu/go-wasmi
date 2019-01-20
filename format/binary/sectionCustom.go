package binary

import (
	"io"
	"io/ioutil"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/pkg/errors"
)

// SectionEntityCustom stores an entity of custom section
type SectionEntityCustom struct {
	Name string
	Data []byte
}

// UnmarshalSectionEntity parses custom section payload
func (s *SectionEntityCustom) UnmarshalSectionEntity(r io.Reader) error {
	name, err := bintypes.ReadString(r)

	if err != nil {
		return errors.WithStack(err)
	}

	data, err := ioutil.ReadAll(r)

	if err != nil {
		return errors.WithStack(err)
	}
	s.Name = name
	s.Data = data

	return nil
}

// SectionID returns wasm section id
func (s *SectionEntityCustom) SectionID() SectionIDType {
	return SectionCustom
}
