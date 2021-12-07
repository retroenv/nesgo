package neslib

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
	. "github.com/retroenv/nesgo/pkg/nes"
)

func TestParser(t *testing.T) {
	ClearRAM()
	// TODO add test
	assert.Equal(t, 0, *X)
}
