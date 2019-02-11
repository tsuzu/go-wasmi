package testutil

import (
	"reflect"
	"runtime"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/k0kubun/pp"
)

func caller() string {
	_, f, l, _ := runtime.Caller(1)

	return f + ":" + strconv.Itoa(l)
}

var (
	// https://godoc.org/github.com/google/go-cmp/cmp#AllowUnexported
	alwaysEqual = cmp.Comparer(func(_, _ interface{}) bool { return true })
	cmpOpt      = cmp.FilterValues(func(x, y interface{}) bool {
		vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
		return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
			(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
			(vx.Len() == 0 && vy.Len() == 0)
	}, alwaysEqual)
)

func setNoColorScheme() {
	pp.SetColorScheme(pp.ColorScheme{
		pp.NoColor,
		pp.NoColor,
		pp.NoColor,
		pp.NoColor,
		pp.NoColor,
		pp.NoColor,
		pp.NoColor,
		pp.NoColor,
		pp.NoColor,
		pp.NoColor,
		pp.NoColor,
		pp.NoColor,
	})
}

// Assert asserts expected == actual
func Assert(t *testing.T, expected, actual interface{}, message string) {
	if diff := cmp.Diff(expected, actual, cmpOpt); diff != "" {
		t.Errorf("%s: expected value is %v, but got %v: %s\n%s", caller(), expected, actual, message, diff)
	}
}

// AssertNotEqual asserts expected != actual
func AssertNotEqual(t *testing.T, expected, actual interface{}, message string) {
	if diff := cmp.Diff(expected, actual, cmpOpt); diff == "" {
		setNoColorScheme()
		t.Errorf("%s: expected value is not %v, but got this: %s", caller(), pp.Sprint(actual), message)
	}
}

// AssertNotError asserts err == nil
func AssertNotError(t *testing.T, err error, message string) {
	if err != nil {
		t.Errorf("%s: error occurred(%+v): %s", caller(), err, message)
	}
}
