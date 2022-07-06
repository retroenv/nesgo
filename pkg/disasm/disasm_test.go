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

var testCodeDefault = []byte{
	0x78,             // sei
	0x4C, 0x04, 0x80, // jmp + 3
	0xAD, 0x30, 0x80, // lda a:$8030
	0xBD, 0x20, 0x80, // lda a:$8020,X
	0x1a,             // nop
	0xdc, 0xae, 0x8b, // nop $8BAE,X
	// TODO jump into instruction
	0x40, // rti
}

var expectedDefault = `Reset:
  sei                            ; $8000 78
  jmp _label_8004                ; $8001 4C 04 80

_label_8004:
  lda a:_data_8030               ; $8004 AD 30 80
  lda a:_data_8020_indexed,X     ; $8007 BD 20 80
.byte $1a                        ; $800A unofficial nop instruction: nop
.byte $dc, $ae, $8b              ; $800B unofficial nop instruction: nop $8BAE,X
  rti                            ; $800E 40

.byte $00, $00, $00, $00, $00, $00, $00, $00, $00, $00, $00, $00, $00, $00, $00, $00
.byte $00

_data_8020_indexed:
.byte $12, $00, $00, $00, $00, $00, $00, $00, $00, $00, $00, $00, $00, $00, $00, $00

_data_8030:
.byte $34
`

var testCodeNoHexNoAddress = []byte{
	0x78,             // sei
	0x4C, 0x05, 0x80, // jmp + 3
	0x1a, // nop
	0x40, // rti
}

var expectedNoOffsetNoHex = `Reset:
  sei
  jmp _label_8005

.byte $1a

_label_8005:
  rti
`

func testProgram(t *testing.T, options *disasmoptions.Options, cart *cartridge.Cartridge, code []byte) *Disasm {
	t.Helper()

	// point reset handler to offset 0 of PRG buffer, aka 0x8000 address
	cart.PRG[0x7FFD] = 0x80

	copy(cart.PRG, code)

	disasm, err := New(cart, options)
	assert.NoError(t, err)

	return disasm
}

func TestDisasmDefault(t *testing.T) {
	options := disasmoptions.New()
	options.CodeOnly = true
	options.Assembler = ca65.Name

	cart := cartridge.New()
	cart.PRG[0x0020] = 0x12
	cart.PRG[0x0030] = 0x34

	disasm := testProgram(t, &options, cart, testCodeDefault)

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
	cart := cartridge.New()
	disasm := testProgram(t, &options, cart, testCodeNoHexNoAddress)

	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)

	err := disasm.Process(writer)
	assert.NoError(t, err)

	assert.NoError(t, writer.Flush())

	buf := buffer.String()
	assert.Equal(t, expectedNoOffsetNoHex, buf)
}
