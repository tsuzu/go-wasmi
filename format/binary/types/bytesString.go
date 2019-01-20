package bintypes

import (
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/util/binrw"
)

// ReadBytes reads vector(byte) values encoded in wasm binary format
func ReadBytes(r io.Reader) ([]byte, error) {
	return binrw.ReadVecBytes(r, 1)
}

// ReadString reads vector(byte) values encoded in wasm binary format as string
func ReadString(r io.Reader) (string, error) {
	buf, err := ReadBytes(r)

	return string(buf), err
}
