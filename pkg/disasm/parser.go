package disasm

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/nesgo/pkg/disasm/param"
	"github.com/retroenv/nesgo/pkg/disasm/program"
	"github.com/retroenv/nesgo/pkg/nes"
)

// followExecutionFlow parses opcodes and follows the execution flow to parse all code.
func (dis *Disasm) followExecutionFlow() error {
	for len(dis.targetsToParse) > 0 {
		dis.popTarget()
		if *nes.PC == 0 {
			break
		}

		offset := dis.addressToOffset(*nes.PC)
		offsetInfo := &dis.offsets[offset]

		if offsetInfo.Type&program.CodeOffset != 0 {
			continue // was set by CDL
		}

		offsetInfo.OpcodeBytes = make([]byte, 1, 3)
		offsetInfo.OpcodeBytes[0] = dis.sys.ReadMemory(*nes.PC)

		if offsetInfo.Type&program.DataOffset != 0 {
			continue // was set by CDL
		}

		var err error
		offsetInfo.opcode, err = nes.DecodePCInstruction(dis.sys)
		if err != nil {
			// consider an unknown instruction as start of data
			offsetInfo.Type |= program.DataOffset
			continue
		}

		var params string
		instruction := offsetInfo.opcode.Instruction

		// process instructions with parameters, ignore special case of unofficial nop
		// that also has an implied addressing without parameters.
		if instruction.ParamFunc != nil && offsetInfo.opcode.Addressing != ImpliedAddressing {
			params, err = dis.processParamInstruction(*nes.PC, offsetInfo)
			if err != nil {
				return err
			}
		}

		opcodeLength := uint16(len(offsetInfo.OpcodeBytes))
		nextTarget := dis.sys.PC + opcodeLength

		if _, ok := cpu.NotExecutingFollowingOpcodeInstructions[instruction.Name]; !ok {
			dis.addTarget(nextTarget, instruction, false)
		}

		if params == "" {
			offsetInfo.Code = instruction.Name
		} else {
			offsetInfo.Code = fmt.Sprintf("%s %s", instruction.Name, params)
		}

		if instruction.Name == "nop" && instruction.Unofficial {
			dis.handleUnofficialNop(offset)
			continue
		}

		for i := uint16(0); i < opcodeLength && int(offset)+int(i) < len(dis.offsets); i++ {
			dis.offsets[offset+i].Type |= program.CodeOffset
		}
	}
	return nil
}

// processParamInstruction processes an instruction with parameters.
// Special handling is required as this instruction could branch to a different location.
func (dis *Disasm) processParamInstruction(offset uint16, offsetInfo *offset) (string, error) {
	opcode := offsetInfo.opcode
	params, opcodes, _ := nes.ReadOpParams(dis.sys, opcode.Addressing, false)
	offsetInfo.OpcodeBytes = append(offsetInfo.OpcodeBytes, opcodes...)

	paramAsString, err := param.String(dis.converter, opcode.Addressing, params...)
	if err != nil {
		return "", err
	}

	paramAsString = dis.replaceParamByAlias(offset, opcode, params[0], paramAsString)

	if _, ok := cpu.BranchingInstructions[opcode.Instruction.Name]; ok {
		addr, ok := params[0].(Absolute)
		if ok {
			dis.addTarget(uint16(addr), opcode.Instruction, true)
		}
	}

	return paramAsString, nil
}

// replaceParamByAlias replaces the absolute address with an alias name if it can match it to
// a constant, zero page variable or a code reference.
func (dis *Disasm) replaceParamByAlias(offset uint16, opcode cpu.Opcode, param interface{}, paramAsString string) string {
	if _, ok := cpu.BranchingInstructions[opcode.Instruction.Name]; ok {
		return paramAsString
	}

	var address uint16
	absolute, ok := param.(Absolute)
	if ok { // not the addressing type found that accesses known addresses
		address = uint16(absolute)
	} else {
		indirect, ok := param.(Indirect)
		if !ok {
			return paramAsString
		}
		address = uint16(indirect)
	}

	constantInfo, ok := dis.constants[address]
	if ok {
		return dis.replaceParamByConstant(opcode, paramAsString, address, constantInfo)
	}

	if !dis.addVariableReference(offset, opcode, address) {
		return paramAsString
	}

	// force using absolute address to not generate a different opcode by using zeropage access mode
	// TODO check if other assemblers use the same prefix
	switch opcode.Addressing {
	case ZeroPageAddressing, ZeroPageXAddressing, ZeroPageYAddressing:
		return "z:" + paramAsString
	case AbsoluteAddressing, AbsoluteXAddressing, AbsoluteYAddressing:
		return "a:" + paramAsString
	default: // indirect x, ...
		return paramAsString
	}
}

// popTarget pops the next target to disassemble and sets it into the program counter.
func (dis *Disasm) popTarget() {
	dis.sys.PC = dis.targetsToParse[0]
	dis.targetsToParse = dis.targetsToParse[1:]
}

// addTarget adds a target to the list to be processed if the address has not been processed yet.
func (dis *Disasm) addTarget(target uint16, currentInstruction *cpu.Instruction, jumpTarget bool) {
	offset := dis.addressToOffset(target)

	if currentInstruction != nil && currentInstruction.Name == "jsr" {
		dis.offsets[offset].Type |= program.CallTarget
	}
	if jumpTarget {
		dis.offsets[offset].JumpFrom = append(dis.offsets[offset].JumpFrom, *nes.PC)
		dis.jumpTargets[target] = struct{}{}
	}

	typ := dis.offsets[offset].Type
	if typ&program.CodeOffset != 0 || typ&program.DataOffset != 0 {
		return // already disassembled
	}
	dis.targetsToParse = append(dis.targetsToParse, target)
}
