//go:build !nesgo
// +build !nesgo

// Package controller provides hardware controller functionality.
package controller

import "sync/atomic"

// Button defines a button on the controller.
type Button uint64

const (
	A      Button = 0b0000_0001
	B      Button = 0b0000_0010
	Select Button = 0b0000_0100
	Start  Button = 0b0000_1000
	Up     Button = 0b0001_0000
	Down   Button = 0b0010_0000
	Left   Button = 0b0100_0000
	Right  Button = 0b1000_0000
)

// Controller represents a hardware controller.
type Controller struct {
	// if strobeMode is set, it resets the pointer to the state to read
	// to the A button. The pointer is not advanced on every state read
	// until it strobe mode is set off again.
	strobeMode bool
	// button pressed state as flags. needs to be read atomically as it
	// written by the main goroutine that is locked for SDL/OpenGL usage
	// and the emulator running in a separate goroutine.
	buttons uint64
	// index (mask) of next button state to read
	index uint8
}

// New returns a new Controller.
func New() *Controller {
	c := &Controller{}
	c.reset()
	return c
}

func (c *Controller) reset() {
	c.buttons = 0
	c.strobeMode = false
	c.index = 1
}

// SetStrobeMode sets the strobe mode flag of the controller.
func (c *Controller) SetStrobeMode(mode uint8) {
	if mode&1 == 1 {
		c.strobeMode = true
		c.index = 1
	} else {
		c.strobeMode = false
	}
}

// Read returns the current button state.
func (c *Controller) Read() uint8 {
	state := atomic.LoadUint64(&c.buttons)
	if c.strobeMode {
		return uint8(state & uint64(A))
	}

	val := state & uint64(c.index) // nolint:ifshort
	c.index <<= 1
	if val != 0 {
		return 1
	}
	return 0
}

// SetButtonState sets the current button state.
func (c *Controller) SetButtonState(key Button, pressed bool) {
	state := atomic.LoadUint64(&c.buttons)
	if pressed {
		state |= uint64(key)
	} else {
		state &= ^uint64(key)
	}
	atomic.StoreUint64(&c.buttons, state)
}
