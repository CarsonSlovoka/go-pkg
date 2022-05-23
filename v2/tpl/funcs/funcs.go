package funcs

import (
	"errors"
	"github.com/CarsonSlovoka/go-pkg/v2/spf13/cast"
	"reflect"
	"strings"
)

func Dict(values ...any) (map[string]any, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("parameters must be even")
	}
	dict := make(map[string]any)
	var key, val any
	for {
		key, val, values = values[0], values[1], values[2:]
		switch reflect.ValueOf(key).Kind() {
		case reflect.String:
			dict[key.(string)] = val
		default:
			return nil, errors.New(`type must equal to "string"`)
		}
		if len(values) == 0 {
			break
		}
	}
	return dict, nil
}

func List(args ...any) any {
	if len(args) == 0 {
		return args
	}
	first := args[0]
	firstType := reflect.TypeOf(first)

	if firstType == nil {
		return args
	}

	if len(args) > 1 {
		// make sure all type are the same as the first
		for i := 1; i < len(args); i++ {
			if firstType != reflect.TypeOf(args[i]) {
				return args
			}
		}
	}

	// if t represents int, SliceOf(t) represents []int.
	slice := reflect.MakeSlice(reflect.SliceOf(firstType), len(args), len(args))
	for i, arg := range args {
		slice.Index(i).Set(reflect.ValueOf(arg))
	}
	return slice.Interface() // {interface() | []yourType
}

func Split(obj any, delimiter string) ([]string, error) {
	objStr, err := cast.ToStringE(obj)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(objStr, delimiter), nil
}

func Replace(target, old, new any, n int) (string, error) {
	// n == -1 => ReplaceALL

	sTarget, err := cast.ToStringE(target)
	if err != nil {
		return "", err
	}

	sOld, err := cast.ToStringE(old)
	if err != nil {
		return "", err
	}

	sNew, err := cast.ToStringE(new)
	if err != nil {
		return "", err
	}

	return strings.Replace(sTarget, sOld, sNew, n), nil
}
