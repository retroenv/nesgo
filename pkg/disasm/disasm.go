// Package disasm provides an NES program disassembler.
package disasm

import (
	"fmt"
	"io"
	"strings"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/nesgo/pkg/disasm/ca65"
	"github.com/retroenv/nesgo/pkg/disasm/program"
	. "github.com/retroenv/nesgo/pkg/nes"
	"github.com/retroenv/nesgo/pkg/system"
)

const codeBaseAddress = 0x8000

type fileWriter interface {
	Write(app *program.Program, writer io.Writer) error
}

// offset defines the content of an offset in a program that can represent data or code.
type offset struct {
	opcode cpu.Opcode    // opcode that the byte at this offset represents
	params []interface{} // internal representation of the instruction parameters

	IsProcessed  bool     // flag whether current offset and following opcode bytes have been processed
	IsCallTarget bool     // opcode is target of a jsr call, indicating a subroutine
	JumpFrom     []uint16 // list of all addresses that jump to this offset

	Label     string // name of label or subroutine if identified as a jump target
	Output    string // asm output of this instruction
	JumpingTo string // label to jump to if instruction branches
}

// Disasm implements a NES disassembler.
type Disasm struct {
	sys        *system.System
	converter  paramConverter
	fileWriter fileWriter
	cart       *cartridge.Cartridge
	constants  map[uint16]string

	jumpTargets map[uint16]struct{} // jumpTargets is a set of all addresses that branched to
	offsets     []offset

	targets  []uint16
	handlers program.Handlers
}

// New creates a new NES disassembler that creates output compatible with the chosen assembler.
func New(cart *cartridge.Cartridge, assembler string) (*Disasm, error) {
	opts := NewOptions(WithCartridge(cart))
	dis := &Disasm{
		sys:         InitializeSystem(opts),
		cart:        cart,
		constants:   buildConstMap(),
		offsets:     make([]offset, len(cart.PRG)),
		jumpTargets: map[uint16]struct{}{},
		handlers: program.Handlers{
			NMI:   "0",
			Reset: "Reset",
			IRQ:   "0",
		},
	}

	if err := dis.initializeCompatibleMode(assembler); err != nil {
		return nil, fmt.Errorf("initializing compatible mode: %w", err)
	}

	dis.initializeIrqHandlers()
	return dis, nil
}

// Process disassembles the cartridge.
func (dis *Disasm) Process(writer io.Writer) error {
	if err := dis.followExecutionFlow(); err != nil {
		return err
	}

	dis.processJumpTargets()

	app := dis.convertToProgram()
	return dis.fileWriter.Write(app, writer)
}

// initializeCompatibleMode sets the chosen assembler specific instances to be used to output
// compatible code.
func (dis *Disasm) initializeCompatibleMode(assembler string) error {
	switch strings.ToLower(assembler) {
	case ca65.Name:
		dis.converter = ca65.ParamConverter{}
		dis.fileWriter = ca65.FileWriter{}

	default:
		return fmt.Errorf("unsupported assembler '%s'", assembler)
	}
	return nil
}

// initializeIrqHandlers reads the 3 handler addresses and adds them to the addresses to be
// followed for execution flow.
func (dis *Disasm) initializeIrqHandlers() {
	nmi := dis.sys.ReadMemory16(0xFFFA)
	if nmi != 0 {
		dis.addTarget(nmi, nil, false)
		offset := nmi - codeBaseAddress
		dis.offsets[offset].Label = "NMI"
		dis.offsets[offset].IsCallTarget = true
		dis.handlers.NMI = "NMI"
	}

	reset := dis.sys.ReadMemory16(0xFFFC)
	dis.addTarget(reset, nil, false)
	offset := reset - codeBaseAddress
	dis.offsets[offset].Label = "Reset"
	dis.offsets[offset].IsCallTarget = true

	irq := dis.sys.ReadMemory16(0xFFFE)
	if irq != 0 {
		dis.addTarget(irq, nil, false)
		offset = irq - codeBaseAddress
		dis.offsets[offset].Label = "IRQ"
		dis.offsets[offset].IsCallTarget = true
		dis.handlers.IRQ = "IRQ"
	}
}

// popTarget pops the next target to disassemble and sets it into the program counter.
func (dis *Disasm) popTarget() {
	dis.sys.PC = dis.targets[0]
	dis.targets = dis.targets[1:]
}

// converts the internal disasm type representation to a program type that will be used by
// the chosen assembler output instance to generate the asm file.
func (dis *Disasm) convertToProgram() *program.Program {
	app := program.New(dis.cart)
	app.Handlers = dis.handlers

	for i := 0; i < len(dis.offsets); i++ {
		res := dis.offsets[i]
		if !res.IsProcessed || res.Output == "" {
			continue
		}

		app.PRG[i] = program.Offset{
			IsCallTarget: res.IsCallTarget,
			Label:        res.Label,
			Output:       res.Output,
		}

		if res.JumpingTo != "" {
			app.PRG[i].Output = fmt.Sprintf("%s %s", res.Output, res.JumpingTo)
		}
	}

	return app
}
