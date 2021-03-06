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
		offsetInfo, inspectCode := dis.initializeOffsetInfo(offset)
		if !inspectCode {
			continue
		}

		instruction := offsetInfo.opcode.Instruction

		// process instructions with parameters, ignore special case of unofficial nop
		// that also has an implied addressing without parameters.
		if instruction.ParamFunc != nil && offsetInfo.opcode.Addressing != ImpliedAddressing {
			params, err := dis.processParamInstruction(*nes.PC, offsetInfo)
			if err != nil {
				return err
			}
			offsetInfo.Code = fmt.Sprintf("%s %s", instruction.Name, params)
		} else {
			offsetInfo.Code = instruction.Name
		}

		opcodeLength := uint16(len(offsetInfo.OpcodeBytes))
		nextTarget := dis.sys.PC + opcodeLength

		if _, ok := cpu.NotExecutingFollowingOpcodeInstructions[instruction.Name]; !ok {
			dis.addTarget(nextTarget, instruction, false)
		}

		dis.checkInstructionOverlap(offsetInfo, offset)

		if instruction.Name == "nop" && instruction.Unofficial {
			dis.handleUnofficialNop(offset)
			continue
		}

		dis.changeOffsetRangeToCode(offsetInfo.OpcodeBytes, offset)
	}
	return nil
}

// in case the current instruction overlaps with an already existing instruction,
// cut the current one short.
func (dis *Disasm) checkInstructionOverlap(offsetInfo *offset, offset uint16) {
	for i := 1; i < len(offsetInfo.OpcodeBytes) && int(offset)+i < len(dis.offsets); i++ {
		ins := &dis.offsets[offset+uint16(i)]
		if ins.IsType(program.CodeOffset) {
			offsetInfo.OpcodeBytes = offsetInfo.OpcodeBytes[:i]
			return
		}
	}
}

// initializeOffsetInfo initializes the offset info for the given offset and returns
// whether the offset should process inspection for code parameters.
func (dis *Disasm) initializeOffsetInfo(offset uint16) (*offset, bool) {
	offsetInfo := &dis.offsets[offset]

	if offsetInfo.IsType(program.CodeOffset) {
		return offsetInfo, false // was set by CDL
	}

	offsetInfo.OpcodeBytes = make([]byte, 1, 3)
	offsetInfo.OpcodeBytes[0] = dis.sys.Bus.Memory.Read(*nes.PC)

	if offsetInfo.IsType(program.DataOffset) {
		return offsetInfo, false // was set by CDL
	}

	var err error
	offsetInfo.opcode, err = nes.DecodePCInstruction(dis.sys)
	if err != nil {
		// consider an unknown instruction as start of data
		offsetInfo.SetType(program.DataOffset)
		return offsetInfo, false
	}
	return offsetInfo, true
}

// processParamInstruction processes an instruction with parameters.
// Special handling is required as this instruction could branch to a different location.
func (dis *Disasm) processParamInstruction(offset uint16, offsetInfo *offset) (string, error) {
	opcode := offsetInfo.opcode
	params, opcodes, _ := nes.ReadOpParams(dis.sys.Bus.Memory, opcode.Addressing, false)
	offsetInfo.OpcodeBytes = append(offsetInfo.OpcodeBytes, opcodes...)

	firstParam := params[0]
	paramAsString, err := param.String(dis.converter, opcode.Addressing, firstParam)
	if err != nil {
		return "", err
	}

	paramAsString = dis.replaceParamByAlias(offset, opcode, firstParam, paramAsString)

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
	offsetInfo := &dis.offsets[offset]

	if currentInstruction != nil && currentInstruction.Name == "jsr" {
		offsetInfo.SetType(program.CallTarget)
	}
	if jumpTarget {
		offsetInfo.JumpFrom = append(offsetInfo.JumpFrom, *nes.PC)
		dis.jumpTargets[target] = struct{}{}
	}

	if _, ok := dis.targetsAdded[target]; ok {
		return
	}
	dis.targetsAdded[target] = struct{}{}
	dis.targetsToParse = append(dis.targetsToParse, target)
}
