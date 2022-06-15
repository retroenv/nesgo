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

// result contains the processing result that represents a single byte
// in the program.
type result struct {
	opcode cpu.Opcode
	params []interface{}

	IsProcessed  bool
	IsCallTarget bool // opcode is target of a jsr call, indicating a subroutine
	JumpFrom     []uint16

	Label     string // name of label or subroutine if identified as a jump target
	Output    string
	JumpingTo string // label to jump to if instruction branches
}

// Disasm implements a NES disassembler.
type Disasm struct {
	sys        *system.System
	converter  ParamConverter
	fileWriter fileWriter

	// jumpTargets is a set of all addresses that
	jumpTargets map[uint16]struct{}
	results     []result

	targets []uint16
}

// New creates a new NES disassembler that creates output compatible with the
// chose assembler.
func New(cart *cartridge.Cartridge, assembler string) (*Disasm, error) {
	opts := NewOptions(WithCartridge(cart))
	dis := &Disasm{
		sys:         InitializeSystem(opts),
		results:     make([]result, len(cart.PRG)),
		jumpTargets: map[uint16]struct{}{},
	}
	if err := dis.initializeCompatibleMode(assembler); err != nil {
		return nil, fmt.Errorf("initializing compatible mode: %w", err)
	}

	dis.initializeIrqHandlers()
	return dis, nil
}

// initializeCompatibleMode sets the chosen assembler specific instances
// to be used to output compatible code.
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

func (dis *Disasm) initializeIrqHandlers() {
	resetHandler := dis.sys.ReadMemory16(0xFFFC)
	dis.targets = []uint16{resetHandler}
	offset := resetHandler - codeBaseAddress
	dis.results[offset].Label = "resetHandler"
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

// popTarget pops the next target to disassemble and sets it into PC.
func (dis *Disasm) popTarget() {
	dis.sys.PC = dis.targets[0]
	dis.targets = dis.targets[1:]
}

func (dis *Disasm) convertToProgram() *program.Program {
	app := program.New(len(dis.results))

	for i := 0; i < len(dis.results); i++ {
		res := dis.results[i]
		if !res.IsProcessed || res.Output == "" {
			continue
		}

		app.Offsets[i] = program.Offset{
			IsCallTarget: res.IsCallTarget,
			Label:        res.Label,
			Output:       res.Output,
		}

		if res.JumpingTo != "" {
			app.Offsets[i].Output = fmt.Sprintf("%s %s", res.Output, res.JumpingTo)
		}
	}

	return app
}
