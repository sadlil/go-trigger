package trigger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifferentEventHandlers(t *testing.T) {
	one := New()
	two := New()

	one.On("one", func() {})
	assert.Equal(t, 1, one.EventCount())
	assert.Equal(t, 0, two.EventCount())
}
