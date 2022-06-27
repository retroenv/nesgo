package compiler

import "testing"

var functionRegisterParam = []byte(`
func test() {
  Dex()
  testParam(0)
}

func testParam(index uint8) {
  Ldy(index)
  Dey()
}
`)
var functionRegisterParamAssembly = `
.proc test
  dex
  ldy #$00
  jsr testParam
  rti
.endproc

.proc testParam
  dey
  rts
.endproc
`

var instructionRegisterParam = []byte(`
func test() {
  Sta(0x200, X)
}
`)
var instructionRegisterParamAssembly = `
.proc test
  sta $0200, X
  rti
.endproc
`

var functionTestCases = []testCase{
	{
		"instruction with register as index param",
		instructionRegisterParam,
		instructionRegisterParamAssembly,
	},
	{
		"function with register as param",
		functionRegisterParam,
		functionRegisterParamAssembly,
	},
}

func TestFunction(t *testing.T) {
	for _, test := range functionTestCases {
		runCompileTest(t, test)
	}
}
