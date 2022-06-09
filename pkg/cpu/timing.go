//go:build !nesgo
// +build !nesgo

package cpu

import (
	"fmt"
	"math"
	"time"

	. "github.com/retroenv/nesgo/pkg/addressing"
)

// instructionHook is a hook that is executed before a CPU instruction is executed.
// It allows for accounting of the instruction timing and trace logging.
// TODO add option to disable timing in unit tests
func (c *CPU) instructionHook(instruction *Instruction, params ...interface{}) {
	if c.tracing != NoTracing {
		c.trace(instruction, params...)
	}

	// TODO account for exact cycles
	time.Sleep(time.Microsecond)
}

// addressModeFromCall gets the addressing mode from the passed params
func addressModeFromCall(instruction *Instruction, params ...interface{}) (Mode, string) {
	if len(params) == 0 {
		return addressModeFromCallNoParam(instruction)
	}

	param := params[0]
	var register interface{}
	if len(params) > 1 {
		register = params[1]
	}

	switch address := param.(type) {
	case int:
		if instruction.HasAddressing(ImmediateAddressing) && register == nil && address <= math.MaxUint8 {
			return ImmediateAddressing, fmt.Sprintf("#%00x", address)
		}
		if register == nil {
			return AbsoluteAddressing, fmt.Sprintf("#%0000x", address)
		}
		panic("X/Y support not implemented") // TODO

	case uint8:
		return ImmediateAddressing, fmt.Sprintf("#%00x", address)

	case *uint8: // variable
		return ImmediateAddressing, fmt.Sprintf("#%00x", *address)

	case Absolute:
		// branches in emulation mode
		if instruction.HasAddressing(RelativeAddressing) {
			return RelativeAddressing, fmt.Sprintf("#%0000x", address)
		}

		return AbsoluteAddressing, fmt.Sprintf("#%0000x", address)

	case Indirect:
		if register == nil {
			return IndirectAddressing, fmt.Sprintf("#%0000x", address)
		}
		panic("X/Y support not implemented") // TODO

	default:
		panic(fmt.Sprintf("unsupported addressing mode type %T", param))
	}
}

func addressModeFromCallNoParam(instruction *Instruction) (Mode, string) {
	if instruction.HasAddressing(AccumulatorAddressing) {
		return AccumulatorAddressing, ""
	}
	// branches have no target in go mode
	if instruction.HasAddressing(RelativeAddressing) {
		return RelativeAddressing, ""
	}
	return ImpliedAddressing, ""
}
