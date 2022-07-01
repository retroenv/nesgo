package disasm

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
)

type variable struct {
	usageAt []uint16 // list of all addresses that use this offset
}

func (dis *Disasm) addVariableReference(offset uint16, address Absolute) {
	varInfo := dis.variables[uint16(address)]
	if varInfo == nil {
		varInfo = &variable{}
		dis.variables[uint16(address)] = varInfo
	}
	varInfo.usageAt = append(varInfo.usageAt, offset)
}

// processJumpTargets processes all variables and updates the instructions that use them
// with a generated alias name.
func (dis *Disasm) processVariables() {
	for _, varInfo := range dis.variables {
		if len(varInfo.usageAt) == 1 {
			continue // ignore once usages
		}

		// TODO implement
	}
}
