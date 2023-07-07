/*
go test -v -coverprofile="coverage.txt"
go tool cover -func="coverage.txt"
go tool cover -html="coverage.txt"
*/

package strings_test

import (
	"errors"
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/strings"
	"strconv"
	"testing"
)

func TestCmpVer(t *testing.T) {
	for _, d := range []struct {
		v1       string
		op       string
		v2       string
		expected any // bool, error
	}{
		{"1.20", ">", "1.19.5", true},
		{"1.20", ">=", "1.19.5", true},
		{"1.2.3.20230707", ">=", "1.2.4", false},

		{"1.01", "==", "1.001", true},
		{"1.0", "==", "1.0.0", true},
		{"1.0", "=", "1.0.0", true},

		{"1.1", "!=", "1.2", true},

		{"0.1", "<", "1.1", true},
		{"0.1", "<=", "1.1", true},

		{"1.20", "@", "1.19.5", strings.ErrInvalidOperator},
		{"one", ">=", "2", strconv.ErrSyntax},
		{"1", ">=", "two", strconv.ErrSyntax},
	} {
		actual, err := strings.CmpVer(d.v1, d.op, d.v2)
		if err == nil {
			if expected := d.expected.(bool); actual != expected {
				t.Error(d.v1, d.op, d.v2, actual, expected)
				continue
			}
		} else if !errors.Is(err, d.expected.(error)) {
			t.Error("should error", d.v1, d.op, d.v2, err, d.expected)
		}
	}
}

func ExampleCmpVer() {
	for _, d := range []struct {
		v1 string
		op string
		v2 string
	}{
		{"1.01", "==", "1.001"},
		{"1.0", "==", "1.0.0"},
		{"0.1", "<", "1.1"},
		{"1.20", ">", "1.19.5"},
	} {
		result, _ := strings.CmpVer(d.v1, d.op, d.v2)
		fmt.Println(d.v1, d.op, d.v2, result)
	}

	// Output:
	// 1.01 == 1.001 true
	// 1.0 == 1.0.0 true
	// 0.1 < 1.1 true
	// 1.20 > 1.19.5 true
}

func TestMustCmpVer(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("should get an error")
		} else if !errors.Is(err.(error), strconv.ErrSyntax) {
			t.Error("error type != strconv.ErrSyntax")
		}
	}()
	strings.MustCmpVer("one", "==", "2")
}

func ExampleMustCmpVer() {
	fmt.Println(strings.MustCmpVer("1.01", "==", "1.001"))
	// Output:
	// true
}
