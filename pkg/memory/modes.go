//go:build !nesgo

package memory

import (
	"fmt"
	"math"

	. "github.com/retroenv/nesgo/pkg/addressing"
)

// WriteAddressModes writes to memory using different address modes:
// Absolute: the absolut memory address is used to write the value
// Absolute, X: the absolut memory address with offset from X is used
// Absolute, Y: the absolut memory address with offset from Y is used
// (Indirect, X): the absolut memory address to write the value to is read from (indirect address + X)
// (Indirect), Y: the pointer to the memory address is read from the indirect parameter and adjusted after
// reading it by adding Y. The value is written to this pointer.
func (m *Memory) WriteAddressModes(value byte, params ...any) {
	param := params[0]
	var register any
	if len(params) > 1 {
		register = params[1]
	}

	switch address := param.(type) {
	case int:
		m.writeMemoryAbsolute(address, value, register)
	case *uint8: // variable
		*address = value
	case Absolute, AbsoluteX, AbsoluteY:
		m.writeMemoryAbsolute(address, value, register)
	case ZeroPage:
		m.writeMemoryZeroPage(address, value, register)
	case Indirect, IndirectResolved:
		m.writeMemoryIndirect(address, value, register)
	default:
		panic(fmt.Sprintf("unsupported memory write addressing mode type %T", param))
	}
}

func (m *Memory) writeMemoryIndirect(address any, value byte, register any) {
	pointer := m.indirectMemoryPointer(address, register)
	m.Write(pointer, value)
}

func (m *Memory) writeMemoryAbsolute(address any, value byte, register any) {
	if register == nil {
		m.writeMemoryAbsoluteOffset(address, value, 0)
		return
	}

	var offset uint16
	switch val := register.(type) {
	case *uint8: // X/Y register referenced in normal code
		offset = uint16(*val)
	case uint8: // X/Y register referenced in unit test as system.X
		offset = uint16(val)
	default:
		panic(fmt.Sprintf("unsupported extra parameter type %T for absolute memory write", register))
	}

	m.writeMemoryAbsoluteOffset(address, value, offset)
}

// Support 6502 bug, index will not leave zeropage when page boundary is crossed.
func (m *Memory) writeMemoryZeroPage(address ZeroPage, value byte, register any) {
	if register == nil {
		m.writeMemoryAbsoluteOffset(address, value, 0)
		return
	}

	var offset byte
	switch val := register.(type) {
	case *uint8: // X/Y register referenced in normal code
		offset = *val
	case uint8: // X/Y register referenced in unit test as system.X
		offset = val
	default:
		panic(fmt.Sprintf("unsupported extra parameter type %T for zero page memory write", register))
	}

	addr := uint16(byte(address) + offset)
	m.writeMemoryAbsoluteOffset(addr, value, 0)
}

func (m *Memory) writeMemoryAbsoluteOffset(address any, value byte, offset uint16) {
	switch addr := address.(type) {
	case int8:
		m.Write(uint16(addr)+offset, value)
	case uint8:
		m.Write(uint16(addr)+offset, value)
	case *uint8:
		*addr = value
	case uint16:
		m.Write(addr+offset, value)
	case *uint16:
		*addr = uint16(value)
	case int:
		m.Write(uint16(addr)+offset, value)
	case Absolute:
		m.Write(uint16(addr)+offset, value)
	case AbsoluteX:
		m.Write(uint16(addr)+offset, value)
	case AbsoluteY:
		m.Write(uint16(addr)+offset, value)
	case ZeroPage:
		m.Write(uint16(addr)+offset, value)
	default:
		panic(fmt.Sprintf("unsupported address type %T for absolute memory write with register", address))
	}
}

// ReadAddressModes reads memory using different address modes:
// Immediate: if immediate is true and the passed first param fits into a byte, it's immediate value is returned
// Absolute: the absolut memory address is used to read the value
// Absolute, X: the absolut memory address with offset from X is used
// Absolute, Y: the absolut memory address with offset from Y is used
// (Indirect, X): the absolut memory address to write the value to is read from (indirect address + X)
// (Indirect), Y: the pointer to the memory address is read from the indirect parameter and adjusted after
// reading it by adding Y. The value is read from this pointer.
func (m *Memory) ReadAddressModes(immediate bool, params ...any) byte {
	param := params[0]
	var register any
	if len(params) > 1 {
		register = params[1]
	}

	switch address := param.(type) {
	case int:
		if immediate && register == nil && address <= math.MaxUint8 {
			return uint8(address) // immediate, not an address
		}
		return m.ReadAbsolute(address, register)
	case uint8:
		return address // immediate, not an address
	case *uint8: // variable
		return *address
	case Absolute, AbsoluteX, AbsoluteY:
		return m.ReadAbsolute(address, register)
	case ZeroPage:
		return m.ReadMemoryZeroPage(address, register)
	case Indirect, IndirectResolved:
		return m.readMemoryIndirect(address, register)
	default:
		panic(fmt.Sprintf("unsupported memory read addressing mode type %T", param))
	}
}

// ReadAbsolute reads a byte from an address using absolute addressing.
func (m *Memory) ReadAbsolute(address any, register any) byte {
	if register == nil {
		return m.readAbsoluteOffset(address, 0)
	}

	var offset uint16
	switch val := register.(type) {
	case *uint8: // X/Y register referenced in normal code
		offset = uint16(*val)
	case uint8: // X/Y register referenced in unit test as system.X
		offset = uint16(val)
	default:
		panic(fmt.Sprintf("unsupported extra parameter type %T for absolute memory read", register))
	}
	return m.readAbsoluteOffset(address, offset)
}

// ReadMemoryZeroPage reads a byte from an address in zeropage using absolute addressing.
// Support 6502 bug, index will not leave zeropage when page boundary is crossed.
func (m *Memory) ReadMemoryZeroPage(address ZeroPage, register any) byte {
	if register == nil {
		return m.readAbsoluteOffset(address, 0)
	}

	var offset byte
	switch val := register.(type) {
	case *uint8: // X/Y register referenced in normal code
		offset = *val
	case uint8: // X/Y register referenced in unit test as system.X
		offset = val
	default:
		panic(fmt.Sprintf("unsupported extra parameter type %T for zero page memory read", register))
	}
	addr := uint16(byte(address) + offset)
	return m.readAbsoluteOffset(addr, 0)
}

func (m *Memory) readAbsoluteOffset(address any, offset uint16) byte {
	switch addr := address.(type) {
	case *uint8:
		if offset != 0 {
			panic("memory pointer read with offset is not supported")
		}
		return *addr
	case uint16:
		return m.Read(addr + offset)
	case int:
		return m.Read(uint16(addr) + offset)
	case Absolute:
		return m.Read(uint16(addr) + offset)
	case AbsoluteX:
		val := m.Read(uint16(addr))
		val += byte(offset)
		return val
	case AbsoluteY:
		val := m.Read(uint16(addr))
		val += byte(offset)
		return val
	case ZeroPage:
		return m.Read(uint16(addr) + offset)
	default:
		panic(fmt.Sprintf("unsupported address type %T for absolute memory write", address))
	}
}

func (m *Memory) readMemoryIndirect(address any, register any) byte {
	pointer := m.indirectMemoryPointer(address, register)
	return m.Read(pointer)
}

func (m *Memory) indirectMemoryPointer(addressParam any, register any) uint16 {
	if register == nil {
		panic("register parameter missing for indirect memory addressing")
	}

	p, ok := register.(*uint8)
	if !ok {
		panic(fmt.Sprintf("unsupported extra parameter type %T for indirect memory addressing", register))
	}

	var address uint16
	indirectAddress, ok := addressParam.(Indirect)
	if ok {
		address = uint16(indirectAddress)
		if address > 0xff {
			panic(fmt.Sprintf("indirect address parameter 0x%04X exceeds byte", address))
		}
	} else {
		address = uint16(addressParam.(IndirectResolved))
		return address
	}

	var pointer uint16
	switch {
	case p == m.globalX, p == m.x:
		pointer = m.ReadWord(address + uint16(*p))
	case p == m.globalY, p == m.y:
		pointer = m.ReadWord(address)
		pointer += uint16(*p)
	default:
		panic("only X and Y registers are supported for indirect addressing")
	}
	return pointer
}
