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

	for len(dis.targets) > 0 {
		dis.popTarget()
		if *PC == 0 {
			break
		}

		opcode := DecodePCInstruction(sys)

		var params string
		opcodeLength := uint16(1)

		if opcode.Instruction.ParamFunc != nil { // instruction has parameters
			opcodeLength, params, err = dis.processParamFunction(opcode)
			if err != nil {
				return err
			}
		}
		nextTarget := sys.PC + 1 + opcodeLength

		if _, ok := cpu.NotExecutingFollowingOpcodeInstructions[opcode.Instruction.Name]; !ok {
			dis.addTarget(nextTarget, opcode.Instruction, false)
		}

		dis.completeInstructionProcessing(opcodeLength, opcode, params)
	}
	return nil
}

func (dis *Disasm) processParamFunction(opcode cpu.Opcode) (uint16, string, error) {
	params, opcodes, _ := ReadOpParams(dis.sys, opcode.Addressing)

	offset := *PC - codeBaseAddress
	dis.offsets[offset].params = params

	paramAsString, err := paramString(dis.converter, opcode, params...)
	if err != nil {
		return 0, "", err
	}

	if _, ok := cpu.BranchingInstructions[opcode.Instruction.Name]; ok {
		addr := params[0].(Absolute)
		dis.addTarget(uint16(addr), opcode.Instruction, true)
	}
	return uint16(len(opcodes) + 1), paramAsString, nil
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
	if target == 0 {
		return
	}

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

func (dis *Disasm) completeInstructionProcessing(opcodeLength uint16, opcode cpu.Opcode, params string) {
	offset := *PC - codeBaseAddress
	dis.offsets[offset].opcode = opcode

	if params == "" {
		dis.offsets[offset].Output = opcode.Instruction.Name
	} else {
		dis.offsets[offset].Output = fmt.Sprintf("%s %s", opcode.Instruction.Name, params)
	}

	for i := uint16(0); i < opcodeLength; i++ {
		dis.offsets[offset+i].IsProcessed = true
	}
}
