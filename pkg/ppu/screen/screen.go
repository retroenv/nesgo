// Package screen handles PPU screen drawing.
package screen

import (
	"image"
	"image/color"
)

const (
	Width  = 256
	Height = 240
)

// Screen implements PPU screen drawing support.
type Screen struct {
	back  *image.RGBA // rendering in progress image
	front *image.RGBA // currently visible image
}

// New returns a new screen manager.
func New() *Screen {
	return &Screen{
		back:  image.NewRGBA(image.Rect(0, 0, Width, Height)),
		front: image.NewRGBA(image.Rect(0, 0, Width, Height)),
	}
}

// SetPixel sets a pixel in the rendering image.
func (s *Screen) SetPixel(x, y int, color color.RGBA) {
	s.back.SetRGBA(x, y, color)
}

// Image returns the rendered image to display.
func (s *Screen) Image() *image.RGBA {
	return s.front
}

// FinishRendering finishes rendering by switching the visible image with the rendered one.
func (s *Screen) FinishRendering() {
	s.front, s.back = s.back, s.front
}
