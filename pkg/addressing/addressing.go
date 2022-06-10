// Package addressing provides addressing helpers.
package addressing

// ZeroPage indicates that the parameter for the instruction is addressing
// the zero page.
type ZeroPage uint8

// Absolute indicates that the parameter for the instruction is an
// absolute address.
type Absolute uint16

// Indirect indicates that the parameter for the instruction is using
// indirect addressing using an address and an optional X or Y register.
// For usage with a register, the indirect address is a byte and refers
// to the zero page.
type Indirect uint16

// Accumulator indicates that the parameter for the instruction is the
// accumulator.
type Accumulator int
