package nes

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

func TestController(t *testing.T) {
	c := newController()
	c.setButtonState(buttonB, true)

	assert.Equal(t, 0, c.read())
	assert.Equal(t, 1, c.read())
	assert.Equal(t, 0, c.read())

	c.setStrobeMode(1)
	assert.Equal(t, 0, c.read())
	assert.Equal(t, 0, c.read())

	c.setStrobeMode(0)
	assert.Equal(t, 0, c.read())
	assert.Equal(t, 1, c.read())

	c.setButtonState(buttonB, false)
	c.setStrobeMode(1)
	c.setStrobeMode(0)
	assert.Equal(t, 0, c.read())
	assert.Equal(t, 0, c.read())
}
