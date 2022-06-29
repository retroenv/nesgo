package disasm

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/disasm/ca65"
	"github.com/retroenv/nesgo/pkg/disasm/disasmoptions"
)

var testCode = []byte{
	0x78,             // $8000 sei
	0x1a,             // $8001 nop
	0xdc, 0xae, 0x8b, // $8001 nop $8BAE,X
	// TODO jump into instruction
	0x40, // $8004 rti
}

var expectedDefault = `Reset:
  sei                            ; $8000 78
.byte $1a                        ; $8001 unofficial nop instruction: nop
.byte $dc, $ae, $8b              ; $8002 unofficial nop instruction: nop $8BAE,X
  rti                            ; $8005 40
`

var expectedNoOffsetNoHex = `Reset:
  sei
.byte $1a                        ; unofficial nop instruction: nop
.byte $dc, $ae, $8b              ; unofficial nop instruction: nop $8BAE,X
  rti
`

func testProgram(t *testing.T, options *disasmoptions.Options) *Disasm {
	t.Helper()

	cart := cartridge.New()

	// point reset handler to offset 0 of PRG buffer, aka 0x8000 address
	cart.PRG[0x7FFD] = 0x80

	copy(cart.PRG, testCode)

	disasm, err := New(cart, options)
	assert.NoError(t, err)

	return disasm
}

func TestDisasmDefault(t *testing.T) {
	options := disasmoptions.New()
	options.CodeOnly = true
	options.Assembler = ca65.Name
	disasm := testProgram(t, &options)

	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)

	err := disasm.Process(writer)
	assert.NoError(t, err)

	assert.NoError(t, writer.Flush())

	buf := buffer.String()
	assert.Equal(t, expectedDefault, buf)
}

func TestDisasmNoHexNoAddress(t *testing.T) {
	options := disasmoptions.New()
	options.CodeOnly = true
	options.Assembler = ca65.Name
	options.OffsetComments = false
	options.HexComments = false
	disasm := testProgram(t, &options)

	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)

	err := disasm.Process(writer)
	assert.NoError(t, err)

	assert.NoError(t, writer.Flush())

	buf := buffer.String()
	assert.Equal(t, expectedNoOffsetNoHex, buf)
}
