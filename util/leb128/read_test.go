package leb128_test

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"
	"github.com/pkg/errors"
)

type ReadFuncTypeForUnsigned func(io.Reader) (uint64, error)
type ReadFuncTypeForSigned func(io.Reader) (int64, error)

func testReadUnsigned(fnName string, fn ReadFuncTypeForUnsigned, t *testing.T) {
	from := [][]byte{
		[]byte{0x02},
		[]byte{0x7f},
		[]byte{0x80, 0x01},
		[]byte{0x81, 0x01},
		[]byte{0x82, 0x01},
		[]byte{0xb9, 0x64},
	}
	to := []uint64{
		2,
		127,
		128,
		129,
		130,
		12857,
	}

	for i := range from {
		if v, err := fn(bytes.NewReader(from[i])); err != nil {
			t.Errorf("%s error=> from: %v, to: %d, error: %v", fnName, from[i], to[i], err)
		} else if !reflect.DeepEqual(to[i], v) {
			t.Errorf("%s error=> from: %v, to: %d, expected: %d", fnName, from[i], v, to[i])
		}
	}
}

func testReadSigned(fnName string, fn ReadFuncTypeForSigned, t *testing.T) {
	from := [][]byte{
		[]byte{0x02},
		[]byte{0x7e},
		[]byte{0xff, 0x00},
		[]byte{0x81, 0x7f},
		[]byte{0x80, 0x01},
		[]byte{0x80, 0x7f},
		[]byte{0x81, 0x01},
		[]byte{0xff, 0x7e},
	}
	to := []int64{
		2,
		-2,
		127,
		-127,
		128,
		-128,
		129,
		-129,
	}

	for i := range from {
		if v, err := fn(bytes.NewReader(from[i])); err != nil {
			t.Errorf("%s error=> from: %v, to: %d, error: %v", fnName, from[i], to[i], err)
		} else if !reflect.DeepEqual(to[i], v) {
			t.Errorf("%s error=> from: %v, to: %d, expected: %d", fnName, from[i], v, to[i])
		}
	}
}

func TestReadUint64WithGivenParameters(t *testing.T) {
	fn := func(r io.Reader) (uint64, error) {
		v, err := leb128.ReadUint64(r)
		return uint64(v), err
	}
	testReadUnsigned("ReadUint64", fn, t)
}

func TestReadUint32WithGivenParameters(t *testing.T) {
	fn := func(r io.Reader) (uint64, error) {
		v, err := leb128.ReadUint32(r)
		return uint64(v), err
	}
	testReadUnsigned("ReadUint32", fn, t)
}

func TestReadUintWithGivenParameters(t *testing.T) {
	fn := func(r io.Reader) (uint64, error) {
		v, err := leb128.ReadUint(r)
		return uint64(v), err
	}
	testReadUnsigned("ReadUint", fn, t)
}

func TestReadInt64WithGivenParameters(t *testing.T) {
	fn := func(r io.Reader) (int64, error) {
		v, err := leb128.ReadInt64(r)
		return int64(v), err
	}
	testReadSigned("ReadInt64", fn, t)
}

func TestReadInt32WithGivenParameters(t *testing.T) {
	fn := func(r io.Reader) (int64, error) {
		v, err := leb128.ReadInt32(r)
		return int64(v), err
	}
	testReadSigned("ReadInt32", fn, t)
}

func TestReadIntWithGivenParameters(t *testing.T) {
	fn := func(r io.Reader) (int64, error) {
		v, err := leb128.ReadInt(r)
		return int64(v), err
	}
	testReadSigned("ReadInt", fn, t)
}

func TestReadInt64WithInsufficientData(t *testing.T) {
	b := []byte{0x80}

	_, err := leb128.ReadInt64(bytes.NewReader(b))

	if errors.Cause(err) != io.ErrUnexpectedEOF {
		t.Errorf("error should be unexpected EOF: %v", err)
	}
}

func TestReadInt32WithInsufficientData(t *testing.T) {
	b := []byte{0x80}

	_, err := leb128.ReadInt32(bytes.NewReader(b))

	if errors.Cause(err) != io.ErrUnexpectedEOF {
		t.Errorf("error should be unexpected EOF: %v", err)
	}
}

func TestReadIntWithInsufficientData(t *testing.T) {
	b := []byte{0x80}

	_, err := leb128.ReadInt(bytes.NewReader(b))

	if errors.Cause(err) != io.ErrUnexpectedEOF {
		t.Errorf("error should be unexpected EOF: %v", err)
	}
}
