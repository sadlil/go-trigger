package trigger

import "reflect"

func On(event string, task interface{}) error {
	return add(event, task)
}

func Fire(event string, params ...interface{}) ([]reflect.Value, error) {
	return invoke(event, params...)
}

func FireBackground(event string, params ...interface{}) (chan []reflect.Value, error) {
	return invokeParallel(event, params...)
}

func Clear(event string) error {
	return clear(event)
}

func ClearEvents() error {
	return deleteAll()
}

func HasEvent(event string) bool {
	return hasEvent(event)
}

func Events() []string {
	return eventList()
}

func EventCount() int {
	return eventCount()
}
