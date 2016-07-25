package trigger

import (
	"runtime"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOn(t *testing.T) {
	err := On("test-event", func() {})
	assert.Equal(t, err, nil)
	assert.Equal(t, 1, EventCount())
	ClearEvents()
}

func TestDualOn(t *testing.T) {
	err := On("test-event", func() {})
	assert.Equal(t, err, nil)
	err2 := On("test-event", func() {})
	assert.NotEqual(t, err2, nil)
	assert.Equal(t, err2.Error(), "event already defined")
	assert.Equal(t, 1, EventCount())
	ClearEvents()
}

func TestTrigger(t *testing.T) {
	On("test-event", func() {})
	_, err := Fire("test-event")
	assert.Equal(t, err, nil)

	On("test-event2", func(a, b int) int {
		return a + b
	})
	vales, err := Fire("test-event2", 100, 5)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, vales, nil)
	assert.Equal(t, vales[0].Int(), int64(105))

	vales, err = Fire("test-event2", -100, 5)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, vales, nil)
	assert.Equal(t, vales[0].Int(), int64(-95))

	ClearEvents()
}

func TestClear(t *testing.T) {
	On("test-event", func() {})
	On("test-event2", func() {})
	assert.Equal(t, 2, EventCount())
	Clear("test-event")
	assert.Equal(t, 1, EventCount())
	err := On("test-event", func() {})
	assert.Equal(t, err, nil)
	assert.Equal(t, 2, EventCount())
	ClearEvents()
}

func TestClearEvents(t *testing.T) {
	On("test-event", func() {})
	On("test-event2", func() {})
	assert.Equal(t, 2, EventCount())
	ClearEvents()
	assert.Equal(t, 0, EventCount())
	On("test-event", func() {})
	assert.Equal(t, 1, EventCount())
	ClearEvents()
	assert.Equal(t, 0, EventCount())
}

func TestEventCount(t *testing.T) {
	On("test-event", func() {})
	assert.Equal(t, 1, EventCount())
	ClearEvents()
	assert.Equal(t, 0, EventCount())
	On("test-event1", func() {})
	On("test-event2", func() {})
	On("test-event3", func() {})
	On("test-event4", func() {})
	assert.Equal(t, 4, EventCount())
	On("test-event4", func() {})
	assert.Equal(t, 4, EventCount())
	ClearEvents()
}

func TestEvents(t *testing.T) {
	On("test-event1", func() {})
	On("test-event2", func() {})
	On("test-event3", func() {})
	On("test-event4", func() {})
	eventList := Events()
	assert.Equal(t, 4, len(eventList))
	sort.Strings(eventList)
	assert.Equal(t, []string{"test-event1", "test-event2", "test-event3", "test-event4"}, eventList)
	ClearEvents()
}

func TestHasEvent(t *testing.T) {
	On("test-event1", func() {})
	ret := HasEvent("test-event1")
	assert.Equal(t, ret, true)

	ret = HasEvent("test-event-not-found")
	assert.Equal(t, ret, false)
	ClearEvents()
}

func TestParallel(t *testing.T) {
	On("p-1", func() {
		for i := 1; i <= 10000; i++ {

		}
	})

	On("p-2", func() {
		for i := 1; i <= 10000; i++ {

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
	assert.Equal(t, 8, now-prev)
	ClearEvents()
}

func TestNotFunc(t *testing.T) {
	type test struct{}
	err := On("err-event", &test{})
	assert.NotNil(t, err)
	assert.Equal(t, "task is not a function", err.Error())
}
