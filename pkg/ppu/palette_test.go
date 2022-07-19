//go:build !nesgo
// +build !nesgo

package ppu

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

func TestPalette(t *testing.T) {
	t.Parallel()

	p := &palette{}
	p.write(0, 1)
	value := p.read(0)
	assert.Equal(t, 1, value)

	p.reset()
	value = p.read(0)
	assert.Equal(t, 0, value)

	p.write(0x21, 1)
	value = p.read(1)
	assert.Equal(t, 1, value)
}
