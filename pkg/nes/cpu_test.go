package nes

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

func TestAdc(t *testing.T) {
	reset()
	cpu.A = 2
	Adc(0xff)
	assert.Equal(t, 1, cpu.A)
	assert.Equal(t, 1, C)

	Adc(2)
	assert.Equal(t, 4, cpu.A)
	assert.Equal(t, 0, C)
}

func TestClc(t *testing.T) {
	reset()
	C = 1
	Clc()
	assert.Equal(t, 0, C)
}

func TestCld(t *testing.T) {
	reset()
	D = 1
	Cld()
	assert.Equal(t, 0, D)
}

func TestCli(t *testing.T) {
	reset()
	I = 1
	Cli()
	assert.Equal(t, 0, I)
}

func TestClv(t *testing.T) {
	reset()
	V = 1
	Clv()
	assert.Equal(t, 0, V)
}

func TestDex(t *testing.T) {
	reset()
	cpu.X = 2
	Dex()
	assert.Equal(t, 1, cpu.X)
}

func TestDey(t *testing.T) {
	reset()
	cpu.Y = 2
	Dey()
	assert.Equal(t, 1, cpu.Y)
}

func TestInx(t *testing.T) {
	reset()
	Inx()
	assert.Equal(t, 1, cpu.X)
}

func TestIny(t *testing.T) {
	reset()
	Iny()
	assert.Equal(t, 1, cpu.Y)
}

func TestLda(t *testing.T) {
	reset()
	Lda(1)
	assert.Equal(t, 1, cpu.A)
}

func TestLdx(t *testing.T) {
	reset()
	Ldx(1)
	assert.Equal(t, 1, cpu.X)
}

func TestLdy(t *testing.T) {
	reset()
	Ldy(1)
	assert.Equal(t, 1, cpu.Y)
}

func TestNop(t *testing.T) {
	reset()
	Nop()
}

func TestRti(t *testing.T) {
	reset()
	Rti()
}

func TestSec(t *testing.T) {
	reset()
	Sec()
	assert.Equal(t, 1, C)
}

func TestSed(t *testing.T) {
	reset()
	Sed()
	assert.Equal(t, 1, D)
}

func TestSei(t *testing.T) {
	reset()
	Sei()
	assert.Equal(t, 1, I)
}

func TestSta(t *testing.T) {
	reset()
	cpu.A = 11
	Sta(0)
	b := readMemory(0)
	assert.Equal(t, cpu.A, b)

	cpu.X = 0x22
	Sta(Absolute(0), X)
	b = readMemory(0x22)
	assert.Equal(t, cpu.A, b)
}

func TestStx(t *testing.T) {
	reset()
	cpu.X = 11
	Stx(0)
	b := readMemory(0)
	assert.Equal(t, cpu.X, b)

	cpu.Y = 0x22
	Stx(Absolute(0), Y)
	b = readMemory(0x22)
	assert.Equal(t, cpu.X, b)
}

func TestSty(t *testing.T) {
	reset()
	cpu.Y = 11
	Sty(0)
	b := readMemory(0)
	assert.Equal(t, cpu.Y, b)

	cpu.X = 0x22
	Sty(Absolute(0), X)
	b = readMemory(0x22)
	assert.Equal(t, cpu.Y, b)
}

func TestTax(t *testing.T) {
	reset()
	cpu.A = 2
	Tax()
	assert.Equal(t, cpu.A, cpu.X)
}

func TestTay(t *testing.T) {
	reset()
	cpu.A = 2
	Tay()
	assert.Equal(t, cpu.A, cpu.Y)
}

func TestTsx(t *testing.T) {
	reset()
	Tsx()
	assert.Equal(t, initialStack, SP)
}

func TestTxa(t *testing.T) {
	reset()
	cpu.X = 2
	Txa()
	assert.Equal(t, cpu.X, cpu.A)
}

func TestTxs(t *testing.T) {
	reset()
	cpu.X = 2
	Txs()
	assert.Equal(t, cpu.X, SP)
}

func TestTya(t *testing.T) {
	reset()
	cpu.Y = 2
	Tya()
	assert.Equal(t, cpu.Y, cpu.A)
}
