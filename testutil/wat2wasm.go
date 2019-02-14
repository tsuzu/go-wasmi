package testutil

import (
	"io"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// RunWat2Wasm executes wat2wasm for tests
func RunWat2Wasm(wat string) ([]byte, error) {
	in, err := ioutil.TempFile("", "go_wasmi_test_in")

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if _, err := io.Copy(in, strings.NewReader(wat)); err != nil {
		return nil, errors.WithStack(err)
	}

	inName := in.Name()
	in.Close()

	out, err := ioutil.TempFile("", "go_wasmi_test_out")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	outName := out.Name()
	out.Close()

	if b, err := exec.Command("wat2wasm", inName, "-o", outName).CombinedOutput(); err != nil {
		return nil, errors.Wrap(err, string(b))
	}

	b, err := ioutil.ReadFile(outName)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return b, nil
}

// MustRunWat2Wasm executes wat2wasm for tests
// if errors occur, this will panic
func MustRunWat2Wasm(wat string) []byte {
	b, err := RunWat2Wasm(wat)

	if err != nil {
		panic(err)
	}

	return b
}
