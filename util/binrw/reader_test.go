package binrw_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
)

func TestReadLEUint32(t *testing.T) {
	b := []byte{0xf1, 0xf2, 0xf3, 0xf4, 0xf5} // The last element will be ignored

	v, err := binrw.ReadLEUint32(bytes.NewReader(b))

	if err != nil {
		t.Error("ReadLEUint32 unexpected error: ", err)
	}

	var expected uint32
	for i := range b {
		expected += uint32(b[i]) << (uint(i) * 8)
	}

	if v != expected {
		t.Errorf("cannot get expected value=> expected: %d, but got %d", expected, v)
	}
}

func TestReadLEUint64(t *testing.T) {
	b := []byte{0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7, 0xf8, 0xf9} // The last element will be ignores

	v, err := binrw.ReadLEUint64(bytes.NewReader(b))

	if err != nil {
		t.Error("ReadLEUint32 unexpected error: ", err)
	}

	var expected uint64
	for i := range b {
		expected += uint64(b[i]) << (uint(i) * 8)
	}

	if v != expected {
		t.Errorf("cannot get expected value=> expected: %d, but got %d", expected, v)
	}
}

func TestReadByte(t *testing.T) {
	b := []byte{0xf1, 0xf2}

	v, err := binrw.ReadByte(bytes.NewReader(b))

	if err != nil {
		t.Error("ReadByte unexpected error: ", err)
	}

	if b[0] != v {
		t.Errorf("cannot get expected value=> expected: %d, but got %d", b[0], v)
	}
}

func TestUnexpectedEOF(t *testing.T) {
	b := []byte{0xf1, 0xf2, 0xf3}

	_, err := binrw.ReadLEUint32(bytes.NewReader(b))

	if err != io.ErrUnexpectedEOF {
		t.Error("ReadLEUint32: unexpected error should be returned: ", err)
	}

	_, err = binrw.ReadLEUint64(bytes.NewReader(b))

	if err != io.ErrUnexpectedEOF {
		t.Error("ReadLEUint64: unexpected error should be returned: ", err)
	}
}
