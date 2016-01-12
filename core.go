package trigger

import (
	"reflect"
	"errors"
	"runtime"
)

var functionMap map[string][]interface{}

func init() {
	functionMap = make(map[string][]interface{})
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func add(event string, task interface{}) {
	functionMap[event] = append(functionMap[event], task);
}

func invoke(event string, params ...interface{}) ([][]reflect.Value, error) {
	result := make([][]reflect.Value, 0)
	for _, regif := range(functionMap[event]) {
		f := reflect.ValueOf(regif)
		if len(params) != f.Type().NumIn() {
			return nil, errors.New("Parameter Mismatched")
		}
		in := make([]reflect.Value, len(params))
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
		}
		res := f.Call(in)
		result = append(result, res)
	}
	return result, nil
}

func invokeParallel(event string, params ...interface{}) ([]chan []reflect.Value, error) {
	result := make([]chan []reflect.Value, 0)
	for _, regif := range(functionMap[event]) {
		f := reflect.ValueOf(regif)
		if len(params) != f.Type().NumIn() {
			return nil, errors.New("Parameter Mismatched")
		}
		in := make([]reflect.Value, len(params))
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
		}
		results := make(chan []reflect.Value)
		result = append(result, results)
		go func() {
			results <- f.Call(in)
		}()

	}
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
	functionMap = make(map[string][]interface{})
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
