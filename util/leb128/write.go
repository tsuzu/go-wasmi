package leb128

import (
	"io"
	"math"
)

// WriteUint64 writes uint64 value in LSB128 format
func WriteUint64(w io.Writer, v uint64) error {
	return write(w, uint64(v), false)
}

// WriteInt64 writes int64 value in LSB128 format
func WriteInt64(w io.Writer, v int64) error {
	return write(w, uint64(v), true)
}

// WriteUint32 writes uint32 value in LSB128 format
func WriteUint32(w io.Writer, v uint32) error {
	return write(w, uint64(v), false)
}

// WriteInt32 writes int32 value in LSB128 format
func WriteInt32(w io.Writer, v int32) error {
	return write(w, uint64(v), true)
}

// WriteInt writes int value in LSB128 format
func WriteInt(w io.Writer, v int) error {
	return write(w, uint64(v), true)
}

// WriteUint writes uint value in LSB128 format
func WriteUint(w io.Writer, v uint) error {
	return write(w, uint64(v), false)
}

func fillLowerBit(size uint) uint64 {
	return ^(math.MaxUint64 << size)
}

func write(w io.Writer, v uint64, signed bool) error {
	msZeroBit := 63
	plus := false
	if signed {
		if (v & (1 << 63)) == 0 {
			v = ^v
			plus = true
		}

		for msZeroBit >= 0 && (v&(1<<uint(msZeroBit))) != 0 {
			msZeroBit--
		}

		if msZeroBit >= 0 {
			mask := fillLowerBit(uint((msZeroBit+1)/7+1)*7) ^ fillLowerBit(uint(msZeroBit+2))

			v &= fillLowerBit(uint(msZeroBit + 2))
			v |= mask
		}
	}

	buf := make([]byte, 0, 8)
	for v != 0 {
		b := byte(v & 0xff)
		v >>= 7

		if signed && plus {
			b = ^(b | 0x80)
		}
		if v != 0 {
			b |= 0x80
		}

		buf = append(buf, b)
	}

	_, err := w.Write(buf)

	return err
}
