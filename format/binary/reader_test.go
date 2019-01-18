package binary_test

import (
	"github.com/cs3238-tsuzu/go-wasmi/format/binary"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	fp, err := os.Open("testdata/sum.wasm")

	if err != nil {
		t.Error("testdata/sum.wasm is missing", err)
	}
	defer fp.Close()

	if err := binary.NewParser(fp).Parse(); err != nil {
		t.Error("parsing sum.wasm error", err)
	}
}
