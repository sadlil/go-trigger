package trigger

import "reflect"

func On(event string, task interface{}) {
	add(event, task)
}

func Fire(event string, params ...interface{}) (result []reflect.Value, err error) {
	return invoke(event, params...)
}
