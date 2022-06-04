package cpu

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
)

// Instruction contains information about a NES CPU instruction.
type Instruction struct {
	Name string

	// instruction has no parameters
	NoParamFunc func()
	// instruction has parameters
	ParamFunc func(params ...interface{})

	Addressing Mode // TODO change to map[Mode] byte for opcode + timing
}

// HasAddressing ...
func (c Instruction) HasAddressing(flags Mode) bool {
	return c.Addressing&flags != 0
}

// LinkInstructionFuncs links cpu instruction emulation functions
// to the CPU instance.
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
	Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing,
}

var and = &Instruction{
	Name: "and",
	Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing,
}

var asl = &Instruction{
	Name: "asl",
	Addressing: AccumulatorAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing,
}

var bcc = &Instruction{
	Name:       "bcc",
	Addressing: RelativeAddressing,
}

var bcs = &Instruction{
	Name:       "bcs",
	Addressing: RelativeAddressing,
}

var beq = &Instruction{
	Name:       "beq",
	Addressing: RelativeAddressing,
}

var bit = &Instruction{
	Name: "bit",
	Addressing: AbsoluteAddressing |
		ZeroPageAddressing,
}

var bmi = &Instruction{
	Name:       "bmi",
	Addressing: RelativeAddressing,
}

var bne = &Instruction{
	Name:       "bne",
	Addressing: RelativeAddressing,
}

var bpl = &Instruction{
	Name:       "bpl",
	Addressing: RelativeAddressing,
}

var brk = &Instruction{
	Name:       "brk",
	Addressing: ImpliedAddressing,
}

var bvc = &Instruction{
	Name:       "bvc",
	Addressing: RelativeAddressing,
}

var bvs = &Instruction{
	Name:       "bvs",
	Addressing: RelativeAddressing,
}

var clc = &Instruction{
	Name:       "clc",
	Addressing: ImpliedAddressing,
}

var cld = &Instruction{
	Name:       "cld",
	Addressing: ImpliedAddressing,
}

var cli = &Instruction{
	Name:       "cli",
	Addressing: ImpliedAddressing,
}

var clv = &Instruction{
	Name:       "clv",
	Addressing: ImpliedAddressing,
}

var cmp = &Instruction{
	Name: "cmp",
	Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing,
}

var cpx = &Instruction{
	Name: "cpx",
	Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		AbsoluteAddressing,
}

var cpy = &Instruction{
	Name: "cpy",
	Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		AbsoluteAddressing,
}

var dec = &Instruction{
	Name: "dec",
	Addressing: ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing,
}

var dex = &Instruction{
	Name:       "dex",
	Addressing: ImpliedAddressing,
}

var dey = &Instruction{
	Name:       "dey",
	Addressing: ImpliedAddressing,
}

var eor = &Instruction{
	Name: "eor",
	Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing,
}

var inc = &Instruction{
	Name: "inc",
	Addressing: ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing,
}

var inx = &Instruction{
	Name:       "inx",
	Addressing: ImpliedAddressing,
}

var iny = &Instruction{
	Name:       "iny",
	Addressing: ImpliedAddressing,
}

var jmp = &Instruction{
	Name: "jmp",
	Addressing: AbsoluteAddressing |
		IndirectAddressing,
}

var jsr = &Instruction{
	Name:       "jsr",
	Addressing: AbsoluteAddressing,
}

var lda = &Instruction{
	Name: "lda",
	Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing,
}

var ldx = &Instruction{
	Name: "ldx",
	Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageYAddressing |
		AbsoluteAddressing |
		AbsoluteYAddressing,
}

var ldy = &Instruction{
	Name: "ldy",
	Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing,
}

var lsr = &Instruction{
	Name: "lsr",
	Addressing: AccumulatorAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing,
}

var nop = &Instruction{
	Name:       "nop",
	Addressing: ImpliedAddressing,
}

var ora = &Instruction{
	Name: "ora",
	Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing,
}

var pha = &Instruction{
	Name:       "pha",
	Addressing: ImpliedAddressing,
}

var php = &Instruction{
	Name:       "php",
	Addressing: ImpliedAddressing,
}

var pla = &Instruction{
	Name:       "pla",
	Addressing: ImpliedAddressing,
}

var plp = &Instruction{
	Name:       "plp",
	Addressing: ImpliedAddressing,
}

var rol = &Instruction{
	Name: "rol",
	Addressing: AccumulatorAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing,
}

var ror = &Instruction{
	Name: "ror",
	Addressing: AccumulatorAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing,
}

var rti = &Instruction{
	Name:       "rti",
	Addressing: ImpliedAddressing,
}

var rts = &Instruction{
	Name:       "rts",
	Addressing: ImpliedAddressing,
}

var sbc = &Instruction{
	Name: "sbc",
	Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing,
}

var sec = &Instruction{
	Name:       "sec",
	Addressing: ImpliedAddressing,
}

var sed = &Instruction{
	Name:       "sed",
	Addressing: ImpliedAddressing,
}

var sei = &Instruction{
	Name:       "sei",
	Addressing: ImpliedAddressing,
}

var sta = &Instruction{
	Name: "sta",
	Addressing: ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing,
}

var stx = &Instruction{
	Name: "stx",
	Addressing: ZeroPageAddressing |
		ZeroPageYAddressing |
		AbsoluteAddressing,
}

var sty = &Instruction{
	Name: "sty",
	Addressing: ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing,
}

var tax = &Instruction{
	Name:       "tax",
	Addressing: ImpliedAddressing,
}

var tay = &Instruction{
	Name:       "tay",
	Addressing: ImpliedAddressing,
}

var tsx = &Instruction{
	Name:       "tsx",
	Addressing: ImpliedAddressing,
}

var txa = &Instruction{
	Name:       "txa",
	Addressing: ImpliedAddressing,
}

var txs = &Instruction{
	Name:       "txs",
	Addressing: ImpliedAddressing,
}

var tya = &Instruction{
	Name:       "tya",
	Addressing: ImpliedAddressing,
}

// Instructions maps to NES CPU instruction information.
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
