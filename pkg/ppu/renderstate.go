//go:build !nesgo
// +build !nesgo

package ppu

type renderState struct {
	cycle    int // 0-340, 0=idle,1-336=tile data fetching,337-340=nameTable fetching
	scanLine int // 0-261, 0-239=visible, 240=post-render, 241-260=vertical blank, 261=pre-render
	frame    uint64
}

func newRenderState() *renderState {
	return &renderState{
		cycle:    340,
		scanLine: 240,
	}
}

// tick updates cycle, scanLine and frame counters.
func (p *renderState) tick(mask mask) {
	if mask.RenderBackground || mask.RenderSprites {
		// for odd frames, the cycle at the end of the scanline is skipped
		if p.scanLine == 261 && p.cycle == 339 && p.frame%2 == 1 {
			p.nextFrame()
			return
		}
	}

	p.cycle++
	if p.cycle <= 340 {
		return
	}
	p.cycle = 0

	p.scanLine++
	if p.scanLine <= 261 {
		return
	}
	p.nextFrame()
}

func (p *renderState) nextFrame() {
	p.cycle = 0
	p.scanLine = 0
	p.frame++
}
