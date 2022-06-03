package addressing

// ZeroPage indicates that the parameter for the instruction is addressing
// the zero page.
type ZeroPage uint8

// Absolute indicates that the parameter for the instruction is an
// absolute address.
type Absolute uint16

// Indirect indicates that the parameter for the instruction is using
// indirect addressing using an address and X or Y register.
type Indirect uint8
