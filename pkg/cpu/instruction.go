package cpu

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
)

// AddressingInfo contains the opcode and timing info for an instruction addressing mode.
type AddressingInfo struct {
	Opcode byte
}

// Instruction contains information about a NES CPU instruction.
type Instruction struct {
	Name       string
	Unofficial bool

	// instruction has no parameters
	NoParamFunc func()
	// instruction has parameters
	ParamFunc func(params ...interface{})

	// maps addressing mode to cpu cycles
	Addressing map[Mode]AddressingInfo
}

// HasAddressing returns whether the instruction has any of the passed addressing modes.
func (c Instruction) HasAddressing(flags ...Mode) bool {
	for _, flag := range flags {
		_, ok := c.Addressing[flag]
		if ok {
			return ok
		}
	}
	return false
}

// LinkInstructionFuncs links cpu instruction emulation functions
// to the CPU instance.
// nolint: funlen
func LinkInstructionFuncs(cpu *CPU) {
	adc.ParamFunc = cpu.Adc
	and.ParamFunc = cpu.And
	asl.ParamFunc = cpu.Asl
	bcc.ParamFunc = cpu.BccInternal
	bcs.ParamFunc = cpu.BcsInternal
	beq.ParamFunc = cpu.BeqInternal
	bit.ParamFunc = cpu.Bit
	bmi.ParamFunc = cpu.BmiInternal
	bne.ParamFunc = cpu.BneInternal
	bpl.ParamFunc = cpu.BplInternal
	brk.NoParamFunc = cpu.Brk
	bvc.ParamFunc = cpu.BvcInternal
	bvs.ParamFunc = cpu.BvsInternal
	clc.NoParamFunc = cpu.Clc
	cld.NoParamFunc = cpu.Cld
	cli.NoParamFunc = cpu.Cli
	clv.NoParamFunc = cpu.Clv
	cmp.ParamFunc = cpu.Cmp
	cpx.ParamFunc = cpu.Cpx
	cpy.ParamFunc = cpu.Cpy
	dec.ParamFunc = cpu.Dec
	dex.NoParamFunc = cpu.Dex
	dey.NoParamFunc = cpu.Dey
	eor.ParamFunc = cpu.Eor
	inc.ParamFunc = cpu.Inc
	inx.NoParamFunc = cpu.Inx
	iny.NoParamFunc = cpu.Iny
	jmp.ParamFunc = cpu.Jmp
	jsr.ParamFunc = cpu.Jsr
	lda.ParamFunc = cpu.Lda
	ldx.ParamFunc = cpu.Ldx
	ldy.ParamFunc = cpu.Ldy
	lsr.ParamFunc = cpu.Lsr
	nop.NoParamFunc = cpu.Nop
	ora.ParamFunc = cpu.Ora
	pha.NoParamFunc = cpu.Pha
	php.NoParamFunc = cpu.Php
	pla.NoParamFunc = cpu.Pla
	plp.NoParamFunc = cpu.Plp
	rol.ParamFunc = cpu.Rol
	ror.ParamFunc = cpu.Ror
	rti.NoParamFunc = cpu.RtiInternal
	rts.NoParamFunc = cpu.Rts
	sbc.ParamFunc = cpu.Sbc
	sec.NoParamFunc = cpu.Sec
	sed.NoParamFunc = cpu.Sed
	sei.NoParamFunc = cpu.Sei
	sta.ParamFunc = cpu.Sta
	stx.ParamFunc = cpu.Stx
	sty.ParamFunc = cpu.Sty
	tax.NoParamFunc = cpu.Tax
	tay.NoParamFunc = cpu.Tay
	tsx.NoParamFunc = cpu.Tsx
	txa.NoParamFunc = cpu.Txa
	txs.NoParamFunc = cpu.Txs
	tya.NoParamFunc = cpu.Tya

	linkUnofficialInstructionFuncs(cpu)
}

var adc = &Instruction{
	Name: "adc",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0x69},
		ZeroPageAddressing:  {Opcode: 0x65},
		ZeroPageXAddressing: {Opcode: 0x75},
		AbsoluteAddressing:  {Opcode: 0x6d},
		AbsoluteXAddressing: {Opcode: 0x7d},
		AbsoluteYAddressing: {Opcode: 0x79},
		IndirectXAddressing: {Opcode: 0x61},
		IndirectYAddressing: {Opcode: 0x71},
	},
}

var and = &Instruction{
	Name: "and",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0x29},
		ZeroPageAddressing:  {Opcode: 0x25},
		ZeroPageXAddressing: {Opcode: 0x35},
		AbsoluteAddressing:  {Opcode: 0x2d},
		AbsoluteXAddressing: {Opcode: 0x3d},
		AbsoluteYAddressing: {Opcode: 0x39},
		IndirectXAddressing: {Opcode: 0x21},
		IndirectYAddressing: {Opcode: 0x31},
	},
}

var asl = &Instruction{
	Name: "asl",
	Addressing: map[Mode]AddressingInfo{
		AccumulatorAddressing: {Opcode: 0x0a},
		ZeroPageAddressing:    {Opcode: 0x06},
		ZeroPageXAddressing:   {Opcode: 0x16},
		AbsoluteAddressing:    {Opcode: 0x0e},
		AbsoluteXAddressing:   {Opcode: 0x1e},
	},
}

var bcc = &Instruction{
	Name: "bcc",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0x90},
	},
}

var bcs = &Instruction{
	Name: "bcs",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0xb0},
	},
}

var beq = &Instruction{
	Name: "beq",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0xf0},
	},
}

var bit = &Instruction{
	Name: "bit",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing: {Opcode: 0x24},
		AbsoluteAddressing: {Opcode: 0x2c},
	},
}

var bmi = &Instruction{
	Name: "bmi",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0x30},
	},
}

var bne = &Instruction{
	Name: "bne",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0xd0},
	},
}

var bpl = &Instruction{
	Name: "bpl",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0x10},
	},
}

var brk = &Instruction{
	Name: "brk",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x00},
	},
}

var bvc = &Instruction{
	Name: "bvc",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0x50},
	},
}

var bvs = &Instruction{
	Name: "bvs",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0x70},
	},
}

var clc = &Instruction{
	Name: "clc",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x18},
	},
}

var cld = &Instruction{
	Name: "cld",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xd8},
	},
}

var cli = &Instruction{
	Name: "cli",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x58},
	},
}

var clv = &Instruction{
	Name: "clv",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xb8},
	},
}

var cmp = &Instruction{
	Name: "cmp",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xc9},
		ZeroPageAddressing:  {Opcode: 0xc5},
		ZeroPageXAddressing: {Opcode: 0xd5},
		AbsoluteAddressing:  {Opcode: 0xcd},
		AbsoluteXAddressing: {Opcode: 0xdd},
		AbsoluteYAddressing: {Opcode: 0xd9},
		IndirectXAddressing: {Opcode: 0xc1},
		IndirectYAddressing: {Opcode: 0xd1},
	},
}

var cpx = &Instruction{
	Name: "cpx",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xe0},
		ZeroPageAddressing:  {Opcode: 0xe4},
		AbsoluteAddressing:  {Opcode: 0xec},
	},
}

var cpy = &Instruction{
	Name: "cpy",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xc0},
		ZeroPageAddressing:  {Opcode: 0xc4},
		AbsoluteAddressing:  {Opcode: 0xcc},
	},
}

var dec = &Instruction{
	Name: "dec",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0xc6},
		ZeroPageXAddressing: {Opcode: 0xd6},
		AbsoluteAddressing:  {Opcode: 0xce},
		AbsoluteXAddressing: {Opcode: 0xde},
	},
}

var dex = &Instruction{
	Name: "dex",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xca},
	},
}

var dey = &Instruction{
	Name: "dey",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x88},
	},
}

var eor = &Instruction{
	Name: "eor",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0x49},
		ZeroPageAddressing:  {Opcode: 0x45},
		ZeroPageXAddressing: {Opcode: 0x55},
		AbsoluteAddressing:  {Opcode: 0x4d},
		AbsoluteXAddressing: {Opcode: 0x5d},
		AbsoluteYAddressing: {Opcode: 0x59},
		IndirectXAddressing: {Opcode: 0x41},
		IndirectYAddressing: {Opcode: 0x51},
	},
}

var inc = &Instruction{
	Name: "inc",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0xe6},
		ZeroPageXAddressing: {Opcode: 0xf6},
		AbsoluteAddressing:  {Opcode: 0xee},
		AbsoluteXAddressing: {Opcode: 0xfe},
	},
}

var inx = &Instruction{
	Name: "inx",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xe8},
	},
}

var iny = &Instruction{
	Name: "iny",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xc8},
	},
}

var jmp = &Instruction{
	Name: "jmp",
	Addressing: map[Mode]AddressingInfo{
		AbsoluteAddressing: {Opcode: 0x4c},
		IndirectAddressing: {Opcode: 0x6c},
	},
}

var jsr = &Instruction{
	Name: "jsr",
	Addressing: map[Mode]AddressingInfo{
		AbsoluteAddressing: {Opcode: 0x20},
	},
}

var lda = &Instruction{
	Name: "lda",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xa9},
		ZeroPageAddressing:  {Opcode: 0xa5},
		ZeroPageXAddressing: {Opcode: 0xb5},
		AbsoluteAddressing:  {Opcode: 0xad},
		AbsoluteXAddressing: {Opcode: 0xbd},
		AbsoluteYAddressing: {Opcode: 0xb9},
		IndirectXAddressing: {Opcode: 0xa1},
		IndirectYAddressing: {Opcode: 0xb1},
	},
}

var ldx = &Instruction{
	Name: "ldx",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xa2},
		ZeroPageAddressing:  {Opcode: 0xa6},
		ZeroPageYAddressing: {Opcode: 0xb6},
		AbsoluteAddressing:  {Opcode: 0xae},
		AbsoluteYAddressing: {Opcode: 0xbe},
	},
}

var ldy = &Instruction{
	Name: "ldy",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xa0},
		ZeroPageAddressing:  {Opcode: 0xa4},
		ZeroPageXAddressing: {Opcode: 0xb4},
		AbsoluteAddressing:  {Opcode: 0xac},
		AbsoluteXAddressing: {Opcode: 0xbc},
	},
}

var lsr = &Instruction{
	Name: "lsr",
	Addressing: map[Mode]AddressingInfo{
		AccumulatorAddressing: {Opcode: 0x4a},
		ZeroPageAddressing:    {Opcode: 0x46},
		ZeroPageXAddressing:   {Opcode: 0x56},
		AbsoluteAddressing:    {Opcode: 0x4e},
		AbsoluteXAddressing:   {Opcode: 0x5e},
	},
}

var nop = &Instruction{
	Name: "nop",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xea},
	},
}

var ora = &Instruction{
	Name: "ora",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0x09},
		ZeroPageAddressing:  {Opcode: 0x05},
		ZeroPageXAddressing: {Opcode: 0x15},
		AbsoluteAddressing:  {Opcode: 0x0d},
		AbsoluteXAddressing: {Opcode: 0x1d},
		AbsoluteYAddressing: {Opcode: 0x19},
		IndirectXAddressing: {Opcode: 0x01},
		IndirectYAddressing: {Opcode: 0x11},
	},
}

var pha = &Instruction{
	Name: "pha",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x48},
	},
}

var php = &Instruction{
	Name: "php",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x08},
	},
}

var pla = &Instruction{
	Name: "pla",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x68},
	},
}

var plp = &Instruction{
	Name: "plp",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x28},
	},
}

var rol = &Instruction{
	Name: "rol",
	Addressing: map[Mode]AddressingInfo{
		AccumulatorAddressing: {Opcode: 0x2a},
		ZeroPageAddressing:    {Opcode: 0x26},
		ZeroPageXAddressing:   {Opcode: 0x36},
		AbsoluteAddressing:    {Opcode: 0x2e},
		AbsoluteXAddressing:   {Opcode: 0x3e},
	},
}

var ror = &Instruction{
	Name: "ror",
	Addressing: map[Mode]AddressingInfo{
		AccumulatorAddressing: {Opcode: 0x6a},
		ZeroPageAddressing:    {Opcode: 0x66},
		ZeroPageXAddressing:   {Opcode: 0x76},
		AbsoluteAddressing:    {Opcode: 0x6e},
		AbsoluteXAddressing:   {Opcode: 0x7e},
	},
}

var rti = &Instruction{
	Name: "rti",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x40},
	},
}

var rts = &Instruction{
	Name: "rts",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x60},
	},
}

var sbc = &Instruction{
	Name: "sbc",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xe9},
		ZeroPageAddressing:  {Opcode: 0xe5},
		ZeroPageXAddressing: {Opcode: 0xf5},
		AbsoluteAddressing:  {Opcode: 0xed},
		AbsoluteXAddressing: {Opcode: 0xfd},
		AbsoluteYAddressing: {Opcode: 0xf9},
		IndirectXAddressing: {Opcode: 0xe1},
		IndirectYAddressing: {Opcode: 0xf1},
	},
}

var sec = &Instruction{
	Name: "sec",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x38},
	},
}

var sed = &Instruction{
	Name: "sed",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xf8},
	},
}

var sei = &Instruction{
	Name: "sei",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x78},
	},
}

var sta = &Instruction{
	Name: "sta",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x85},
		ZeroPageXAddressing: {Opcode: 0x95},
		AbsoluteAddressing:  {Opcode: 0x8d},
		AbsoluteXAddressing: {Opcode: 0x9d},
		AbsoluteYAddressing: {Opcode: 0x99},
		IndirectXAddressing: {Opcode: 0x81},
		IndirectYAddressing: {Opcode: 0x91},
	},
}

var stx = &Instruction{
	Name: "stx",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x86},
		ZeroPageYAddressing: {Opcode: 0x96},
		AbsoluteAddressing:  {Opcode: 0x8e},
	},
}

var sty = &Instruction{
	Name: "sty",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x84},
		ZeroPageXAddressing: {Opcode: 0x94},
		AbsoluteAddressing:  {Opcode: 0x8c},
	},
}

var tax = &Instruction{
	Name: "tax",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xaa},
	},
}

var tay = &Instruction{
	Name: "tay",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xa8},
	},
}

var tsx = &Instruction{
	Name: "tsx",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xba},
	},
}

var txa = &Instruction{
	Name: "txa",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x8a},
	},
}

var txs = &Instruction{
	Name: "txs",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x9a},
	},
}

var tya = &Instruction{
	Name: "tya",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x98},
	},
}

// Instructions maps instruction names to NES CPU instruction information.
var Instructions = map[string]*Instruction{
	"adc": adc,
	"and": and,
	"asl": asl,
	"bcc": bcc,
	"bcs": bcs,
	"beq": beq,
	"bit": bit,
	"bmi": bmi,
	"bne": bne,
	"bpl": bpl,
	"brk": brk,
	"bvc": bvc,
	"bvs": bvs,
	"clc": clc,
	"cld": cld,
	"cli": cli,
	"clv": clv,
	"cmp": cmp,
	"cpx": cpx,
	"cpy": cpy,
	"dec": dec,
	"dex": dex,
	"dey": dey,
	"eor": eor,
	"inc": inc,
	"inx": inx,
	"iny": iny,
	"jmp": jmp,
	"jsr": jsr,
	"lda": lda,
	"ldx": ldx,
	"ldy": ldy,
	"lsr": lsr,
	"nop": nop,
	"ora": ora,
	"pha": pha,
	"php": php,
	"pla": pla,
	"plp": plp,
	"rol": rol,
	"ror": ror,
	"rti": rti,
	"rts": rts,
	"sbc": sbc,
	"sec": sec,
	"sed": sed,
	"sei": sei,
	"sta": sta,
	"stx": stx,
	"sty": sty,
	"tax": tax,
	"tay": tay,
	"tsx": tsx,
	"txa": txa,
	"txs": txs,
	"tya": tya,
}
