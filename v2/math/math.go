// 目前主要都是針對樣板(text/template, html/template)所設計，也就是對string的類型進行四則運算

package math

import (
	"errors"
	"reflect"
	"strconv"
)

// i2float pretty like cast.ToFloat64E
func i2float(a any) (float64, error) { // interface to number
	aValue := reflect.ValueOf(a)
	switch aValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(aValue.Int()), nil
	case reflect.Float32, reflect.Float64:
		return aValue.Float(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(aValue.Uint()), nil
	/*
	   case reflect.Bool:
	       if a == true {
	           return 1, nil
	       }
	       return 0, nil
	*/
	case reflect.String:
		return strconv.ParseFloat(aValue.String(), 64)
	default:
		return 0, errors.New("type error")
	}
}

func Compute(a, b any, op rune) (float64, error) {
	x, errX := i2float(a)
	y, errY := i2float(b)
	for _, err := range []error{errX, errY} {
		if err != nil {
			return 0, err
		}
	}

	switch op {
	case '+':
		return x + y, nil
	case '-':
		return x - y, nil
	case '*':
		return x * y, nil
	case '/':
		if y == 0 {
			return 0, errors.New("can't divide the value by 0")
		}
		return x / y, nil
	default:
		return 0, errors.New("there is no such an operation")
	}
}
