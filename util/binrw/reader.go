package binrw

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"

	"github.com/cs3238-tsuzu/go-wasmi/util/leb128"
)

// ReadLEUint32 reads uint32 value encoded in little endian
func ReadLEUint32(r io.Reader) (uint32, error) {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(buf), nil
}

// ReadLEUint64 reads uint64 value encoded in little endian
func ReadLEUint64(r io.Reader) (uint64, error) {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(buf), nil
}

// ReadByte reads a byte value from reader
func ReadByte(r io.Reader) (byte, error) {
	buf := make([]byte, 1)
	_, err := io.ReadFull(r, buf)

	return buf[0], err
}

// ReadVector parses vectors in wasm binary format and calls fn for each element
func ReadVector(r io.Reader, fn func(uint32, io.Reader) error) (uint32, error) {
	size, err := leb128.ReadUint32(r)

	if err != nil {
		return 0, err
	}

	for i := uint32(0); i < size; i++ {
		err := fn(size, r)

		if err != nil {
			if errors.Cause(err) != io.EOF {
				return 0, err
			}
			if i != size-1 {
				err = io.ErrUnexpectedEOF
				return 0, err
			}
		}
	}

	return size, nil
}

// ReadVecBytes read bytes as long as len(x) * elemSize
func ReadVecBytes(r io.Reader, elemSize int) ([]byte, error) {
	size, err := leb128.ReadUint32(r)

	if err != nil {
		return nil, err
	}

	buf := make([]byte, size*uint32(elemSize))

	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}

	return buf, nil
}
