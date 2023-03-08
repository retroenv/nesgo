//go:build !nesgo

package nes

import (
	"github.com/retroenv/nesgo/pkg/controller"
	"github.com/retroenv/retrogolib/input"
)

var controllerMapping = map[input.Key]controller.Button{
	input.Up:        controller.Up,
	input.Down:      controller.Down,
	input.Left:      controller.Left,
	input.Right:     controller.Right,
	input.Z:         controller.A,
	input.X:         controller.B,
	input.Enter:     controller.Start,
	input.Backspace: controller.Select,
}

// KeyDown gets called when a key down event is registered.
func (sys *System) KeyDown(key input.Key) {
	controllerKey, ok := controllerMapping[key]
	if !ok {
		return
	}
	sys.Bus.Controller1.SetButtonState(controllerKey, true)
}

// KeyUp gets called when a key up event is registered.
func (sys *System) KeyUp(key input.Key) {
	controllerKey, ok := controllerMapping[key]
	if !ok {
		return
	}
	sys.Bus.Controller1.SetButtonState(controllerKey, false)
}
