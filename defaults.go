package trigger

import "reflect"

type Trigger interface {
	On(event string, task interface{}) error
	Fire(event string, params ...interface{}) ([]reflect.Value, error)
	FireBackground(event string, params ...interface{}) (chan []reflect.Value, error)
	Clear(event string) error
	ClearEvents()
	HasEvent(event string) bool
	Events() []string
	EventCount() int
}

var defaultTrigger = New()

// Default global trigger options.
func On(event string, task interface{}) error {
	return defaultTrigger.On(event, task)
}

func Fire(event string, params ...interface{}) ([]reflect.Value, error) {
	return defaultTrigger.Fire(event, params...)
}

func FireBackground(event string, params ...interface{}) (chan []reflect.Value, error) {
	return defaultTrigger.FireBackground(event, params...)
}

func Clear(event string) error {
	return defaultTrigger.Clear(event)
}

func ClearEvents() {
	defaultTrigger.ClearEvents()
}

func HasEvent(event string) bool {
	return defaultTrigger.HasEvent(event)
}

func Events() []string {
	return defaultTrigger.Events()
}

func EventCount() int {
	return defaultTrigger.EventCount()
}
