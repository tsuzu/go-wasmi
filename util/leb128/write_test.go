package leb128_test

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"
)

type writeFuncTypeForUnsigned func(io.Writer, uint64) error
type writeFuncTypeForSigned func(io.Writer, int64) error

func testWriteUnsigned(fnName string, fn writeFuncTypeForUnsigned, t *testing.T) {
	from := []uint64{
		2,
		127,
		128,
		129,
		130,
		12857,
	}
	to := [][]byte{
		[]byte{0x02},
		[]byte{0x7f},
		[]byte{0x80, 0x01},
		[]byte{0x81, 0x01},
		[]byte{0x82, 0x01},
		[]byte{0xb9, 0x64},
	}

	buf := bytes.NewBuffer(nil)
	for i := range from {
		if err := fn(buf, from[i]); err != nil {
			t.Errorf("%s error=> from: %d, to: %+v, error: %+v", fnName, from[i], to[i], err)
		}
		if !reflect.DeepEqual(to[i], buf.Bytes()) {
			t.Errorf("%s error=> from: %d, to: %+v, expected: %+v", fnName, from[i], buf.Bytes(), to[i])
		}
		buf.Reset()
	}
}

func testWriteSigned(fnName string, fn writeFuncTypeForSigned, t *testing.T) {
	from := []int64{
		2,
		-2,
		127,
		-127,
		128,
		-128,
		129,
		-129,
	}
	to := [][]byte{
		[]byte{0x02},
		[]byte{0x7e},
		[]byte{0xff, 0x00},
		[]byte{0x81, 0x7f},
		[]byte{0x80, 0x01},
		[]byte{0x80, 0x7f},
		[]byte{0x81, 0x01},
		[]byte{0xff, 0x7e},
	}

	buf := bytes.NewBuffer(nil)
	for i := range from {
		if err := fn(buf, from[i]); err != nil {
			t.Errorf("%s error=> from: %d, to: %+v, error: %+v", fnName, from[i], to[i], err)
		}
		if !reflect.DeepEqual(to[i], buf.Bytes()) {
			t.Errorf("%s error=> from: %d, to: %+v, expected: %+v", fnName, from[i], buf.Bytes(), to[i])
		}
		buf.Reset()
	}
}

func TestWriteUint64WithGivenParameters(t *testing.T) {
	fn := func(w io.Writer, v uint64) error {
		return leb128.WriteUint64(w, v)
	}
	testWriteUnsigned("WriteUint64", fn, t)
}

func TestWriteUint32WithGivenParameters(t *testing.T) {
	fn := func(w io.Writer, v uint64) error {
		return leb128.WriteUint32(w, uint32(v))
	}
	testWriteUnsigned("WriteUint32", fn, t)
}

func TestWriteUintWithGivenParameters(t *testing.T) {
	fn := func(w io.Writer, v uint64) error {
		return leb128.WriteUint(w, uint(v))
	}
	testWriteUnsigned("WriteUint", fn, t)
}

func TestWriteInt64WithGivenParameters(t *testing.T) {
	fn := func(w io.Writer, v int64) error {
		return leb128.WriteInt64(w, v)
	}
	testWriteSigned("WriteInt64", fn, t)
}

func TestWriteInt32WithGivenParameters(t *testing.T) {
	fn := func(w io.Writer, v int64) error {
		return leb128.WriteInt32(w, int32(v))
	}
	testWriteSigned("WriteInt32", fn, t)
}

func TestWriteIntWithGivenParameters(t *testing.T) {
	fn := func(w io.Writer, v int64) error {
		return leb128.WriteInt(w, int(v))
	}
	testWriteSigned("WriteInt", fn, t)
}
