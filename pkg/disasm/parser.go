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
		dis.results[offset].opcode = opcode
		ins := opcode.Instruction

		var paramsAsString string
		var opcodes []byte
		nextTarget := sys.PC + 1

		if ins.NoParamFunc == nil {
			params, opcodes, _ = ReadOpParams(sys, opcode.Addressing)
			dis.results[offset].params = params
			paramsAsString, err = ParamStrings(dis.converter, opcode, params...)
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
		name := dis.results[offset].Label
		if name == "" {
			if dis.results[offset].IsCallTarget {
				name = fmt.Sprintf("_func_%04x", target)
			} else {
				name = fmt.Sprintf("_label_%04x", target)
			}
			dis.results[offset].Label = name
		}

		for _, caller := range dis.results[offset].JumpFrom {
			offset = caller - codeBaseAddress
			dis.results[offset].Output = dis.results[offset].opcode.Instruction.Name
			dis.results[offset].JumpingTo = name
		}
	}
}

// addTarget adds a target to the list to be processed if the address
// has not been processed yet.
func (dis *Disasm) addTarget(target uint16, currentInstruction *cpu.Instruction, jumpTarget bool) {
	offset := target - codeBaseAddress

	if currentInstruction != nil && currentInstruction.Name == "jsr" {
		dis.results[offset].IsCallTarget = true
	}
	if jumpTarget {
		dis.results[offset].JumpFrom = append(dis.results[offset].JumpFrom, *PC)
		dis.jumpTargets[target] = struct{}{}
	}

	if dis.results[offset].IsProcessed {
		return // already disassembled
	}
	dis.targets = append(dis.targets, target)
}

func (dis *Disasm) completeInstructionProcessing(offset uint16, ins *cpu.Instruction, opcodes []byte, params string) {
	if params == "" {
		dis.results[offset].Output = ins.Name
	} else {
		dis.results[offset].Output = fmt.Sprintf("%s %s", ins.Name, params)
	}
	for i := 0; i < len(opcodes)+1; i++ {
		dis.results[offset+uint16(i)].IsProcessed = true
	}
}
