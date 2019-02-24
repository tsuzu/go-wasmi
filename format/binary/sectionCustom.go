package binary

import (
	"io"
	"io/ioutil"

	bintypes "github.com/cs3238-tsuzu/go-wasmi/format/binary/types"
	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/pkg/errors"
)

// UnmarshalSectionCustom parses custom section payload
func UnmarshalSectionCustom(r io.Reader) (types.Section, error) {
	var s types.SectionCustom
	name, err := bintypes.ReadString(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	data, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	s.Name = name
	s.Data = data

	return &s, nil
}
