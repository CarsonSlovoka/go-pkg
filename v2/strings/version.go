package strings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidOperator = errors.New("invalid operator")
)

func CmpVer(v1, op, v2 string) (bool, error) {
	v1Slice := strings.Split(v1, ".")
	v2Slice := strings.Split(v2, ".")

	var maxSize int
	// Let each other have the same length.
	if len(v1Slice) < len(v2Slice) {
		maxSize = len(v2Slice)
	} else {
		maxSize = len(v1Slice)
	}

	v1NSlice := make([]int, maxSize)
	v2NSlice := make([]int, maxSize)

	var err error
	// convert string to int
	for i := range v1Slice {
		v1NSlice[i], err = strconv.Atoi(v1Slice[i])
		if err != nil {
			return false, err
		}
	}
	for i := range v2Slice {
		v2NSlice[i], err = strconv.Atoi(v2Slice[i])
		if err != nil {
			return false, err
		}
	}

	var sign rune
	var v2Elem int
	for i, v1Elem := range v1NSlice {
		if sign != '=' && sign != 0 { //  // The previous comparison has got the answer already.
			break
		}

		v2Elem = v2NSlice[i]
		if v1Elem > v2Elem {
			sign = '>'
		} else if v1Elem < v2Elem {
			sign = '<'
		} else {
			sign = '='
		}
	}

	switch op {
	case ">":
		return sign == '>', nil
	case ">=":
		return sign == '>' || sign == '=', nil
	case "==":
		fallthrough
	case "=":
		return sign == '=', nil
	case "!=":
		return sign != '=', nil
	case "<":
		return sign == '<', nil
	case "<=":
		return sign == '<' || sign == '=', nil
	default:
		return false, fmt.Errorf("%w: %s", ErrInvalidOperator, op)
	}
}

func MustCmpVer(v1, op, v2 string) bool {
	r, err := CmpVer(v1, op, v2)
	if err != nil {
		panic(err)
	}
	return r
}
