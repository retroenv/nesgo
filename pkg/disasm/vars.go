package disasm

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/nesgo/pkg/disasm/param"
)

const variableNaming = "VAR_%04X"

type variable struct {
	reads  bool
	writes bool

	usageAt []uint16 // list of all addresses that use this offset
}

func (dis *Disasm) addVariableReference(offset uint16, opcode cpu.Opcode, address uint16) bool {
	var reads, writes bool
	if opcode.ReadWritesMemory() {
		reads = true
		writes = true
	} else {
		reads = opcode.ReadsMemory()
		writes = opcode.WritesMemory()
	}
	if !reads && !writes {
		return false
	}

	varInfo := dis.variables[address]
	if varInfo == nil {
		varInfo = &variable{}
		dis.variables[address] = varInfo
	}
	varInfo.usageAt = append(varInfo.usageAt, offset)

	if reads {
		varInfo.reads = true
	}
	if writes {
		varInfo.writes = true
	}

	return true
}

// processJumpTargets processes all variables and updates the instructions that use them
// with a generated alias name.
func (dis *Disasm) processVariables() error {
	for address, varInfo := range dis.variables {
		if len(varInfo.usageAt) == 1 {
			if !varInfo.reads || !varInfo.writes {
				continue // ignore only once usages or ones that are now read and write
			}
		}

		dis.usedVariables[address] = struct{}{}
		name := fmt.Sprintf(variableNaming, address)

		for _, usedAddress := range varInfo.usageAt {
			offset := dis.addressToOffset(usedAddress)
			offsetInfo := &dis.offsets[offset]

			converted, err := param.String(dis.converter, offsetInfo.opcode.Addressing, name)
			if err != nil {
				return err
			}

			switch offsetInfo.opcode.Addressing {
			case ZeroPageAddressing, ZeroPageXAddressing, ZeroPageYAddressing:
				offsetInfo.Code = fmt.Sprintf("%s z:%s", offsetInfo.opcode.Instruction.Name, converted)
			case AbsoluteAddressing, AbsoluteXAddressing, AbsoluteYAddressing:
				offsetInfo.Code = fmt.Sprintf("%s a:%s", offsetInfo.opcode.Instruction.Name, converted)
			case IndirectAddressing, IndirectXAddressing, IndirectYAddressing:
				offsetInfo.Code = fmt.Sprintf("%s %s", offsetInfo.opcode.Instruction.Name, converted)
			}
		}
	}
	return nil
}
