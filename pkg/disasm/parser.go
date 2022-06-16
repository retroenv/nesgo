package disasm

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cpu"
	. "github.com/retroenv/nesgo/pkg/nes"
)

// followExecutionFlow parses opcodes and follows the execution flow
// to parse all code.
func (dis *Disasm) followExecutionFlow() error {
	sys := dis.sys
	var err error
	var params []interface{}

	for len(dis.targets) > 0 {
		dis.popTarget()
		if *PC == 0 {
			break
		}

		opcode := DecodePCInstruction(sys)
		offset := *PC - codeBaseAddress
		dis.offsets[offset].opcode = opcode
		ins := opcode.Instruction

		var paramsAsString string
		var opcodes []byte
		nextTarget := sys.PC + 1

		if ins.NoParamFunc == nil {
			params, opcodes, _ = ReadOpParams(sys, opcode.Addressing)
			dis.offsets[offset].params = params
			paramsAsString, err = paramStrings(dis.converter, opcode, params...)
			if err != nil {
				return err
			}

			nextTarget += uint16(len(opcodes))
			if _, ok := cpu.BranchingInstructions[ins.Name]; ok {
				addr := params[0].(Absolute)
				dis.addTarget(uint16(addr), ins, true)
			}
		}

		if _, ok := cpu.NotExecutingFollowingOpcodeInstructions[ins.Name]; !ok {
			dis.addTarget(nextTarget, ins, false)
		}

		dis.completeInstructionProcessing(offset, ins, opcodes, paramsAsString)
	}
	return nil
}

// processJumpTargets processes all jump targets and updates the callers with
// the generated jump target label name.
func (dis *Disasm) processJumpTargets() {
	for target := range dis.jumpTargets {
		offset := target - codeBaseAddress
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
			offset = caller - codeBaseAddress
			dis.offsets[offset].Output = dis.offsets[offset].opcode.Instruction.Name
			dis.offsets[offset].JumpingTo = name
		}
	}
}

// addTarget adds a target to the list to be processed if the address
// has not been processed yet.
func (dis *Disasm) addTarget(target uint16, currentInstruction *cpu.Instruction, jumpTarget bool) {
	offset := target - codeBaseAddress

	if currentInstruction != nil && currentInstruction.Name == "jsr" {
		dis.offsets[offset].IsCallTarget = true
	}
	if jumpTarget {
		dis.offsets[offset].JumpFrom = append(dis.offsets[offset].JumpFrom, *PC)
		dis.jumpTargets[target] = struct{}{}
	}

	if dis.offsets[offset].IsProcessed {
		return // already disassembled
	}
	dis.targets = append(dis.targets, target)
}

func (dis *Disasm) completeInstructionProcessing(offset uint16, ins *cpu.Instruction, opcodes []byte, params string) {
	if params == "" {
		dis.offsets[offset].Output = ins.Name
	} else {
		dis.offsets[offset].Output = fmt.Sprintf("%s %s", ins.Name, params)
	}
	for i := 0; i < len(opcodes)+1; i++ {
		dis.offsets[offset+uint16(i)].IsProcessed = true
	}
}
