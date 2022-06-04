package addressing

// Mode defines an address mode.
type Mode int

// addressing modes.
const (
	NoAddressing      Mode = 0
	ImpliedAddressing Mode = 1 << iota
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
