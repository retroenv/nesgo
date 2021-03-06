package disasm

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/disasm/program"
)

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
				insNext.SetType(program.DataOffset)
				noLabelOffsets++
				continue
			}
			break
		}

		ins.OpcodeBytes = data[i : i+noLabelOffsets]
		ins.ClearType(program.CodeOffset)
		ins.SetType(program.DataOffset)
		i += noLabelOffsets - 1
	}
}

// processData sets all data bytes for offsets that have not being identified as code.
func (dis *Disasm) processData() {
	for i, offset := range dis.offsets {
		if offset.IsType(program.CodeOffset) || offset.IsType(program.DataOffset) {
			continue
		}

		address := uint16(i + CodeBaseAddress)
		b := dis.sys.Bus.Memory.Read(address)
		dis.offsets[i].OpcodeBytes = []byte{b}
	}
}
