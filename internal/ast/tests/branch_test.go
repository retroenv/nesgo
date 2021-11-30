package tests

import (
	"testing"
)

var branchLabel = []byte(`
a:
`)
var branchLabelIr = `
label, a
`

var branchLabelStatement = []byte(`
a: Dex()
`)
var branchLabelStatementIr = `
label, a
inst, dex
`

var branchLabelBranching = []byte(`
a: 
  if Bne() {
    goto a
  }
`)
var branchLabelBranchingIr = `
label, a
inst, bne, a
`

var branchTestCases = []struct {
	name          string
	input         []byte
	expectedIr    string
	expectedError string
}{
	{
		"label with branching",
		branchLabelBranching,
		branchLabelBranchingIr,
		"",
	},
	{
		"simple label",
		branchLabel,
		branchLabelIr,
		"",
	},
	{
		"label with instruction",
		branchLabelStatement,
		branchLabelStatementIr,
		"",
	},
}

func TestBranch(t *testing.T) {
	for _, test := range branchTestCases {
		runTest(t, true, test.input, test.expectedIr, test.expectedError, test.name)
	}
}
