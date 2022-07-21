//go:build !nesgo
// +build !nesgo

// Package renderstate handles PPU render handling of cycles, scan lines and frames.
package renderstate

type mask interface {
	RenderBackground() bool
	RenderSprites() bool
}

// RenderState implements a PPU render state manager.
type RenderState struct {
	cycle    int // 0-340, 0=idle,1-336=tile data fetching,337-340=nameTable fetching
	scanLine int // 0-261, 0-239=visible, 240=post-render, 241-260=vertical blank, 261=pre-render
	frame    uint64
}

// New returns a new render state manager.
func New() *RenderState {
	return &RenderState{
		cycle:    340,
		scanLine: 240,
	}
}

// Tick updates cycle, scanLine and frame counters.
func (r *RenderState) Tick(mask mask) {
	if mask.RenderBackground() || mask.RenderSprites() {
		// for odd frames, the cycle at the end of the scanline is skipped
		if r.scanLine == 261 && r.cycle == 339 && r.frame%2 == 1 {
			r.nextFrame()
			return
		}
	}

	r.cycle++
	if r.cycle <= 340 {
		return
	}
	r.cycle = 0

	r.scanLine++
	if r.scanLine <= 261 {
		return
	}
	r.nextFrame()
}

func (r *RenderState) nextFrame() {
	r.cycle = 0
	r.scanLine = 0
	r.frame++
}

// Cycle returns the current cycle, possible values are 0-340.
func (r RenderState) Cycle() int {
	return r.cycle
}

// ScanLine returns the current scanline, possible values are 0-261.
func (r RenderState) ScanLine() int {
	return r.scanLine
}
