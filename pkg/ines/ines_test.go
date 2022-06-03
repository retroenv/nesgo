package ines

import (
	"bytes"
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

func testRom() []byte {
	b := []byte{iNESFileMagic[0], iNESFileMagic[1], iNESFileMagic[2], iNESFileMagic[3]}
	b = append(b, []byte{2, 1, 1, 0, 0}...)       // prg, chr, control 1, control 2, ram
	b = append(b, []byte{0, 0, 0, 0, 0, 0, 0}...) // reserved/padding

	prg := make([]byte, 2*16384)
	prg[0] = 0x80 // marker
	b = append(b, prg...)

	chr := make([]byte, 8192)
	chr[0] = 0x81 // marker
	b = append(b, chr...)

	return b
}

func TestLoadFile(t *testing.T) {
	rom := testRom()
	reader := bytes.NewReader(rom)

	cart, err := LoadFile(reader)
	assert.NoError(t, err)

	assert.Equal(t, 0x80, cart.PRG[0])
	assert.Equal(t, 0x81, cart.CHR[0])
	assert.Equal(t, 0, cart.Mapper)
	assert.Equal(t, 1, cart.Mirror)
	assert.Equal(t, 0, cart.Battery)
}

func TestCartridgeSave(t *testing.T) {
	c := &Cartridge{
		PRG:     make([]byte, 2*16384),
		CHR:     make([]byte, 8192),
		Mapper:  0,
		Mirror:  1,
		Battery: 0,
	}
	c.PRG[0] = 0x80 // marker
	c.CHR[0] = 0x81 // marker

	var buf bytes.Buffer
	assert.NoError(t, c.Save(&buf))

	rom := testRom()
	b := buf.Bytes()
	assert.Equal(t, rom, b)
}
