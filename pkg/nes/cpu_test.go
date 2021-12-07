package nes

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

func TestAdc(t *testing.T) {
	sys := newSystem()
	sys.A = 2
	sys.Adc(0xff)
	assert.Equal(t, 1, sys.A)
	assert.Equal(t, 1, sys.Flags.C)

	sys.Adc(2)
	assert.Equal(t, 4, sys.A)
	assert.Equal(t, 0, sys.Flags.C)
}

func TestAnd(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.And(0)
}

func TestAsl(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Asl()
}

func TestBcc(t *testing.T) {
	sys := newSystem()
	assert.Equal(t, true, sys.Bcc())
	sys.Flags.C = 1
	assert.Equal(t, false, sys.Bcc())
}

func TestBcs(t *testing.T) {
	sys := newSystem()
	assert.Equal(t, false, sys.Bcs())
	sys.Flags.C = 1
	assert.Equal(t, true, sys.Bcs())
}

func TestBeq(t *testing.T) {
	sys := newSystem()
	assert.Equal(t, false, sys.Beq())
	sys.Flags.Z = 1
	assert.Equal(t, true, sys.Beq())
}

func TestBit(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Bit(0)
}

func TestBmi(t *testing.T) {
	sys := newSystem()
	assert.Equal(t, false, sys.Bmi())
	sys.Flags.N = 1
	assert.Equal(t, true, sys.Bmi())
}

func TestBne(t *testing.T) {
	sys := newSystem()
	assert.Equal(t, true, sys.Bne())
	sys.Flags.Z = 1
	assert.Equal(t, false, sys.Bne())
}

func TestBpl(t *testing.T) {
	sys := newSystem()
	assert.Equal(t, true, sys.Bpl())
	sys.Flags.N = 1
	assert.Equal(t, false, sys.Bpl())
}

func TestBvc(t *testing.T) {
	sys := newSystem()
	assert.Equal(t, true, sys.Bvc())
	sys.Flags.V = 1
	assert.Equal(t, false, sys.Bvc())
}

func TestBvs(t *testing.T) {
	sys := newSystem()
	assert.Equal(t, false, sys.Bvs())
	sys.Flags.V = 1
	assert.Equal(t, true, sys.Bvs())
}

func TestClc(t *testing.T) {
	sys := newSystem()
	sys.Flags.C = 1
	sys.Clc()
	assert.Equal(t, 0, sys.Flags.C)
}

func TestCld(t *testing.T) {
	sys := newSystem()
	sys.Flags.D = 1
	sys.Cld()
	assert.Equal(t, 0, sys.Flags.D)
}

func TestCli(t *testing.T) {
	sys := newSystem()
	sys.Flags.I = 1
	sys.Cli()
	assert.Equal(t, 0, sys.Flags.I)
}

func TestClv(t *testing.T) {
	sys := newSystem()
	sys.Flags.V = 1
	sys.Clv()
	assert.Equal(t, 0, sys.Flags.V)
}

func TestCpx(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Cpx(0)
}

func TestCpy(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Cpy(0)
}

func TestDex(t *testing.T) {
	sys := newSystem()
	sys.X = 2
	sys.Dex()
	assert.Equal(t, 1, sys.X)
}

func TestDey(t *testing.T) {
	sys := newSystem()
	sys.Y = 2
	sys.Dey()
	assert.Equal(t, 1, sys.Y)
}

func TestEor(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Eor(0)
}

func TestInx(t *testing.T) {
	sys := newSystem()
	sys.Inx()
	assert.Equal(t, 1, sys.X)
}

func TestIny(t *testing.T) {
	sys := newSystem()
	sys.Iny()
	assert.Equal(t, 1, sys.Y)
}

func TestLda(t *testing.T) {
	sys := newSystem()
	sys.Lda(1)
	assert.Equal(t, 1, sys.A)
}

func TestLdx(t *testing.T) {
	sys := newSystem()
	sys.Ldx(1)
	assert.Equal(t, 1, sys.X)
}

func TestLdy(t *testing.T) {
	sys := newSystem()
	sys.Ldy(1)
	assert.Equal(t, 1, sys.Y)
}

func TestLsr(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Lsr()
}

func TestNop(t *testing.T) {
	sys := newSystem()
	sys.Nop()
}

func TestOra(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Ora(0)
}

func TestPha(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Pha()
}

func TestPhp(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Php()
}

func TestPla(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Lsr()
}

func TestPlp(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Plp()
}

func TestRol(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Rol()
}

func TestRor(t *testing.T) {
	sys := newSystem()
	// TODO add test
	sys.Ror()
}

func TestRti(t *testing.T) {
	sys := newSystem()
	sys.Rti()
}

func TestSbc(t *testing.T) {
	sys := newSystem()
	sys.A = 2
	sys.Sbc(0xff)
	assert.Equal(t, 2, sys.A)
	assert.Equal(t, 0, sys.Flags.C)

	sys.Sbc(2)
	assert.Equal(t, 0xff, sys.A)
	assert.Equal(t, 0, sys.Flags.C)
}

func TestSec(t *testing.T) {
	sys := newSystem()
	sys.Sec()
	assert.Equal(t, 1, sys.Flags.C)
}

func TestSed(t *testing.T) {
	sys := newSystem()
	sys.Sed()
	assert.Equal(t, 1, sys.Flags.D)
}

func TestSei(t *testing.T) {
	sys := newSystem()
	sys.Sei()
	assert.Equal(t, 1, sys.Flags.I)
}

func TestSta(t *testing.T) {
	sys := newSystem()
	sys.A = 11
	sys.Sta(0)
	b := sys.readMemory(0)
	assert.Equal(t, sys.A, b)

	sys.X = 0x22
	sys.Sta(Absolute(0), sys.X)
	b = sys.readMemory(0x22)
	assert.Equal(t, sys.A, b)
}

func TestStx(t *testing.T) {
	sys := newSystem()
	sys.X = 11
	sys.Stx(0)
	b := sys.readMemory(0)
	assert.Equal(t, sys.X, b)

	sys.Y = 0x22
	sys.Stx(Absolute(0), sys.Y)
	b = sys.readMemory(0x22)
	assert.Equal(t, sys.X, b)
}

func TestSty(t *testing.T) {
	sys := newSystem()
	sys.Y = 11
	sys.Sty(0)
	b := sys.readMemory(0)
	assert.Equal(t, sys.Y, b)

	sys.X = 0x22
	sys.Sty(Absolute(0), sys.X)
	b = sys.readMemory(0x22)
	assert.Equal(t, sys.Y, b)
}

func TestTax(t *testing.T) {
	sys := newSystem()
	sys.A = 2
	sys.Tax()
	assert.Equal(t, sys.A, sys.X)
}

func TestTay(t *testing.T) {
	sys := newSystem()
	sys.A = 2
	sys.Tay()
	assert.Equal(t, sys.A, sys.Y)
}

func TestTsx(t *testing.T) {
	sys := newSystem()
	sys.Tsx()
	assert.Equal(t, initialStack, sys.SP)
}

func TestTxa(t *testing.T) {
	sys := newSystem()
	sys.X = 2
	sys.Txa()
	assert.Equal(t, sys.X, sys.A)
}

func TestTxs(t *testing.T) {
	sys := newSystem()
	sys.X = 2
	sys.Txs()
	assert.Equal(t, sys.X, sys.SP)
}

func TestTya(t *testing.T) {
	sys := newSystem()
	sys.Y = 2
	sys.Tya()
	assert.Equal(t, sys.Y, sys.A)
}
