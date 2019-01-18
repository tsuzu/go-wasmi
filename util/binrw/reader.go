package binrw

import "encoding/binary"
import "io"

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
