package w32

import (
	"fmt"
	"reflect"
)

func run(obj any, name string, args ...any) ([]reflect.Value, error) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	values := reflect.ValueOf(obj).MethodByName(name).Call(inputs)
	if len(values) == 0 {
		return nil, fmt.Errorf("the method may not exists")
	}
	return values, nil
}
