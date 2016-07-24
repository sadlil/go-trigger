package trigger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDifferentEventHandlers(t *testing.T) {
	one := New()
	two := New()

	one.On("one", func() {})
	assert.Equal(t, 1, one.EventCount())
	assert.Equal(t, 0, two.EventCount())
}
