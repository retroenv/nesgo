// Package disasm provides an NES program disassembler.
package disasm

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/cpu"
	. "github.com/retroenv/nesgo/pkg/nes"
	"github.com/retroenv/nesgo/pkg/system"
)

const codeBaseAddress = 0x8000

// result contains the processing result that represents a single byte
// in the program.
type result struct {
	opcode cpu.Opcode
	params []interface{}

	IsProcessed  bool
	IsCallTarget bool
	JumpFrom     []uint16

	Label     string // name of label or subroutine if identified as jump target
	Output    string
	JumpingTo string // label to jump to if instruction branches
}

// Disasm implements a NES disassembler.
type Disasm struct {
	sys       *system.System
	converter ParamConverter

	// jumpTargets is a set of all addresses that
	jumpTargets map[uint16]struct{}
	results     []result

	targets []uint16
}

// New creates a new NES disassembler.
func New(cart *cartridge.Cartridge) *Disasm {
	opts := NewOptions(WithCartridge(cart))
	dis := &Disasm{
		sys:         InitializeSystem(opts),
		converter:   paramConverterCa65{},
		results:     make([]result, len(cart.PRG)),
		jumpTargets: map[uint16]struct{}{},
	}

	resetHandler := dis.sys.ReadMemory16(0xFFFC)
	dis.targets = []uint16{resetHandler}
	offset := resetHandler - codeBaseAddress
	dis.results[offset].Label = "resetHandler"
	return dis
}

// Process disassembles the cartridge.
func (dis *Disasm) Process() error {
	if err := dis.followExecutionFlow(); err != nil {
		return err
	}
	dis.processJumpTargets()

	dis.print()
	return nil
}

// popTarget pops the next target to diassemble and sets it into PC.
func (dis *Disasm) popTarget() {
	dis.sys.PC = dis.targets[0]
	dis.targets = dis.targets[1:]
}

func (dis *Disasm) print() {
	for i := 0; i < len(dis.results); i++ {
		res := dis.results[i]
		if !res.IsProcessed || res.Output == "" {
			continue
		}
		if res.Label != "" {
			if res.IsCallTarget {
				fmt.Println()
			}
			fmt.Printf("%s:\n", res.Label)
		}
		if res.JumpingTo != "" {
			fmt.Printf("  %s %s\n", res.Output, res.JumpingTo)
		} else {
			fmt.Printf("  %s\n", res.Output)
		}
	}
}
