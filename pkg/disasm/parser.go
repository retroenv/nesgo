package disasm

import (
	"fmt"
	"strings"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/nesgo/pkg/disasm/param"
	"github.com/retroenv/nesgo/pkg/nes"
)

// followExecutionFlow parses opcodes and follows the execution flow to parse all code.
func (dis *Disasm) followExecutionFlow() error {
	sys := dis.sys

	for len(dis.targetsToParse) > 0 {
		dis.popTarget()
		if *nes.PC == 0 {
			break
		}

		offset := dis.addressToOffset(*nes.PC)
		opcode, err := nes.DecodePCInstruction(sys)
		if err != nil {
			// consider unknown instruction as start of data
			dis.offsets[offset].IsProcessed = true
			continue
		}

		dis.offsets[offset].opcodeBytes = []byte{sys.ReadMemory(*nes.PC)}

		var params string

		// process instructions with parameters, ignore special case of unofficial nop
		// that also has an implied addressing without parameters.
		if opcode.Instruction.ParamFunc != nil && opcode.Addressing != ImpliedAddressing {
			var opcodes []byte
			opcodes, params, err = dis.processParamInstruction(opcode)
			if err != nil {
				return err
			}
			dis.offsets[offset].opcodeBytes = append(dis.offsets[offset].opcodeBytes, opcodes...)
		}

		opcodeLength := uint16(len(dis.offsets[offset].opcodeBytes))
		nextTarget := sys.PC + opcodeLength

		if _, ok := cpu.NotExecutingFollowingOpcodeInstructions[opcode.Instruction.Name]; !ok {
			dis.addTarget(nextTarget, opcode.Instruction, false)
		}

		dis.offsets[offset].opcode = opcode

		if params == "" {
			dis.offsets[offset].Output = opcode.Instruction.Name
		} else {
			dis.offsets[offset].Output = fmt.Sprintf("%s %s", opcode.Instruction.Name, params)
		}

		for i := uint16(0); i < opcodeLength && int(offset)+int(i) < len(dis.offsets); i++ {
			dis.offsets[offset+i].IsProcessed = true
		}
	}
	return nil
}

// processParamInstruction processes an instruction with parameters.
// Special handling is required as this instruction could branch to a different location.
func (dis *Disasm) processParamInstruction(opcode cpu.Opcode) ([]byte, string, error) {
	params, opcodes, _ := nes.ReadOpParams(dis.sys, opcode.Addressing, false)

	paramAsString, err := param.String(dis.converter, opcode.Addressing, params...)
	if err != nil {
		return nil, "", err
	}

	paramAsString = dis.replaceParamByConstant(opcode, params[0], paramAsString)

	if _, ok := cpu.BranchingInstructions[opcode.Instruction.Name]; ok {
		addr, ok := params[0].(Absolute)
		if ok {
			dis.addTarget(uint16(addr), opcode.Instruction, true)
		}
	}
	return opcodes, paramAsString, nil
}

// replaceParamByConstant replaces the absolute address with a constant name if it has a known
// translation for the access mode.
func (dis *Disasm) replaceParamByConstant(opcode cpu.Opcode, param interface{}, paramAsString string) string {
	addr, ok := param.(Absolute)
	if !ok { // not the addressing type found that accesses known addresses
		return paramAsString
	}

	// split parameter string in case of x/y indexing, only the first part will be replaced by a const name
	paramParts := strings.Split(paramAsString, ",")

	constantInfo, ok := dis.constants[uint16(addr)]
	if !ok { // not accessing a known address
		// force using absolute address to not generate a different opcode by using zeropage access mode
		// TODO check if other assemblers use the same prefix
		switch opcode.Addressing {
		case ZeroPageAddressing:
			return "z:" + paramAsString
		case AbsoluteAddressing:
			return "a:" + paramAsString
		default: // indirect x, ...
			return paramAsString
		}
	}

	if constantInfo.Read != "" {
		if _, ok := cpu.MemoryReadInstructions[opcode.Instruction.Name]; ok {
			dis.usedConstants[uint16(addr)] = constantInfo
			paramParts[0] = constantInfo.Read
			return strings.Join(paramParts, ",")
		}
	}
	if constantInfo.Write != "" {
		if _, ok := cpu.MemoryWriteInstructions[opcode.Instruction.Name]; ok {
			dis.usedConstants[uint16(addr)] = constantInfo
			paramParts[0] = constantInfo.Write
			return strings.Join(paramParts, ",")
		}
	}

	return paramAsString
}

// processJumpTargets processes all jump targetsToParse and updates the callers with
// the generated jump target label name.
func (dis *Disasm) processJumpTargets() {
	for target := range dis.jumpTargets {
		offset := dis.addressToOffset(target)
		name := dis.offsets[offset].Label
		if name == "" {
			if dis.offsets[offset].IsCallTarget {
				name = fmt.Sprintf("_func_%04x", target)
			} else {
				name = fmt.Sprintf("_label_%04x", target)
			}
			dis.offsets[offset].Label = name
		}

		for _, caller := range dis.offsets[offset].JumpFrom {
			offset = dis.addressToOffset(caller)
			dis.offsets[offset].Output = dis.offsets[offset].opcode.Instruction.Name
			dis.offsets[offset].JumpingTo = name
		}
	}
}

// addTarget adds a target to the list to be processed if the address has not been processed yet.
func (dis *Disasm) addTarget(target uint16, currentInstruction *cpu.Instruction, jumpTarget bool) {
	offset := dis.addressToOffset(target)

	if currentInstruction != nil && currentInstruction.Name == "jsr" {
		dis.offsets[offset].IsCallTarget = true
	}
	if jumpTarget {
		dis.offsets[offset].JumpFrom = append(dis.offsets[offset].JumpFrom, *nes.PC)
		dis.jumpTargets[target] = struct{}{}
	}

	if dis.offsets[offset].IsProcessed {
		return // already disassembled
	}
	dis.targetsToParse = append(dis.targetsToParse, target)
}

// processData sets all data bytes for offsets that have not being identified as code.
func (dis *Disasm) processData() {
	for i, offset := range dis.offsets {
		if offset.Output != "" {
			continue
		}

		address := uint16(i + codeBaseAddress)
		b := dis.sys.ReadMemory(address)
		dis.offsets[i].Data = b
	}
}
