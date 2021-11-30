package nes

import (
	"os"
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

func TestMain(m *testing.M) {
	reset()
	os.Exit(m.Run())
}

func TestClcSec(t *testing.T) {
	Sec()
	assert.Equal(t, 1, C)
	Clc()
	assert.Equal(t, 0, C)
}

func TestCldSed(t *testing.T) {
	Sed()
	assert.Equal(t, 1, D)
	Cld()
	assert.Equal(t, 0, D)
}

func TestCliSei(t *testing.T) {
	Sei()
	assert.Equal(t, 1, I)
	Cli()
	assert.Equal(t, 0, I)
}

func TestClv(t *testing.T) {
	V = 1
	Clv()
	assert.Equal(t, 0, V)
}

func TestDex(t *testing.T) {
	assert.Equal(t, 0, cpu.X)
	Dex()
	assert.Equal(t, -1, cpu.X)
}

func TestDey(t *testing.T) {
	assert.Equal(t, 0, cpu.Y)
	Dey()
	assert.Equal(t, -1, cpu.Y)
}

func TestLdaSta(t *testing.T) {
	Lda(1)
	assert.Equal(t, 1, cpu.A)

	b := readMemory(0)
	assert.Equal(t, 0, b)

	Sta(0)
	b = readMemory(0)
	assert.Equal(t, 1, b)

	// TODO: fix
	// Dex()
	// Sta(0, X)
	// b = readMemory(0xff)
	// assert.Equal(t, 1, b)
	//
	// Dey()
	// Dey()
	// Lda(2)
	// Sta(0, Y)
	// b = readMemory(0xfe)
	// assert.Equal(t, 2, b)
}

// TODO: fix
// func TestStx(t *testing.T) {
// 	Dex() // x = -1
// 	b := readMemory(0)
// 	assert.Equal(t, 0, b)
//
// 	Stx(0) // *0 = -1
// 	b = readMemory(0)
// 	assert.Equal(t, -1, b)
//
// 	Dey()     // y = -1
// 	Stx(0, Y) // -1 = -1
// 	b = readMemory(0xff)
// 	assert.Equal(t, -1, b)
// }
