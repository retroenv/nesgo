package ast

// AddressingMode defines an address mode.
type AddressingMode int

// addressing modes.
const (
	NoAddressing      AddressingMode = 0
	ImpliedAddressing AddressingMode = 1 << iota
	AccumulatorAddressing
	ImmediateAddressing
	AbsoluteAddressing
	ZeroPageAddressing
	AbsoluteXAddressing
	ZeroPageXAddressing
	AbsoluteYAddressing
	ZeroPageYAddressing
	IndirectAddressing
	IndirectXAddressing
	IndirectYAddressing
	RelativeAddressing
)

// CPUInstruction contains information about a NES CPU instruction.
type CPUInstruction struct {
	Addressing AddressingMode
}

// CPUInstructions maps to NES CPU instruction information.
var CPUInstructions = map[string]*CPUInstruction{
	"adc": {Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing},
	"and": {Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing},
	"asl": {Addressing: AccumulatorAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing},
	"bcc": {Addressing: RelativeAddressing},
	"bcs": {Addressing: RelativeAddressing},
	"beq": {Addressing: RelativeAddressing},
	"bit": {Addressing: AbsoluteAddressing |
		ZeroPageAddressing},
	"bmi": {Addressing: RelativeAddressing},
	"bne": {Addressing: RelativeAddressing},
	"bpl": {Addressing: RelativeAddressing},
	"brk": {Addressing: ImpliedAddressing},
	"bvc": {Addressing: RelativeAddressing},
	"bvs": {Addressing: RelativeAddressing},
	"clc": {Addressing: ImpliedAddressing},
	"cld": {Addressing: ImpliedAddressing},
	"cli": {Addressing: ImpliedAddressing},
	"clv": {Addressing: ImpliedAddressing},
	"cmp": {Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing},
	"cpx": {Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		AbsoluteAddressing},
	"cpy": {Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		AbsoluteAddressing},
	"dec": {Addressing: ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing},
	"dex": {Addressing: ImpliedAddressing},
	"dey": {Addressing: ImpliedAddressing},
	"eor": {Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing},
	"inc": {Addressing: ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing},
	"inx": {Addressing: ImpliedAddressing},
	"iny": {Addressing: ImpliedAddressing},
	"jmp": {Addressing: AbsoluteAddressing |
		IndirectAddressing},
	"jsr": {Addressing: AbsoluteAddressing},
	"lda": {Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing},
	"ldx": {Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageYAddressing |
		AbsoluteAddressing |
		AbsoluteYAddressing},
	"ldy": {Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing},
	"lsr": {Addressing: AccumulatorAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing},
	"nop": {Addressing: ImpliedAddressing},
	"ora": {Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing},
	"pha": {Addressing: ImpliedAddressing},
	"php": {Addressing: ImpliedAddressing},
	"pla": {Addressing: ImpliedAddressing},
	"plp": {Addressing: ImpliedAddressing},
	"rol": {Addressing: AccumulatorAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing},
	"ror": {Addressing: AccumulatorAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing},
	"rti": {Addressing: ImpliedAddressing},
	"rts": {Addressing: ImpliedAddressing},
	"sbc": {Addressing: ImmediateAddressing |
		ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing},
	"sec": {Addressing: ImpliedAddressing},
	"sed": {Addressing: ImpliedAddressing},
	"sei": {Addressing: ImpliedAddressing},
	"sta": {Addressing: ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing |
		AbsoluteXAddressing |
		AbsoluteYAddressing |
		IndirectXAddressing |
		IndirectYAddressing},
	"stx": {Addressing: ZeroPageAddressing |
		ZeroPageYAddressing |
		AbsoluteAddressing},
	"sty": {Addressing: ZeroPageAddressing |
		ZeroPageXAddressing |
		AbsoluteAddressing},
	"tax": {Addressing: ImpliedAddressing},
	"tay": {Addressing: ImpliedAddressing},
	"tsx": {Addressing: ImpliedAddressing},
	"txa": {Addressing: ImpliedAddressing},
	"txs": {Addressing: ImpliedAddressing},
	"tya": {Addressing: ImpliedAddressing},
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
func (c CPUInstruction) HasAddressing(flags AddressingMode) bool {
	return c.Addressing&flags != 0
}
