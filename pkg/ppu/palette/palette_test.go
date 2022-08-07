//go:build !nesgo

package palette

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

func TestPalette(t *testing.T) {
	t.Parallel()

	p := &Palette{}
	p.Write(0, 1)
	value := p.Read(0)
	assert.Equal(t, 1, value)

	p.Write(0x21, 1)
	value = p.Read(1)
	assert.Equal(t, 1, value)
}
