package binrw_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"

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
	b := []byte{0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7, 0xf8, 0xf9} // The last element will be ignored

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

	if errors.Cause(err) != io.ErrUnexpectedEOF {
		t.Error("ReadLEUint32: unexpected error should be returned: ", err)
	}

	_, err = binrw.ReadLEUint64(bytes.NewReader(b))

	if errors.Cause(err) != io.ErrUnexpectedEOF {
		t.Error("ReadLEUint64: unexpected error should be returned: ", err)
	}
}

func TestReadVectorSuccess(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	const size = 5
	const message = "hello"
	leb128.WriteUint32(buf, size)

	for i := 0; i < 10; i++ {
		buf.Write([]byte(message))
	}

	cnt := 0
	s, err := binrw.ReadVector(buf, func(s uint32, r io.Reader) error {
		if size != s {
			t.Errorf("size must be %d, but got %d", size, s)
		}

		b, err := ioutil.ReadAll(io.LimitReader(r, 5))

		if err != nil {
			t.Errorf("reading for each element error: %d times, %+v", cnt, err)
		}

		if string(b) != message {
			t.Errorf("element must be %s, but got %s", message, string(b))
		}

		cnt++

		return nil
	})

	if err != nil {
		t.Errorf("ReadVector error: %+v", err)
	}

	if s != size {
		t.Errorf("size must be %d, but got %d", size, s)
	}
}

func TestReadVectorLEB128Error(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	_, err := binrw.ReadVector(buf, func(s uint32, r io.Reader) error {
		t.Error("callback must not be called")

		return nil
	})

	if errors.Cause(err) != io.ErrUnexpectedEOF {
		t.Errorf("ReadVector must return unexpected EOF, but returned %+v", err)
	}
}

func TestReadVecBytesSuccess(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	payload := bytes.NewBuffer(nil)

	const size = 10

	leb128.WriteUint32(buf, size)

	b := []byte{0x00, 0x01, 0x02, 0x03, 0x04}
	for i := 0; i < size; i++ {
		payload.Write(b)

		for i := range b {
			b[i] += 0x05
		}
	}
	buf.Write(payload.Bytes())

	res, err := binrw.ReadVecBytes(buf, len(b))

	if err != nil {
		t.Errorf("ReadVecBytes error: %+v", err)
	}

	if !reflect.DeepEqual(payload.Bytes(), res) {
		t.Errorf("data is not equal: expected=>%+v, actual=>%+v", payload.Bytes(), res)
	}
}
