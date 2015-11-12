package trigger

import (
	"testing"
	. "gopkg.in/go-playground/assert.v1"
	"sort"
	"fmt"
"runtime"
)

func TestOn(t *testing.T) {
	err := On("test-event", func() {	})
	Equal(t, err, nil)
	Equal(t, 1, EventCount())
	ClearEvents()
}

func TestDualOn(t *testing.T) {
	err := On("test-event", func() {	})
	Equal(t, err, nil)
	err2 := On("test-event", func() {	})
	NotEqual(t, err2, nil)
	Equal(t, err2.Error(), "Event Already Defined")
	Equal(t, 1, EventCount())
	ClearEvents()
}

func TestTrigger(t *testing.T) {
	On("test-event", func() {
		fmt.Println("Testing Triggered Ok.")
	})
	_, err := Fire("test-event")
	Equal(t, err, nil)


	On("test-event2", func(a, b int) int {
		return a + b
	})
	vales, err := Fire("test-event2", 100, 5)
	Equal(t, err, nil)
	NotEqual(t, vales, nil)
	Equal(t, vales[0].Int(), int64(105))


	vales, err = Fire("test-event2", -100, 5)
	Equal(t, err, nil)
	NotEqual(t, vales, nil)
	Equal(t, vales[0].Int(), int64(-95))

	ClearEvents()
}

func TestClear(t *testing.T) {
	On("test-event", func() {	})
	On("test-event2", func() {	})
	Equal(t, 2, EventCount())
	Clear("test-event")
	Equal(t, 1, EventCount())
	err := On("test-event", func() {	})
	Equal(t, err, nil)
	Equal(t, 2, EventCount())
	ClearEvents()
}

func TestClearEvents(t *testing.T) {
	On("test-event", func() {	})
	On("test-event2", func() {	})
	Equal(t, 2, EventCount())
	ClearEvents()
	Equal(t, 0, EventCount())
	On("test-event", func() {	})
	Equal(t, 1, EventCount())
	ClearEvents()
	Equal(t, 0, EventCount())
}

func TestEventCount(t *testing.T) {
	On("test-event", func() {	})
	Equal(t, 1, EventCount())
	ClearEvents()
	Equal(t, 0, EventCount())
	On("test-event1", func() {	})
	On("test-event2", func() {	})
	On("test-event3", func() {	})
	On("test-event4", func() {	})
	Equal(t, 4, EventCount())
	On("test-event4", func() {	})
	Equal(t, 4, EventCount())
	ClearEvents()
}

func TestEvents(t *testing.T) {
	On("test-event1", func() {	})
	On("test-event2", func() {	})
	On("test-event3", func() {	})
	On("test-event4", func() {	})
	eventList := Events()
	Equal(t, 4, len(eventList))
	sort.Strings(eventList)
	Equal(t, []string{"test-event1", "test-event2", "test-event3", "test-event4"}, eventList)
	ClearEvents()
}

func TestHasEvent(t *testing.T) {
	On("test-event1", func() {	})
	ret := HasEvent("test-event1")
	Equal(t, ret, true)

	ret = HasEvent("test-event-not-found")
	Equal(t, ret, false)
	ClearEvents()
}


func TestParallel(t *testing.T) {
	On("p-1", func () {
		for i:=1; i<=10000; i++ {

		}
	})

	On("p-2", func () {
		for i:=1; i<=10000; i++ {

		}
	})
	prev := runtime.NumGoroutine()
	FireBackground("p-1")
	FireBackground("p-2")
	FireBackground("p-2")
	FireBackground("p-2")
	FireBackground("p-2")
	FireBackground("p-2")
	FireBackground("p-2")
	FireBackground("p-2")

	now := runtime.NumGoroutine()
	fmt.Println("Number of go routine running ", now - prev)
	Equal(t, 8, now - prev)
	ClearEvents();
}
