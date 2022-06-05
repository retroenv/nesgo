package cpu

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
)

// AddressingInfo contains the opcode and timing info for an instruction addressing mode.
type AddressingInfo struct {
	Opcode byte
	Timing int
}

// Instruction contains information about a NES CPU instruction.
type Instruction struct {
	Name string

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
	rti.NoParamFunc = cpu.Rti
	rts.NoParamFunc = cpu.Rts
	sbc.ParamFunc = cpu.Sbc
	sec.NoParamFunc = cpu.Sec
	sed.NoParamFunc = cpu.Sed
	sei.NoParamFunc = cpu.Sei
	sta.ParamFunc = cpu.Sta
	stx.ParamFunc = cpu.Stx
	sty.ParamFunc = cpu.Sty
	tax.NoParamFunc = cpu.Tax
	tay.NoParamFunc = cpu.Tax
	tsx.NoParamFunc = cpu.Tsx
	txa.NoParamFunc = cpu.Txa
	txs.NoParamFunc = cpu.Txs
	tya.NoParamFunc = cpu.Tya
}

var adc = &Instruction{
	Name: "adc",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0x69, Timing: 2},
		ZeroPageAddressing:  {Opcode: 0x65, Timing: 3},
		ZeroPageXAddressing: {Opcode: 0x75, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0x6d, Timing: 4},
		AbsoluteXAddressing: {Opcode: 0x7d, Timing: 4},
		AbsoluteYAddressing: {Opcode: 0x79, Timing: 4},
		IndirectXAddressing: {Opcode: 0x61, Timing: 6},
		IndirectYAddressing: {Opcode: 0x71, Timing: 5},
	},
}

var and = &Instruction{
	Name: "and",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0x29, Timing: 2},
		ZeroPageAddressing:  {Opcode: 0x25, Timing: 3},
		ZeroPageXAddressing: {Opcode: 0x35, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0x2d, Timing: 4},
		AbsoluteXAddressing: {Opcode: 0x3d, Timing: 4},
		AbsoluteYAddressing: {Opcode: 0x39, Timing: 4},
		IndirectXAddressing: {Opcode: 0x21, Timing: 6},
		IndirectYAddressing: {Opcode: 0x31, Timing: 5},
	},
}

var asl = &Instruction{
	Name: "asl",
	Addressing: map[Mode]AddressingInfo{
		AccumulatorAddressing: {Opcode: 0x0a, Timing: 2},
		ZeroPageAddressing:    {Opcode: 0x06, Timing: 5},
		ZeroPageXAddressing:   {Opcode: 0x16, Timing: 6},
		AbsoluteAddressing:    {Opcode: 0x0e, Timing: 6},
		AbsoluteXAddressing:   {Opcode: 0x1e, Timing: 7},
	},
}

var bcc = &Instruction{
	Name: "bcc",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0x90, Timing: 3},
	},
}

var bcs = &Instruction{
	Name: "bcs",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0xb0, Timing: 3},
	},
}

var beq = &Instruction{
	Name: "beq",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0xf0, Timing: 3},
	},
}

var bit = &Instruction{
	Name: "bit",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing: {Opcode: 0x24, Timing: 3},
		AbsoluteAddressing: {Opcode: 0x2c, Timing: 4},
	},
}

var bmi = &Instruction{
	Name: "bmi",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0x30, Timing: 3},
	},
}

var bne = &Instruction{
	Name: "bne",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0xd0, Timing: 3},
	},
}

var bpl = &Instruction{
	Name: "bpl",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0x10, Timing: 3},
	},
}

var brk = &Instruction{
	Name: "brk",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x00, Timing: 7},
	},
}

var bvc = &Instruction{
	Name: "bvc",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0x50, Timing: 3},
	},
}

var bvs = &Instruction{
	Name: "bvs",
	Addressing: map[Mode]AddressingInfo{
		RelativeAddressing: {Opcode: 0x70, Timing: 3},
	},
}

var clc = &Instruction{
	Name: "clc",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x18, Timing: 2},
	},
}

var cld = &Instruction{
	Name: "cld",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xd8, Timing: 2},
	},
}

var cli = &Instruction{
	Name: "cli",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x58, Timing: 2},
	},
}

var clv = &Instruction{
	Name: "clv",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xb8, Timing: 2},
	},
}

var cmp = &Instruction{
	Name: "cmp",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xc9, Timing: 2},
		ZeroPageAddressing:  {Opcode: 0xc5, Timing: 3},
		ZeroPageXAddressing: {Opcode: 0xd5, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0xcd, Timing: 4},
		AbsoluteXAddressing: {Opcode: 0xdd, Timing: 4},
		AbsoluteYAddressing: {Opcode: 0xd9, Timing: 4},
		IndirectXAddressing: {Opcode: 0xc1, Timing: 6},
		IndirectYAddressing: {Opcode: 0xd1, Timing: 5},
	},
}

var cpx = &Instruction{
	Name: "cpx",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xe0, Timing: 2},
		ZeroPageAddressing:  {Opcode: 0xe4, Timing: 3},
		AbsoluteAddressing:  {Opcode: 0xec, Timing: 4},
	},
}

var cpy = &Instruction{
	Name: "cpy",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xc0, Timing: 2},
		ZeroPageAddressing:  {Opcode: 0xc4, Timing: 3},
		AbsoluteAddressing:  {Opcode: 0xcc, Timing: 4},
	},
}

var dec = &Instruction{
	Name: "dec",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0xc6, Timing: 5},
		ZeroPageXAddressing: {Opcode: 0xd6, Timing: 6},
		AbsoluteAddressing:  {Opcode: 0xce, Timing: 6},
		AbsoluteXAddressing: {Opcode: 0xde, Timing: 7},
	},
}

var dex = &Instruction{
	Name: "dex",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xca, Timing: 2},
	},
}

var dey = &Instruction{
	Name: "dey",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x88, Timing: 2},
	},
}

var eor = &Instruction{
	Name: "eor",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0x49, Timing: 2},
		ZeroPageAddressing:  {Opcode: 0x45, Timing: 3},
		ZeroPageXAddressing: {Opcode: 0x55, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0x4d, Timing: 4},
		AbsoluteXAddressing: {Opcode: 0x5d, Timing: 4},
		AbsoluteYAddressing: {Opcode: 0x59, Timing: 4},
		IndirectXAddressing: {Opcode: 0x41, Timing: 6},
		IndirectYAddressing: {Opcode: 0x51, Timing: 5},
	},
}

var inc = &Instruction{
	Name: "inc",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0xe6, Timing: 5},
		ZeroPageXAddressing: {Opcode: 0xf6, Timing: 6},
		AbsoluteAddressing:  {Opcode: 0xee, Timing: 6},
		AbsoluteXAddressing: {Opcode: 0xfe, Timing: 7},
	},
}

var inx = &Instruction{
	Name: "inx",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xe8, Timing: 2},
	},
}

var iny = &Instruction{
	Name: "iny",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xc8, Timing: 2},
	},
}

var jmp = &Instruction{
	Name: "jmp",
	Addressing: map[Mode]AddressingInfo{
		AbsoluteAddressing: {Opcode: 0x4c, Timing: 3},
		IndirectAddressing: {Opcode: 0x6c, Timing: 5},
	},
}

var jsr = &Instruction{
	Name: "jsr",
	Addressing: map[Mode]AddressingInfo{
		AbsoluteAddressing: {Opcode: 0x20, Timing: 6},
	},
}

var lda = &Instruction{
	Name: "lda",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xa9, Timing: 2},
		ZeroPageAddressing:  {Opcode: 0xa5, Timing: 3},
		ZeroPageXAddressing: {Opcode: 0xb5, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0xad, Timing: 4},
		AbsoluteXAddressing: {Opcode: 0xbd, Timing: 4},
		AbsoluteYAddressing: {Opcode: 0xb9, Timing: 4},
		IndirectXAddressing: {Opcode: 0xa1, Timing: 6},
		IndirectYAddressing: {Opcode: 0xb1, Timing: 5},
	},
}

var ldx = &Instruction{
	Name: "ldx",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xa2, Timing: 2},
		ZeroPageAddressing:  {Opcode: 0xa6, Timing: 3},
		ZeroPageYAddressing: {Opcode: 0xb6, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0xae, Timing: 4},
		AbsoluteYAddressing: {Opcode: 0xbe, Timing: 4},
	},
}

var ldy = &Instruction{
	Name: "ldy",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xa0, Timing: 2},
		ZeroPageAddressing:  {Opcode: 0xa4, Timing: 3},
		ZeroPageXAddressing: {Opcode: 0xb4, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0xac, Timing: 4},
		AbsoluteXAddressing: {Opcode: 0xbc, Timing: 4},
	},
}

var lsr = &Instruction{
	Name: "lsr",
	Addressing: map[Mode]AddressingInfo{
		AccumulatorAddressing: {Opcode: 0x4a, Timing: 2},
		ZeroPageAddressing:    {Opcode: 0x46, Timing: 5},
		ZeroPageXAddressing:   {Opcode: 0x56, Timing: 6},
		AbsoluteAddressing:    {Opcode: 0x4e, Timing: 6},
		AbsoluteXAddressing:   {Opcode: 0x5e, Timing: 7},
	},
}

var nop = &Instruction{
	Name: "nop",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xea, Timing: 2},
	},
}

var ora = &Instruction{
	Name: "ora",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0x09, Timing: 2},
		ZeroPageAddressing:  {Opcode: 0x05, Timing: 3},
		ZeroPageXAddressing: {Opcode: 0x15, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0x0d, Timing: 4},
		AbsoluteXAddressing: {Opcode: 0x1d, Timing: 4},
		AbsoluteYAddressing: {Opcode: 0x19, Timing: 4},
		IndirectXAddressing: {Opcode: 0x01, Timing: 6},
		IndirectYAddressing: {Opcode: 0x11, Timing: 5},
	},
}

var pha = &Instruction{
	Name: "pha",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x48, Timing: 3},
	},
}

var php = &Instruction{
	Name: "php",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x08, Timing: 3},
	},
}

var pla = &Instruction{
	Name: "pla",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x68, Timing: 4},
	},
}

var plp = &Instruction{
	Name: "plp",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x28, Timing: 4},
	},
}

var rol = &Instruction{
	Name: "rol",
	Addressing: map[Mode]AddressingInfo{
		AccumulatorAddressing: {Opcode: 0x2a, Timing: 2},
		ZeroPageAddressing:    {Opcode: 0x26, Timing: 5},
		ZeroPageXAddressing:   {Opcode: 0x36, Timing: 6},
		AbsoluteAddressing:    {Opcode: 0x2e, Timing: 6},
		AbsoluteXAddressing:   {Opcode: 0x3e, Timing: 7},
	},
}

var ror = &Instruction{
	Name: "ror",
	Addressing: map[Mode]AddressingInfo{
		AccumulatorAddressing: {Opcode: 0x6a, Timing: 2},
		ZeroPageAddressing:    {Opcode: 0x66, Timing: 5},
		ZeroPageXAddressing:   {Opcode: 0x76, Timing: 6},
		AbsoluteAddressing:    {Opcode: 0x6e, Timing: 6},
		AbsoluteXAddressing:   {Opcode: 0x7e, Timing: 7},
	},
}

var rti = &Instruction{
	Name: "rti",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x40, Timing: 6},
	},
}

var rts = &Instruction{
	Name: "rts",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x60, Timing: 6},
	},
}

var sbc = &Instruction{
	Name: "sbc",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xe9, Timing: 2},
		ZeroPageAddressing:  {Opcode: 0xe5, Timing: 3},
		ZeroPageXAddressing: {Opcode: 0xf5, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0xed, Timing: 4},
		AbsoluteXAddressing: {Opcode: 0xfd, Timing: 4},
		AbsoluteYAddressing: {Opcode: 0xf9, Timing: 4},
		IndirectXAddressing: {Opcode: 0xe1, Timing: 6},
		IndirectYAddressing: {Opcode: 0xf1, Timing: 5},
	},
}

var sec = &Instruction{
	Name: "sec",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x38, Timing: 2},
	},
}

var sed = &Instruction{
	Name: "sed",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xf8, Timing: 2},
	},
}

var sei = &Instruction{
	Name: "sei",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x78, Timing: 2},
	},
}

var sta = &Instruction{
	Name: "sta",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x85, Timing: 3},
		ZeroPageXAddressing: {Opcode: 0x95, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0x8d, Timing: 4},
		AbsoluteXAddressing: {Opcode: 0x9d, Timing: 4},
		AbsoluteYAddressing: {Opcode: 0x99, Timing: 4},
		IndirectXAddressing: {Opcode: 0x81, Timing: 6},
		IndirectYAddressing: {Opcode: 0x91, Timing: 5},
	},
}

var stx = &Instruction{
	Name: "stx",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x86, Timing: 3},
		ZeroPageYAddressing: {Opcode: 0x96, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0x8e, Timing: 4},
	},
}

var sty = &Instruction{
	Name: "sty",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x84, Timing: 3},
		ZeroPageXAddressing: {Opcode: 0x94, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0x8c, Timing: 4},
	},
}

var tax = &Instruction{
	Name: "tax",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xaa, Timing: 2},
	},
}

var tay = &Instruction{
	Name: "tay",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xa8, Timing: 2},
	},
}

var tsx = &Instruction{
	Name: "tsx",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0xba, Timing: 2},
	},
}

var txa = &Instruction{
	Name: "txa",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x8a, Timing: 2},
	},
}

var txs = &Instruction{
	Name: "txs",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x9a, Timing: 2},
	},
}

var tya = &Instruction{
	Name: "tya",
	Addressing: map[Mode]AddressingInfo{
		ImpliedAddressing: {Opcode: 0x98, Timing: 2},
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
