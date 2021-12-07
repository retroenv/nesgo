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
	assert.Equal(t, 1, cpu.Flags.C)

	Adc(2)
	assert.Equal(t, 4, cpu.A)
	assert.Equal(t, 0, cpu.Flags.C)
}

func TestAnd(t *testing.T) {
	reset()
	// TODO add test
	And(0)
}

func TestAsl(t *testing.T) {
	reset()
	// TODO add test
	Asl()
}

func TestBcc(t *testing.T) {
	reset()
	assert.Equal(t, true, Bcc())
	cpu.Flags.C = 1
	assert.Equal(t, false, Bcc())
}

func TestBcs(t *testing.T) {
	reset()
	assert.Equal(t, false, Bcs())
	cpu.Flags.C = 1
	assert.Equal(t, true, Bcs())
}

func TestBeq(t *testing.T) {
	reset()
	assert.Equal(t, false, Beq())
	cpu.Flags.Z = 1
	assert.Equal(t, true, Beq())
}

func TestBit(t *testing.T) {
	reset()
	// TODO add test
	Bit(0)
}

func TestBmi(t *testing.T) {
	reset()
	assert.Equal(t, false, Bmi())
	cpu.Flags.N = 1
	assert.Equal(t, true, Bmi())
}

func TestBne(t *testing.T) {
	reset()
	assert.Equal(t, true, Bne())
	cpu.Flags.Z = 1
	assert.Equal(t, false, Bne())
}

func TestBpl(t *testing.T) {
	reset()
	assert.Equal(t, true, Bpl())
	cpu.Flags.N = 1
	assert.Equal(t, false, Bpl())
}

func TestBvc(t *testing.T) {
	reset()
	assert.Equal(t, true, Bvc())
	cpu.Flags.V = 1
	assert.Equal(t, false, Bvc())
}

func TestBvs(t *testing.T) {
	reset()
	assert.Equal(t, false, Bvs())
	cpu.Flags.V = 1
	assert.Equal(t, true, Bvs())
}

func TestClc(t *testing.T) {
	reset()
	cpu.Flags.C = 1
	Clc()
	assert.Equal(t, 0, cpu.Flags.C)
}

func TestCld(t *testing.T) {
	reset()
	cpu.Flags.D = 1
	Cld()
	assert.Equal(t, 0, cpu.Flags.D)
}

func TestCli(t *testing.T) {
	reset()
	cpu.Flags.I = 1
	Cli()
	assert.Equal(t, 0, cpu.Flags.I)
}

func TestClv(t *testing.T) {
	reset()
	cpu.Flags.V = 1
	Clv()
	assert.Equal(t, 0, cpu.Flags.V)
}

func TestCpx(t *testing.T) {
	reset()
	// TODO add test
	Cpx(0)
}

func TestCpy(t *testing.T) {
	reset()
	// TODO add test
	Cpy(0)
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

func TestEor(t *testing.T) {
	reset()
	// TODO add test
	Eor(0)
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

func TestLsr(t *testing.T) {
	reset()
	// TODO add test
	Lsr()
}

func TestNop(t *testing.T) {
	reset()
	Nop()
}

func TestOra(t *testing.T) {
	reset()
	// TODO add test
	Ora(0)
}

func TestPha(t *testing.T) {
	reset()
	// TODO add test
	Pha()
}

func TestPhp(t *testing.T) {
	reset()
	// TODO add test
	Php()
}

func TestPla(t *testing.T) {
	reset()
	// TODO add test
	Lsr()
}

func TestPlp(t *testing.T) {
	reset()
	// TODO add test
	Plp()
}

func TestRol(t *testing.T) {
	reset()
	// TODO add test
	Rol()
}

func TestRor(t *testing.T) {
	reset()
	// TODO add test
	Ror()
}

func TestRti(t *testing.T) {
	reset()
	Rti()
}

func TestSbc(t *testing.T) {
	reset()
	cpu.A = 2
	Sbc(0xff)
	assert.Equal(t, 2, cpu.A)
	assert.Equal(t, 0, cpu.Flags.C)

	Sbc(2)
	assert.Equal(t, 0xff, cpu.A)
	assert.Equal(t, 0, cpu.Flags.C)
}

func TestSec(t *testing.T) {
	reset()
	Sec()
	assert.Equal(t, 1, cpu.Flags.C)
}

func TestSed(t *testing.T) {
	reset()
	Sed()
	assert.Equal(t, 1, cpu.Flags.D)
}

func TestSei(t *testing.T) {
	reset()
	Sei()
	assert.Equal(t, 1, cpu.Flags.I)
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
	assert.Equal(t, initialStack, cpu.SP)
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
	assert.Equal(t, cpu.X, cpu.SP)
}

func TestTya(t *testing.T) {
	reset()
	cpu.Y = 2
	Tya()
	assert.Equal(t, cpu.Y, cpu.A)
}
