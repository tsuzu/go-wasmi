package binary_test

import (
	"os"
	"testing"

	"github.com/k0kubun/pp"

	"github.com/cs3238-tsuzu/go-wasmi/format/binary"
)

func TestParse(t *testing.T) {
	fp, err := os.Open("testdata/sum.wasm")

	if err != nil {
		t.Error("testdata/sum.wasm is missing", err)
	}
	defer fp.Close()

	if sections, err := binary.ParseBinaryFormat(fp); err != nil {
		t.Errorf("parsing sum.wasm error %+v %+v", sections, err)
	} else {
		t.Log(pp.Sprint(sections))
	}
}

func TestParse2(t *testing.T) {
	fp, err := os.Open("testdata/go_sum.wasm")

	if err != nil {
		t.Error("testdata/sum.wasm is missing", err)
	}
	defer fp.Close()

	if sections, err := binary.ParseBinaryFormat(fp); err != nil {
		t.Errorf("parsing sum.wasm error %+v %+v", sections, err)
	} else {
		t.Log(pp.Sprint(sections))
	}
}
