package disasm

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/nesgo/pkg/disasm/param"
)

const (
	dataNaming            = "_data_%04x"
	dataNamingIndexed     = "_data_%04x_indexed"
	variableNaming        = "_var_%04x"
	variableNamingIndexed = "_var_%04x_indexed"
)

type variable struct {
	reads  bool
	writes bool

	name         string
	indexedUsage bool     // access with X/Y registers indicates table
	usageAt      []uint16 // list of all addresses that use this offset
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

	if address == 0xc135 {
		fmt.Println("asdf")
	}

	switch opcode.Addressing {
	case ZeroPageXAddressing, ZeroPageYAddressing,
		AbsoluteXAddressing, AbsoluteYAddressing,
		IndirectXAddressing, IndirectYAddressing:
		varInfo.indexedUsage = true
	}

	return true
}

// processJumpTargets processes all variables and updates the instructions that use them
// with a generated alias name.
func (dis *Disasm) processVariables() error {
	for address, varInfo := range dis.variables {
		if len(varInfo.usageAt) == 1 && !varInfo.indexedUsage && address < CodeBaseAddress {
			if !varInfo.reads || !varInfo.writes {
				continue // ignore only once usages or ones that are not read and write
			}
		}

		var offsetInfo *offset
		if address >= CodeBaseAddress {
			offset := dis.addressToOffset(address)
			offsetInfo = &dis.offsets[offset]
		} else {
			dis.usedVariables[address] = struct{}{}
		}

		varInfo.name = dataName(offsetInfo, varInfo.indexedUsage, address)

		for _, usedAddress := range varInfo.usageAt {
			offset := dis.addressToOffset(usedAddress)
			offsetInfo := &dis.offsets[offset]

			converted, err := param.String(dis.converter, offsetInfo.opcode.Addressing, varInfo.name)
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

func dataName(offsetInfo *offset, indexedUsage bool, address uint16) string {
	if offsetInfo != nil && offsetInfo.Label != "" {
		return offsetInfo.Label
	}

	var name string
	prgAccess := offsetInfo != nil

	switch {
	case prgAccess && indexedUsage:
		name = fmt.Sprintf(dataNamingIndexed, address)
	case prgAccess && !indexedUsage:
		name = fmt.Sprintf(dataNaming, address)
	case !prgAccess && indexedUsage:
		name = fmt.Sprintf(variableNamingIndexed, address)
	default:
		name = fmt.Sprintf(variableNaming, address)
	}

	if offsetInfo != nil {
		offsetInfo.Label = name
	}
	return name
}
