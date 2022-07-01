package cpu

import (
	"fmt"
	"math"

	. "github.com/retroenv/nesgo/pkg/addressing"
)

// addressModeFromCall gets the addressing mode from the passed params.
func (c *CPU) addressModeFromCall(instruction *Instruction, params ...interface{}) Mode {
	if len(params) == 0 {
		mode := addressModeFromCallNoParam(instruction)
		return mode
	}

	firstParam := params[0]
	var register interface{}
	if len(params) > 1 {
		register = params[1]
	}

	switch address := firstParam.(type) {
	case int:
		return c.addressModeInt(address, instruction, firstParam, register)

	case uint8:
		return ImmediateAddressing

	case *uint8: // variable
		return ImmediateAddressing

	case Absolute:
		return c.addressModeAbsolute(instruction)

	case Indirect, IndirectResolved:
		return c.addressModeIndirect(register)

	case ZeroPage:
		return c.addressModeZeroPage(register)

	case Accumulator:
		return AccumulatorAddressing

	default:
		panic(fmt.Sprintf("unsupported addressing mode type %T", firstParam))
	}
}

func (c *CPU) addressModeInt(address int, instruction *Instruction, firstParam, register interface{}) Mode {
	if instruction.HasAddressing(ImmediateAddressing) && register == nil && address <= math.MaxUint8 {
		return ImmediateAddressing
	}
	if register == nil {
		return AbsoluteAddressing
	}
	panic(fmt.Sprintf("unsupported int parameter %v", firstParam))
}

func (c *CPU) addressModeAbsolute(instruction *Instruction) Mode {
	// branches in emulation mode
	if instruction.HasAddressing(RelativeAddressing) {
		return RelativeAddressing
	}

	return AbsoluteAddressing
}

func (c *CPU) addressModeIndirect(register interface{}) Mode {
	if register == nil {
		return IndirectAddressing
	}

	ptr := register.(*uint8)
	switch ptr {
	case &c.X:
		return IndirectXAddressing
	case &c.Y:
		return IndirectYAddressing
	default:
		panic(fmt.Sprintf("unsupported indirect parameter %v", register))
	}
}

func (c *CPU) addressModeZeroPage(register interface{}) Mode {
	if register == 0 {
		return ZeroPageAddressing
	}

	ptr := register.(*uint8)
	switch ptr {
	case &c.X:
		return ZeroPageXAddressing
	case &c.Y:
		return ZeroPageYAddressing
	default:
		panic(fmt.Sprintf("unsupported zeropage parameter %v", register))
	}
}
