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


func invoke(event string, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(functionMap[event])
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("Parameter Mismatched")
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result := f.Call(in)
	return result, nil
}


func clear(event string) error {
	if _, ok := functionMap[event]; !ok {
		return errors.New("Event Not Defined")
	}
	delete(functionMap, event)
	return nil
}

func deleteAll() error {
	functionMap = make(map[string]interface{})
	return nil
}

func eventList() []string {
	events := make([]string, 0)
	for k := range functionMap {
		events = append(events, k)
	}
	return events
}

func eventCount() int {
	return len(functionMap)
}

func hasEvent(event string) bool {
	_, ok := functionMap[event]
	return ok
}
