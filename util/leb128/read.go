package leb128

import "io"

// ReadUint64 reads uint64 value encoded in LEB128 from reader
func ReadUint64(r io.Reader) (uint64, error) {
	v, _, err := readLEB128(r)

	return v, err
}

// ReadInt64 reads int64 value encoded in LEB128 from reader
func ReadInt64(r io.Reader) (int64, error) {
	v, s, err := readLEB128(r)

	if err != nil {
		return 0, err
	}

	res := int64(v)
	if (v & (uint64(1) << (s - 1))) != 0 {
		res |= int64(-1) << s
	}

	return res, nil
}

// ReadUint32 reads uint32 value encoded in LEB128 from reader
func ReadUint32(r io.Reader) (uint32, error) {
	v, _, err := readLEB128(r)

	return uint32(v), err
}

// ReadInt32 reads int32 value encoded in LEB128 from reader
func ReadInt32(r io.Reader) (int32, error) {
	v, s, err := readLEB128(r)

	if err != nil {
		return 0, err
	}

	res := int32(v)
	if (v & (uint64(1) << (s - 1))) != 0 {
		res |= int32(-1) << s
	}

	return res, nil
}

// ReadInt reads int(4 bytes) value encoded in LEB128 from reader
func ReadInt(r io.Reader) (int, error) {
	v, err := ReadInt32(r)

	return int(v), err
}

// ReadUint reads uint(4 bytes) value encoded in LEB128 from reader
func ReadUint(r io.Reader) (uint, error) {
	v, err := ReadUint32(r)

	return uint(v), err
}

func readLEB128(r io.Reader) (uint64, uint, error) {
	var res uint64
	var shift uint
	arr := make([]byte, 1)

	for {
		l, err := r.Read(arr)

		if l == 0 && err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}

			return 0, 0, err
		}

		res |= uint64(arr[0]&0x7f) << shift
		shift += 7

		if (0x80 & arr[0]) == 0 {
			break
		}
	}

	return res, shift, nil
}
