package controller

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

func TestController(t *testing.T) {
	c := New()
	c.SetButtonState(ButtonB, true)

	assert.Equal(t, 0, c.Read())
	assert.Equal(t, 1, c.Read())
	assert.Equal(t, 0, c.Read())

	c.SetStrobeMode(1)
	assert.Equal(t, 0, c.Read())
	assert.Equal(t, 0, c.Read())

	c.SetStrobeMode(0)
	assert.Equal(t, 0, c.Read())
	assert.Equal(t, 1, c.Read())

	c.SetButtonState(ButtonB, false)
	c.SetStrobeMode(1)
	c.SetStrobeMode(0)
	assert.Equal(t, 0, c.Read())
	assert.Equal(t, 0, c.Read())
}