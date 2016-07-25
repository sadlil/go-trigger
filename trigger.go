package trigger

import (
	"errors"
	"reflect"
	"sync"
)

func New() Trigger {
	return &trigger{
		functionMap: make(map[string]interface{}),
	}
}

type trigger struct {
	functionMap map[string]interface{}

	mu sync.Mutex
}

func (t *trigger) On(event string, task interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if _, ok := t.functionMap[event]; ok {
		return errors.New("event already defined")
	}
	if reflect.ValueOf(task).Type().Kind() != reflect.Func {
		return errors.New("task is not a function")
	}
	t.functionMap[event] = task
	return nil
}

func (t *trigger) Fire(event string, params ...interface{}) ([]reflect.Value, error) {
	f, in, err := t.read(event, params...)
	if err != nil {
		return nil, err
	}
	result := f.Call(in)
	return result, nil
}

func (t *trigger) FireBackground(event string, params ...interface{}) (chan []reflect.Value, error) {
	f, in, err := t.read(event, params...)
	if err != nil {
		return nil, err
	}
	results := make(chan []reflect.Value)
	go func() {
		results <- f.Call(in)
	}()
	return results, nil
}

func (t *trigger) Clear(event string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if _, ok := t.functionMap[event]; !ok {
		return errors.New("event not defined")
	}
	delete(t.functionMap, event)
	return nil
}

func (t *trigger) ClearEvents() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.functionMap = make(map[string]interface{})
}

func (t *trigger) HasEvent(event string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	_, ok := t.functionMap[event]
	return ok
}

func (t *trigger) Events() []string {
	t.mu.Lock()
	defer t.mu.Unlock()
	events := make([]string, 0)
	for k := range t.functionMap {
		events = append(events, k)
	}
	return events
}

func (t *trigger) EventCount() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return len(t.functionMap)
}

func (t *trigger) read(event string, params ...interface{}) (reflect.Value, []reflect.Value, error) {
	t.mu.Lock()
	task, ok := t.functionMap[event]
	t.mu.Unlock()
	if !ok {
		return reflect.Value{}, nil, errors.New("no task found for event")
	}
	f := reflect.ValueOf(task)
	if len(params) != f.Type().NumIn() {
		return reflect.Value{}, nil, errors.New("parameter mismatched")
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return f, in, nil
}
