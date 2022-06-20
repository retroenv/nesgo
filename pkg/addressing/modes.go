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

// AccessMode defines an address access mode.
type AccessMode int

// address accessing modes.
const (
	NoAccess        AccessMode = 0
	ReadAccess      AccessMode = 1
	WriteAccess     AccessMode = 2
	ReadWriteAccess AccessMode = 3
)

// AccessModeConstant is used to specify for every memory address what access mode applies to it.
// A memory address like 0x4017 has a different meaning depending on the type of access.
type AccessModeConstant struct {
	Constant string
	Mode     AccessMode
}
