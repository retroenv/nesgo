package ast

// addressing modes.
const (
	NoAddressing      = 0
	ImpliedAddressing = 1 << iota
	AccumulatorAddressing
	ImmediateAddressing
	AbsoluteAddressing
	ZeroPageAddressing
	AbsoluteXAddressing
	ZeroPageXAddressing
	AbsoluteYAddressing
	ZeroPageYAddressing
	IndirectXAddressing
	IndirectYAddressing
	RelativeAddressing
)

// CPUInstruction contains information about a NES CPU instruction.
type CPUInstruction struct {
	Alias      string
	Addressing int
}

// CPUInstructions maps to NES CPU instruction information.
var CPUInstructions = map[string]*CPUInstruction{
	"adc": nil, // TODO
	"and": nil, // TODO
	"asl": {Alias: "asl", Addressing: AccumulatorAddressing |
		AbsoluteAddressing | ZeroPageAddressing | AbsoluteXAddressing | ZeroPageXAddressing},
	"bcc":     {Alias: "bcc", Addressing: RelativeAddressing},
	"bcs":     {Alias: "bcs", Addressing: RelativeAddressing},
	"beq":     {Alias: "beq", Addressing: RelativeAddressing},
	"bit":     {Alias: "bit", Addressing: AbsoluteAddressing | ZeroPageAddressing},
	"bmi":     {Alias: "bmi", Addressing: RelativeAddressing},
	"bne":     {Alias: "bne", Addressing: RelativeAddressing},
	"bpl":     {Alias: "bpl", Addressing: RelativeAddressing},
	"brk":     {Alias: "brk", Addressing: ImpliedAddressing},
	"bvc":     {Alias: "bvc", Addressing: RelativeAddressing},
	"bvs":     {Alias: "bvs", Addressing: RelativeAddressing},
	"clc":     {Alias: "clc", Addressing: ImpliedAddressing},
	"cld":     {Alias: "cld", Addressing: ImpliedAddressing},
	"cli":     {Alias: "cli", Addressing: ImpliedAddressing},
	"clv":     {Alias: "clv", Addressing: ImpliedAddressing},
	"cmp":     nil, // TODO
	"cpx":     {Alias: "cpx", Addressing: ImmediateAddressing},
	"cpxaddr": {Alias: "cpx", Addressing: AbsoluteAddressing | ZeroPageAddressing},
	"cpy":     {Alias: "cpy", Addressing: ImmediateAddressing},
	"cpyaddr": {Alias: "cpy", Addressing: AbsoluteAddressing | ZeroPageAddressing},
	"dex":     {Alias: "dex", Addressing: ImpliedAddressing},
	"dey":     {Alias: "dey", Addressing: ImpliedAddressing},
	"eor":     {Alias: "eor", Addressing: ImmediateAddressing},
	"eoraddr": {Alias: "eor", Addressing: AbsoluteAddressing |
		ZeroPageAddressing | AbsoluteXAddressing | AbsoluteYAddressing | ZeroPageXAddressing},
	"eorind": {Alias: "eor", Addressing: IndirectXAddressing |
		IndirectYAddressing},
	"inx": {Alias: "inx", Addressing: ImpliedAddressing},
	"iny": {Alias: "iny", Addressing: ImpliedAddressing},
	"jmp": {Alias: "jmp", Addressing: RelativeAddressing},
	"lda": {Alias: "lda", Addressing: ImmediateAddressing},
	"ldaaddr": {Alias: "lda", Addressing: AbsoluteAddressing |
		ZeroPageAddressing | AbsoluteXAddressing | AbsoluteYAddressing | ZeroPageXAddressing},
	"ldaind": {Alias: "lda", Addressing: IndirectXAddressing | IndirectYAddressing},
	"ldx":    {Alias: "ldx", Addressing: ImmediateAddressing},
	"ldxaddr": {Alias: "ldx", Addressing: AbsoluteAddressing |
		ZeroPageAddressing | AbsoluteYAddressing | ZeroPageYAddressing},
	"ldy": {Alias: "ldy", Addressing: ImmediateAddressing},
	"ldyaddr": {Alias: "ldy", Addressing: AbsoluteAddressing |
		ZeroPageAddressing | AbsoluteXAddressing | ZeroPageXAddressing},
	"lsr": {Alias: "lsr", Addressing: AccumulatorAddressing |
		AbsoluteAddressing | ZeroPageAddressing | AbsoluteXAddressing | ZeroPageXAddressing},
	"nop": {Alias: "nop", Addressing: ImpliedAddressing},
	"ora": nil, // TODO
	"pha": {Alias: "pha", Addressing: ImpliedAddressing},
	"php": {Alias: "php", Addressing: ImpliedAddressing},
	"pla": {Alias: "pla", Addressing: ImpliedAddressing},
	"plp": {Alias: "plp", Addressing: ImpliedAddressing},
	"rol": {Alias: "rol", Addressing: AccumulatorAddressing |
		AbsoluteAddressing | ZeroPageAddressing | AbsoluteXAddressing | ZeroPageXAddressing},
	"ror": {Alias: "ror", Addressing: AccumulatorAddressing |
		AbsoluteAddressing | ZeroPageAddressing | AbsoluteXAddressing | ZeroPageXAddressing},
	"rti": {Alias: "rti", Addressing: ImpliedAddressing},
	"rts": {Alias: "rts", Addressing: ImpliedAddressing},
	"sbc": nil, // TODO
	"sec": {Alias: "sec", Addressing: ImpliedAddressing},
	"sed": {Alias: "sed", Addressing: ImpliedAddressing},
	"sei": {Alias: "sei", Addressing: ImpliedAddressing},
	"sta": {Alias: "sta", Addressing: AbsoluteAddressing |
		ZeroPageAddressing | AbsoluteXAddressing | AbsoluteYAddressing | ZeroPageXAddressing},
	"staind": {Alias: "sta", Addressing: IndirectXAddressing |
		IndirectYAddressing},
	"stx": {Alias: "stx", Addressing: AbsoluteAddressing |
		ZeroPageAddressing | AbsoluteYAddressing},
	"sty": {Alias: "sty", Addressing: AbsoluteAddressing |
		ZeroPageAddressing | AbsoluteXAddressing},
	"tax": {Alias: "tax", Addressing: ImpliedAddressing},
	"tay": {Alias: "tay", Addressing: ImpliedAddressing},
	"tsx": {Alias: "tsx", Addressing: ImpliedAddressing},
	"txa": {Alias: "txa", Addressing: ImpliedAddressing},
	"txs": {Alias: "txs", Addressing: ImpliedAddressing},
	"tya": {Alias: "tya", Addressing: ImpliedAddressing},
}

// CPUBranchingInstructions ...
var CPUBranchingInstructions = map[string]struct{}{
	"bcc": {},
	"bcs": {},
	"beq": {},
	"bmi": {},
	"bne": {},
	"bpl": {},
	"bvc": {},
	"bvs": {},
	"jmp": {},
}

// CPURegisters ...
var CPURegisters = map[string]struct{}{
	"A": {},
	"X": {},
	"Y": {},
}

// HasAddressing ...
func (c CPUInstruction) HasAddressing(flags int) bool {
	return c.Addressing&flags != 0
}
