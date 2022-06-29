package disasm

import (
	"fmt"
	"strings"

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
		dis.offsets[offset].OpcodeBytes = make([]byte, 1, 3)
		dis.offsets[offset].OpcodeBytes[0] = dis.sys.ReadMemory(*nes.PC)

		opcode, err := nes.DecodePCInstruction(dis.sys)
		if err != nil {
			// consider an unknown instruction as start of data
			dis.offsets[offset].Type |= program.DataOffset
			continue
		}

		var params string
		instruction := opcode.Instruction

		// process instructions with parameters, ignore special case of unofficial nop
		// that also has an implied addressing without parameters.
		if instruction.ParamFunc != nil && opcode.Addressing != ImpliedAddressing {
			var opcodes []byte
			opcodes, params, err = dis.processParamInstruction(opcode)
			if err != nil {
				return err
			}
			dis.offsets[offset].OpcodeBytes = append(dis.offsets[offset].OpcodeBytes, opcodes...)
		}

		opcodeLength := uint16(len(dis.offsets[offset].OpcodeBytes))
		nextTarget := dis.sys.PC + opcodeLength

		if _, ok := cpu.NotExecutingFollowingOpcodeInstructions[instruction.Name]; !ok {
			dis.addTarget(nextTarget, instruction, false)
		}

		dis.offsets[offset].opcode = opcode

		if params == "" {
			dis.offsets[offset].Code = instruction.Name
		} else {
			dis.offsets[offset].Code = fmt.Sprintf("%s %s", instruction.Name, params)
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

// processJumpTargets processes all jump targets and updates the callers with
// the generated jump target label name.
func (dis *Disasm) processJumpTargets() {
	for target := range dis.jumpTargets {
		offset := dis.addressToOffset(target)
		name := dis.offsets[offset].Label
		if name == "" {
			if dis.offsets[offset].Type&program.CallTarget != 0 {
				name = fmt.Sprintf("_func_%04x", target)
			} else {
				name = fmt.Sprintf("_label_%04x", target)
			}
			dis.offsets[offset].Label = name
		}

		// if the offset is marked as code but does not have opcode bytes, the jumping target
		// is inside the second or third byte of an instruction.
		if dis.offsets[offset].Type&program.CodeOffset != 0 && len(dis.offsets[offset].OpcodeBytes) == 0 {
			dis.handleJumpIntoInstruction(offset)
		}

		for _, caller := range dis.offsets[offset].JumpFrom {
			offset = dis.addressToOffset(caller)
			dis.offsets[offset].Code = dis.offsets[offset].opcode.Instruction.Name
			dis.offsets[offset].JumpingTo = name
		}
	}
}

// handleJumpIntoInstruction converts an instruction that has a jump destination label inside
// its second or third opcode bytes into data.
func (dis *Disasm) handleJumpIntoInstruction(offset uint16) {
	// look backwards for instruction start
	instructionStart := offset - 1
	for ; len(dis.offsets[instructionStart].OpcodeBytes) == 0; instructionStart-- {
	}

	ins := &dis.offsets[instructionStart]
	ins.Comment = fmt.Sprintf("branch into instruction detected: %s", ins.Code)
	ins.Code = ""
	data := ins.OpcodeBytes
	dis.changeOffsetRangeToData(data, instructionStart)
}

// handleUnofficialNop translates unofficial nop codes into data bytes as the instruction
// has multiple opcodes for the same addressing mode which will result in a different
// bytes being assembled.
func (dis *Disasm) handleUnofficialNop(offset uint16) {
	ins := &dis.offsets[offset]
	ins.Comment = fmt.Sprintf("unofficial nop instruction: %s", ins.Code)
	ins.Code = ""
	data := ins.OpcodeBytes
	dis.changeOffsetRangeToData(data, offset)
}

// changeOffsetRangeToData sets a range of code offsets to data types.
// It combines all data bytes that are not split by a label.
func (dis *Disasm) changeOffsetRangeToData(data []byte, offset uint16) {
	for i := 0; i < len(data); i++ {
		ins := &dis.offsets[offset+uint16(i)]

		noLabelOffsets := 1
		for j := i + 1; j < len(data); j++ {
			insNext := &dis.offsets[offset+uint16(j)]
			if insNext.Label == "" {
				insNext.OpcodeBytes = nil
				insNext.Type |= program.DataOffset
				noLabelOffsets++
				continue
			}
			break
		}

		ins.OpcodeBytes = data[i : i+noLabelOffsets]
		ins.Type ^= program.CodeOffset
		ins.Type |= program.DataOffset
		i += noLabelOffsets - 1
	}
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

// processData sets all data bytes for offsets that have not being identified as code.
func (dis *Disasm) processData() {
	for i, offset := range dis.offsets {
		if offset.Type&program.CodeOffset != 0 || offset.Type&program.DataOffset != 0 {
			continue
		}

		address := uint16(i + codeBaseAddress)
		b := dis.sys.ReadMemory(address)
		dis.offsets[i].OpcodeBytes = []byte{b}
	}
}
