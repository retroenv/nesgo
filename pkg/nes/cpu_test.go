package nes

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

func TestAdc(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
	sys := newSystem()
	// TODO add test
	sys.And(0)
}

func TestAsl(t *testing.T) {
	t.Parallel()
	sys := newSystem()

	sys.A = 0b00000001
	sys.Asl()
	assert.Equal(t, 0b00000010, sys.A)
	assert.Equal(t, 0, sys.Flags.C)

	sys.A = 0b11111110
	sys.Asl()
	assert.Equal(t, 0b11111100, sys.A)
	assert.Equal(t, 1, sys.Flags.C)

	sys.writeMemory(1, 0b00000010)
	sys.Asl(Absolute(1))
	assert.Equal(t, 0b00000100, sys.readMemory(1))

	sys.writeMemory(4, 0b00000010)
	sys.X = 3
	sys.Asl(Absolute(1), sys.X)
	assert.Equal(t, 0b00000100, sys.readMemory(4))
}

func TestBcc(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	assert.Equal(t, true, sys.Bcc())
	sys.Flags.C = 1
	assert.Equal(t, false, sys.Bcc())
}

func TestBcs(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	assert.Equal(t, false, sys.Bcs())
	sys.Flags.C = 1
	assert.Equal(t, true, sys.Bcs())
}

func TestBeq(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	assert.Equal(t, false, sys.Beq())
	sys.Flags.Z = 1
	assert.Equal(t, true, sys.Beq())
}

func TestBit(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	// TODO add test
	sys.Bit(0)
}

func TestBmi(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	assert.Equal(t, false, sys.Bmi())
	sys.Flags.N = 1
	assert.Equal(t, true, sys.Bmi())
}

func TestBne(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	assert.Equal(t, true, sys.Bne())
	sys.Flags.Z = 1
	assert.Equal(t, false, sys.Bne())
}

func TestBpl(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	assert.Equal(t, true, sys.Bpl())
	sys.Flags.N = 1
	assert.Equal(t, false, sys.Bpl())
}

func TestBvc(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	assert.Equal(t, true, sys.Bvc())
	sys.Flags.V = 1
	assert.Equal(t, false, sys.Bvc())
}

func TestBvs(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	assert.Equal(t, false, sys.Bvs())
	sys.Flags.V = 1
	assert.Equal(t, true, sys.Bvs())
}

func TestClc(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Flags.C = 1
	sys.Clc()
	assert.Equal(t, 0, sys.Flags.C)
}

func TestCld(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Flags.D = 1
	sys.Cld()
	assert.Equal(t, 0, sys.Flags.D)
}

func TestCli(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Flags.I = 1
	sys.Cli()
	assert.Equal(t, 0, sys.Flags.I)
}

func TestClv(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Flags.V = 1
	sys.Clv()
	assert.Equal(t, 0, sys.Flags.V)
}

func TestCpx(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	// TODO add test
	sys.Cpx(0)
}

func TestCpy(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	// TODO add test
	sys.Cpy(0)
}

func TestDex(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.X = 2
	sys.Dex()
	assert.Equal(t, 1, sys.X)
}

func TestDey(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Y = 2
	sys.Dey()
	assert.Equal(t, 1, sys.Y)
}

func TestEor(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	// TODO add test
	sys.Eor(0)
}

func TestInx(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Inx()
	assert.Equal(t, 1, sys.X)
}

func TestIny(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Iny()
	assert.Equal(t, 1, sys.Y)
}

func TestLda(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Lda(1)
	assert.Equal(t, 1, sys.A)
}

func TestLdx(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Ldx(1)
	assert.Equal(t, 1, sys.X)
}

func TestLdy(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Ldy(1)
	assert.Equal(t, 1, sys.Y)
}

func TestLsr(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.A = 0b00000010
	sys.Lsr()
	assert.Equal(t, 0b00000001, sys.A)
	assert.Equal(t, 0, sys.Flags.C)

	sys.A = 0b01111111
	sys.Lsr()
	assert.Equal(t, 0b00111111, sys.A)
	assert.Equal(t, 1, sys.Flags.C)
}

func TestNop(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Nop()
}

func TestOra(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	// TODO add test
	sys.Ora(0)
}

func TestPha(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.A = 1
	sys.Pha()
	b := sys.readMemory(stackBase + initialStack)
	assert.Equal(t, sys.A, b)
	assert.Equal(t, stackBase+initialStack-1, sys.SP)
}

func TestPhp(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Php()
	b := sys.readMemory(stackBase + initialStack)
	// I + U are set by default, bit 4 and 5 are set from PHP
	assert.Equal(t, 0b111100, b)
}

func TestPla(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.SP = 1
	sys.writeMemory(stackBase+2, 1)
	sys.Pla()
	assert.Equal(t, 1, sys.A)
	assert.Equal(t, 2, sys.SP)
}

func TestPlp(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.SP = 1
	sys.writeMemory(stackBase+2, 1)
	sys.Plp()
	assert.Equal(t, 1, sys.flags())
	assert.Equal(t, 2, sys.SP)
}

func TestRol(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	// TODO add test
	sys.Rol()
}

func TestRor(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.A = 0b00000010
	sys.Ror()
	assert.Equal(t, 0b00000001, sys.A)
	assert.Equal(t, 0, sys.Flags.C)

	sys.A = 0b01111111
	sys.Ror()
	assert.Equal(t, 0b00111111, sys.A)
	assert.Equal(t, 1, sys.Flags.C)
	sys.Ror()
	assert.Equal(t, 0b10011111, sys.A)
	assert.Equal(t, 1, sys.Flags.C)
}

func TestRti(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Rti()
}

func TestSbc(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
	sys := newSystem()
	sys.Sec()
	assert.Equal(t, 1, sys.Flags.C)
}

func TestSed(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Sed()
	assert.Equal(t, 1, sys.Flags.D)
}

func TestSei(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Sei()
	assert.Equal(t, 1, sys.Flags.I)
}

func TestSta(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	sys := newSystem()
	sys.A = 2
	sys.Tax()
	assert.Equal(t, sys.A, sys.X)
}

func TestTay(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.A = 2
	sys.Tay()
	assert.Equal(t, sys.A, sys.Y)
}

func TestTsx(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Tsx()
	assert.Equal(t, initialStack, sys.SP)
}

func TestTxa(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.X = 2
	sys.Txa()
	assert.Equal(t, sys.X, sys.A)
}

func TestTxs(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.X = 2
	sys.Txs()
	assert.Equal(t, sys.X, sys.SP)
}

func TestTya(t *testing.T) {
	t.Parallel()
	sys := newSystem()
	sys.Y = 2
	sys.Tya()
	assert.Equal(t, sys.Y, sys.A)
}
