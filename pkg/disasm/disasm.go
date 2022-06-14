// Package disasm provides an NES program disassembler.
package disasm

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/cpu"
	. "github.com/retroenv/nesgo/pkg/nes"
	"github.com/retroenv/nesgo/pkg/system"
)

// Result contains the processing result that represents a single byte
// in the program.
type Result struct {
	Processed  bool
	JumpTarget bool
	Label      string
	Output     string
}

// Disasm implements a NES disassembler.
type Disasm struct {
	sys       *system.System
	converter ParamConverter

	results []Result

	targets []uint16
}

// New creates a new NES disassembler.
func New(cart *cartridge.Cartridge) *Disasm {
	opts := NewOptions(WithCartridge(cart))
	dis := &Disasm{
		sys:       InitializeSystem(opts),
		converter: paramConverterCa65{},
		results:   make([]Result, len(cart.PRG)),
	}
	resetHandler := dis.sys.ReadMemory16(0xFFFC)
	dis.targets = []uint16{resetHandler}
	dis.results[resetHandler-0x8000].Label = "resetHandler:"
	return dis
}

// Process desassembles the cartridge.
func (dis *Disasm) Process() error {
	sys := dis.sys
	var err error
	var params []interface{}

	for len(dis.targets) > 0 {
		dis.popTarget()
		if *PC == 0 {
			break
		}
		offset := *PC - 0x8000

		opcode := DecodePCInstruction(sys)
		ins := opcode.Instruction

		var paramsAsString string
		var opcodes []byte
		nextTarget := sys.PC + 1

		if ins.NoParamFunc == nil {
			params, opcodes, _ = ReadOpParams(sys, opcode.Addressing)
			paramsAsString, err = ParamStrings(dis.converter, opcode, params...)
			if err != nil {
				return err
			}

			nextTarget += uint16(len(opcodes))
			if _, ok := cpu.BranchingInstructions[ins.Name]; ok {
				addr := params[0].(Absolute)
				dis.addTarget(uint16(addr), true)
			}
		}

		if ins.Name != "jmp" && ins.Name != "rts" && ins.Name != "rti" {
			dis.addTarget(nextTarget, false)
		}

		dis.completeInstructionProcessing(offset, ins, opcodes, paramsAsString)
	}

	dis.print()
	return nil
}

// addTarget adds a target to the list to be processed if the address
// has not been processed yet.
func (dis *Disasm) addTarget(target uint16, jumpTarget bool) {
	offset := target - 0x8000
	if dis.results[offset].Processed {
		if jumpTarget {
			dis.results[offset].JumpTarget = true
		}
		return // already disassembled
	}
	if jumpTarget {
		dis.results[offset].JumpTarget = true
	}
	dis.targets = append(dis.targets, target)
}

// popTarget pops the next target to diassemble and sets it into PC.
func (dis *Disasm) popTarget() {
	dis.sys.PC = dis.targets[0]
	dis.targets = dis.targets[1:]
}

func (dis *Disasm) completeInstructionProcessing(offset uint16, ins *cpu.Instruction, opcodes []byte, params string) {
	if params == "" {
		dis.results[offset].Output = ins.Name
	} else {
		dis.results[offset].Output = fmt.Sprintf("%s %s", ins.Name, params)
	}
	for i := 0; i < len(opcodes)+1; i++ {
		dis.results[offset+uint16(i)].Processed = true
	}
}

func (dis *Disasm) print() {
	for i := 0; i < len(dis.results); i++ {
		res := dis.results[i]
		if !res.Processed || res.Output == "" {
			continue
		}
		if res.JumpTarget {
			fmt.Printf("_label_%04x:\n", i+0x8000)
		}
		fmt.Printf("  %s\n", res.Output)
	}
}
