package tests

import (
	"testing"
)

var ifBranchingEmpty = []byte(`
if Bne() {
}
`)

var ifBranchingGoto = []byte(`
a:
if Bne() {
  goto a
}
`)
var ifBranchingGotoIr = `
label, a
inst, bne, a
`

var ifBranchingBreak = []byte(`
a:
if Bne() {
  break
}
`)
var ifBranchingBreakIr = `
label, a
inst, bne
inst, break
`

var ifBranchingInstruction = []byte(`
if Bne() {
  Dex()
}
`)
var ifBranchingInstructionIr = `
inst, bne, if_bne
inst, jmp, if_not_bne
label, if_bne
inst, dex
label, if_not_bne
`

var ifNotBranchingInstruction = []byte(`
if !Bne() {
  Dex()
}
`)
var ifNotBranchingInstructionIr = `
inst, bne, if_not_bne
inst, dex
label, if_not_bne
`

var ifTestCases = []testCase{
	{
		"if with not branching and instruction",
		ifNotBranchingInstruction,
		ifNotBranchingInstructionIr,
		"",
	},
	{
		"if with branching and instruction",
		ifBranchingInstruction,
		ifBranchingInstructionIr,
		"",
	},
	{
		"if with branching and break",
		ifBranchingBreak,
		ifBranchingBreakIr,
		"",
	},
	{
		"if with branching and goto",
		ifBranchingGoto,
		ifBranchingGotoIr,
		"",
	},
	{
		"if with branching and empty block",
		ifBranchingEmpty,
		"",
		"if statement with branch can not have an empty block, goto or break expected",
	},
}

func TestIf(t *testing.T) {
	for _, test := range ifTestCases {
		runTest(t, true, test.input, test.expectedIr, test.expectedError, test.name)
	}
}
