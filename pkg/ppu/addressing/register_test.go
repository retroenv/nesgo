package addressing

import (
	"testing"

	"github.com/retroenv/retrogolib/assert"
)

func TestRegisterIncrementX(t *testing.T) {
	t.Parallel()

	r := &register{}

	assert.Equal(t, 0, r.CoarseX)
	r.incrementX()
	assert.Equal(t, 1, r.CoarseX)

	r.CoarseX = 31
	r.incrementX()
	assert.Equal(t, 0, r.CoarseX)
	assert.Equal(t, 1, r.NameTableX)

	r.CoarseX = 31
	r.incrementX()
	assert.Equal(t, 0, r.CoarseX)
	assert.Equal(t, 0, r.NameTableX)
}

func TestRegisterIncrementY(t *testing.T) {
	t.Parallel()

	r := &register{}

	assert.Equal(t, 0, r.FineY)
	r.incrementY()
	assert.Equal(t, 1, r.FineY)

	r.FineY = 7
	r.incrementY()
	assert.Equal(t, 0, r.FineY)
	assert.Equal(t, 1, r.CoarseY)
	assert.Equal(t, 0, r.NameTableY)

	r.FineY = 7
	r.CoarseY = 29
	r.incrementY()
	assert.Equal(t, 0, r.FineY)
	assert.Equal(t, 0, r.CoarseY)
	assert.Equal(t, 1, r.NameTableY)

	r.FineY = 7
	r.CoarseY = 31
	r.incrementY()
	assert.Equal(t, 0, r.FineY)
	assert.Equal(t, 0, r.CoarseY)
	assert.Equal(t, 1, r.NameTableY)
}
