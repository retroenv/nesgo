//go:build !nesgo
// +build !nesgo

package nes

import "sync/atomic"

type button uint64

// nolint: deadcode,varcheck
const (
	buttonRight  button = 0b10000000
	buttonLeft   button = 0b01000000
	buttonDown   button = 0b00100000
	buttonUp     button = 0b00010000
	buttonStart  button = 0b00001000
	buttonSelect button = 0b00000100
	buttonB      button = 0b00000010
	buttonA      button = 0b00000001
)

type controller struct {
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

func (c *controller) reset() {
	c.buttons = 0
	c.strobeMode = false
	c.index = 1
}

func (c *controller) setStrobeMode(mode uint8) {
	if mode&1 == 1 {
		c.strobeMode = true
		c.index = 1
	} else {
		c.strobeMode = false
	}
}

func (c *controller) read() uint8 {
	state := atomic.LoadUint64(&c.buttons)
	if c.strobeMode {
		return uint8(state & uint64(buttonA))
	}

	val := state & uint64(c.index) // nolint:ifshort
	c.index <<= 1
	if val != 0 {
		return 1
	}
	return 0
}

// nolint: unused
func (c *controller) setButtonState(key button, pressed bool) {
	state := atomic.LoadUint64(&c.buttons)
	if pressed {
		state |= uint64(key)
	} else {
		state &= ^uint64(key)
	}
	atomic.StoreUint64(&c.buttons, state)
}
