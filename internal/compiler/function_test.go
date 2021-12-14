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
  rts
.endproc

.proc testParam
  dey
  rts
.endproc
`

var functionTestCases = []testCase{
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
