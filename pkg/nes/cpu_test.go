package nes

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

func TestAdc(t *testing.T) {
	system := reset()
	system.A = 2
	system.Adc(0xff)
	assert.Equal(t, 1, system.A)
	assert.Equal(t, 1, system.Flags.C)

	system.Adc(2)
	assert.Equal(t, 4, system.A)
	assert.Equal(t, 0, system.Flags.C)
}

func TestAnd(t *testing.T) {
	system := reset()
	// TODO add test
	system.And(0)
}

func TestAsl(t *testing.T) {
	system := reset()
	// TODO add test
	system.Asl()
}

func TestBcc(t *testing.T) {
	system := reset()
	assert.Equal(t, true, system.Bcc())
	system.Flags.C = 1
	assert.Equal(t, false, system.Bcc())
}

func TestBcs(t *testing.T) {
	system := reset()
	assert.Equal(t, false, system.Bcs())
	system.Flags.C = 1
	assert.Equal(t, true, system.Bcs())
}

func TestBeq(t *testing.T) {
	system := reset()
	assert.Equal(t, false, system.Beq())
	system.Flags.Z = 1
	assert.Equal(t, true, system.Beq())
}

func TestBit(t *testing.T) {
	system := reset()
	// TODO add test
	system.Bit(0)
}

func TestBmi(t *testing.T) {
	system := reset()
	assert.Equal(t, false, system.Bmi())
	system.Flags.N = 1
	assert.Equal(t, true, system.Bmi())
}

func TestBne(t *testing.T) {
	system := reset()
	assert.Equal(t, true, system.Bne())
	system.Flags.Z = 1
	assert.Equal(t, false, system.Bne())
}

func TestBpl(t *testing.T) {
	system := reset()
	assert.Equal(t, true, system.Bpl())
	system.Flags.N = 1
	assert.Equal(t, false, system.Bpl())
}

func TestBvc(t *testing.T) {
	system := reset()
	assert.Equal(t, true, system.Bvc())
	system.Flags.V = 1
	assert.Equal(t, false, system.Bvc())
}

func TestBvs(t *testing.T) {
	system := reset()
	assert.Equal(t, false, system.Bvs())
	system.Flags.V = 1
	assert.Equal(t, true, system.Bvs())
}

func TestClc(t *testing.T) {
	system := reset()
	system.Flags.C = 1
	system.Clc()
	assert.Equal(t, 0, system.Flags.C)
}

func TestCld(t *testing.T) {
	system := reset()
	system.Flags.D = 1
	system.Cld()
	assert.Equal(t, 0, system.Flags.D)
}

func TestCli(t *testing.T) {
	system := reset()
	system.Flags.I = 1
	system.Cli()
	assert.Equal(t, 0, system.Flags.I)
}

func TestClv(t *testing.T) {
	system := reset()
	system.Flags.V = 1
	system.Clv()
	assert.Equal(t, 0, system.Flags.V)
}

func TestCpx(t *testing.T) {
	system := reset()
	// TODO add test
	system.Cpx(0)
}

func TestCpy(t *testing.T) {
	system := reset()
	// TODO add test
	system.Cpy(0)
}

func TestDex(t *testing.T) {
	system := reset()
	system.X = 2
	system.Dex()
	assert.Equal(t, 1, system.X)
}

func TestDey(t *testing.T) {
	system := reset()
	system.Y = 2
	system.Dey()
	assert.Equal(t, 1, system.Y)
}

func TestEor(t *testing.T) {
	system := reset()
	// TODO add test
	system.Eor(0)
}

func TestInx(t *testing.T) {
	system := reset()
	system.Inx()
	assert.Equal(t, 1, system.X)
}

func TestIny(t *testing.T) {
	system := reset()
	system.Iny()
	assert.Equal(t, 1, system.Y)
}

func TestLda(t *testing.T) {
	system := reset()
	system.Lda(1)
	assert.Equal(t, 1, system.A)
}

func TestLdx(t *testing.T) {
	system := reset()
	system.Ldx(1)
	assert.Equal(t, 1, system.X)
}

func TestLdy(t *testing.T) {
	system := reset()
	system.Ldy(1)
	assert.Equal(t, 1, system.Y)
}

func TestLsr(t *testing.T) {
	system := reset()
	// TODO add test
	system.Lsr()
}

func TestNop(t *testing.T) {
	system := reset()
	system.Nop()
}

func TestOra(t *testing.T) {
	system := reset()
	// TODO add test
	system.Ora(0)
}

func TestPha(t *testing.T) {
	system := reset()
	// TODO add test
	system.Pha()
}

func TestPhp(t *testing.T) {
	system := reset()
	// TODO add test
	system.Php()
}

func TestPla(t *testing.T) {
	system := reset()
	// TODO add test
	system.Lsr()
}

func TestPlp(t *testing.T) {
	system := reset()
	// TODO add test
	system.Plp()
}

func TestRol(t *testing.T) {
	system := reset()
	// TODO add test
	system.Rol()
}

func TestRor(t *testing.T) {
	system := reset()
	// TODO add test
	system.Ror()
}

func TestRti(t *testing.T) {
	system := reset()
	system.Rti()
}

func TestSbc(t *testing.T) {
	system := reset()
	system.A = 2
	system.Sbc(0xff)
	assert.Equal(t, 2, system.A)
	assert.Equal(t, 0, system.Flags.C)

	system.Sbc(2)
	assert.Equal(t, 0xff, system.A)
	assert.Equal(t, 0, system.Flags.C)
}

func TestSec(t *testing.T) {
	system := reset()
	system.Sec()
	assert.Equal(t, 1, system.Flags.C)
}

func TestSed(t *testing.T) {
	system := reset()
	system.Sed()
	assert.Equal(t, 1, system.Flags.D)
}

func TestSei(t *testing.T) {
	system := reset()
	system.Sei()
	assert.Equal(t, 1, system.Flags.I)
}

func TestSta(t *testing.T) {
	system := reset()
	system.A = 11
	system.Sta(0)
	b := readMemory(0)
	assert.Equal(t, system.A, b)

	system.X = 0x22
	system.Sta(Absolute(0), X)
	b = readMemory(0x22)
	assert.Equal(t, system.A, b)
}

func TestStx(t *testing.T) {
	system := reset()
	system.X = 11
	system.Stx(0)
	b := readMemory(0)
	assert.Equal(t, system.X, b)

	system.Y = 0x22
	system.Stx(Absolute(0), Y)
	b = readMemory(0x22)
	assert.Equal(t, system.X, b)
}

func TestSty(t *testing.T) {
	system := reset()
	system.Y = 11
	system.Sty(0)
	b := readMemory(0)
	assert.Equal(t, system.Y, b)

	system.X = 0x22
	system.Sty(Absolute(0), X)
	b = readMemory(0x22)
	assert.Equal(t, system.Y, b)
}

func TestTax(t *testing.T) {
	system := reset()
	system.A = 2
	system.Tax()
	assert.Equal(t, system.A, system.X)
}

func TestTay(t *testing.T) {
	system := reset()
	system.A = 2
	system.Tay()
	assert.Equal(t, system.A, system.Y)
}

func TestTsx(t *testing.T) {
	system := reset()
	system.Tsx()
	assert.Equal(t, initialStack, system.SP)
}

func TestTxa(t *testing.T) {
	system := reset()
	system.X = 2
	system.Txa()
	assert.Equal(t, system.X, system.A)
}

func TestTxs(t *testing.T) {
	system := reset()
	system.X = 2
	system.Txs()
	assert.Equal(t, system.X, system.SP)
}

func TestTya(t *testing.T) {
	system := reset()
	system.Y = 2
	system.Tya()
	assert.Equal(t, system.Y, system.A)
}
