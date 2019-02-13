package binary_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/k0kubun/pp"

	"github.com/cs3238-tsuzu/go-wasmi/format/binary"
	"github.com/cs3238-tsuzu/go-wasmi/testutil"
)

func hexBytes(b []byte) string {
	res := []string{}
	for i := range b {
		res = append(res, fmt.Sprintf("%02x", int64(b[i])))
	}

	return strings.Join(res, " ")
}

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
		t.Error("testdata/go_sum.wasm is missing", err)
	}
	defer fp.Close()

	buf := bytes.NewBuffer(nil)
	tee := io.TeeReader(fp, buf)

	if sections, err := binary.ParseBinaryFormat(tee); err != nil {
		bytes := buf.Bytes()
		i := len(bytes) - 50

		if i < 0 {
			i = 0
		}

		t.Errorf("parsing go_sum.wasm error %+v %+v when reading about %d bytes(%v)", sections, err, buf.Len(), bytes[i:])

	}
}

type mockWriter struct {
	cnt int
}

func (w *mockWriter) Write(b []byte) (int, error) {
	w.cnt += len(b)

	//log.Println(w.cnt)

	return len(b), nil
}

func TestParseEmpty(t *testing.T) {
	fp, err := os.Open("testdata/go_empty.wasm")

	if err != nil {
		t.Error("testdata/go_empty.wasm is missing", err)
	}
	defer fp.Close()

	if sections, err := binary.ParseBinaryFormat(fp); err != nil {
		t.Errorf("parsing go_empty.wasm error %+v %+v when reading about d bytes(s)", sections, err)
	}
}

// splitSExpression ignores "" and ''
func splitSExpression(r io.Reader) ([]string, error) {
	b, err := ioutil.ReadAll(r) // 面倒なので許して

	if err != nil {
		return nil, err
	}

	begin, stat := 0, 0
	res := []string{}
	str := string(b)
	for i, c := range str {
		if c == '(' {
			stat++

			if stat == 1 {
				begin = i
			}
		}
		if c == ')' {
			stat--

			if stat == 0 {
				res = append(res, str[begin:i+1])
			}
		}
	}

	return res, nil
}

func TestParseSpecTest(t *testing.T) {
	filepath.Walk("./testdata/from_spec/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".wast" {
			return nil
		}

		fp, err := os.Open(path)

		testutil.AssertNotError(t, err, fmt.Sprintf("failed to open the file %s", path))

		defer fp.Close()

		res, err := splitSExpression(fp)

		testutil.AssertNotError(t, err, fmt.Sprintf("failed to parse s-expr %s", path))

		for i, wast := range res {
			if !strings.HasPrefix(wast, "(module") {
				continue
			}

			wasmb, err := testutil.RunWat2Wasm(wast)

			if err != nil {
				t.Logf("parse wast error %s:%d %+v", path, i, err)
				continue
			}

			wasm := bytes.NewReader(wasmb)

			_, err = binary.ParseBinaryFormat(wasm)

			testutil.AssertNotError(t, err, "failed to parse function")
		}

		return nil
	})
}
