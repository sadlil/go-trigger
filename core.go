package trigger

import (
	"reflect"
	"errors"
)

var functionMap map[string]interface{}

func init() {
	functionMap = make(map[string]interface{})
}

func add(event string, task interface{}) error {
	if _, ok := functionMap[event]; ok {
		return errors.New("Event Already Defined")
	}
	functionMap[event] = task;
	return nil
}


func invoke(event string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(functionMap[event])
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("Parameter Mismatched")
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return result, nil
}
