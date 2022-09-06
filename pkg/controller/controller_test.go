package controller

import (
	"testing"

	"github.com/retroenv/retrogolib/assert"
)

func TestController(t *testing.T) {
	c := New()
	c.SetButtonState(B, true)

	assert.Equal(t, 0, c.Read())
	assert.Equal(t, 1, c.Read())
	assert.Equal(t, 0, c.Read())

	c.SetStrobeMode(1)
	assert.Equal(t, 0, c.Read())
	assert.Equal(t, 0, c.Read())

	c.SetStrobeMode(0)
	assert.Equal(t, 0, c.Read())
	assert.Equal(t, 1, c.Read())

	c.SetButtonState(B, false)
	c.SetStrobeMode(1)
	c.SetStrobeMode(0)
	assert.Equal(t, 0, c.Read())
	assert.Equal(t, 0, c.Read())
}
