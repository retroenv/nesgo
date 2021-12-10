package tests

import (
	"testing"
)

var forEndless = []byte(`
for {
}
`)
var forEndlessIr = `
label, loop
inst, jmp, loop
`

var forWhileEmpty = []byte(`
for Bne() {
}
`)
var forWhileEmptyIr = `
label, bne_loop
inst, bne, bne_loop
`

var forWhileInstruction = []byte(`
for Bne() {
  Dex()
}
`)
var forWhileInstructionIr = `
label, bne_loop
inst, dex
inst, bne, bne_loop
`

var forWhileInstructionNoBranching = []byte(`
for Dex() {
}
`)

var forEndlessInstruction = []byte(`
for {
  Dex()
}
`)
var forEndlessInstructionIr = `
label, loop
inst, dex
inst, jmp, loop
`

var forClauseTerminator = []byte(`
for ;; {
  Dex()
}
`)
var forClauseTerminatorIr = `
label, loop
inst, dex
inst, jmp, loop
`

var forClauseInit = []byte(`
for *X = 0;; {
  Dex()
}
`)
var forClauseInitIr = `
op, =, X,0
label, loop
inst, dex
inst, jmp, loop
`

var forClauseInitCond = []byte(`
for *X = 0; *X < 8; {
  Dex()
}
`)
var forClauseInitCondIr = `
op, =, X,0
label, bcc_loop
inst, dex
inst, cpx, 8
inst, bcc, bcc_loop
`

var forClauseInitCondPost = []byte(`
for *X = 0; *X < 8; *X++ {
  Dex()
}
`)
var forClauseInitCondPostIr = `
op, =, X,0
label, bcc_loop
inst, dex
inst, inx
inst, cpx, 8
inst, bcc, bcc_loop
`

var forClausePost = []byte(`
for ; ; *X++ {
  Dex()
}
`)
var forClausePostIr = `
label, loop
inst, dex
inst, inx
inst, jmp, loop
`

var forClausePostInstruction = []byte(`
for ; ; Dex() {
}
`)
var forClausePostInstructionIr = `
label, loop
inst, dex
inst, jmp, loop
`

var forClauseCondPost = []byte(`
for ; *X < 8; *X++ {
  Dex()
}
`)
var forClauseCondPostIr = `
label, bcc_loop
inst, dex
inst, inx
inst, cpx, 8
inst, bcc, bcc_loop
`

var forClauseInitPost = []byte(`
for *X = 0; ; *X++ {
  Dex()
}
`)
var forClauseInitPostIr = `
op, =, X,0
label, loop
inst, dex
inst, inx
inst, jmp, loop
`

var forWhileBreak = []byte(`
for {
  if Bne() {
    break
  }
}

`)
var forWhileBreakIr = `
label, loop
inst, bne, loop_end
inst, jmp, loop
label, loop_end
`

var forWhileMultipleBreaks = []byte(`
for {
  Dex()
  if Bne() {
    break
  }
  Dey()
  if Bcc() {
    break
  }
}

`)
var forWhileMultipleBreaksIr = `
label, loop
inst, dex
inst, bne, loop_end
inst, dey
inst, bcc, loop_end
inst, jmp, loop
label, loop_end
`

var forWhileContinue = []byte(`
for {
  if Bne() {
    continue
  }
  Dex()
}

`)
var forWhileContinueIr = `
label, loop
inst, bne, loop_post
inst, dex
label, loop_post
inst, jmp, loop
`

var forTestCases = []testCase{
	{
		"while loop with continue",
		forWhileContinue,
		forWhileContinueIr,
		"",
	},
	{
		"while loop with branch, instructions and multiple breaks",
		forWhileMultipleBreaks,
		forWhileMultipleBreaksIr,
		"",
	},
	{
		"while loop with branch and break",
		forWhileBreak,
		forWhileBreakIr,
		"",
	},
	{
		"for clause with post expression instruction",
		forClausePostInstruction,
		forClausePostInstructionIr,
		"",
	},
	{
		"for clause with init and post expression",
		forClauseInitPost,
		forClauseInitPostIr,
		"",
	},
	{
		"for clause with condition and post expression",
		forClauseCondPost,
		forClauseCondPostIr,
		"",
	},
	{
		"for clause with post expression",
		forClausePost,
		forClausePostIr,
		"",
	},
	{
		"for clause with initializer",
		forClauseInit,
		forClauseInitIr,
		"",
	},
	{
		"for clause with initializer and condition",
		forClauseInitCond,
		forClauseInitCondIr,
		"",
	},
	{
		"for clause with initializer, condition and post statement",
		forClauseInitCondPost,
		forClauseInitCondPostIr,
		"",
	},
	{
		"for clause with 2 terminators",
		forClauseTerminator,
		forClauseTerminatorIr,
		"",
	},
	{
		"invalid while loop without branching instruction",
		forWhileInstructionNoBranching,
		"",
		"expected a branching instruction as for statement condition but found 'dex'",
	},
	{
		"endless loop",
		forEndless,
		forEndlessIr,
		"",
	},
	{
		"empty while loop with branching instruction",
		forWhileEmpty,
		forWhileEmptyIr,
		"",
	},
	{
		"while loop with instruction",
		forWhileInstruction,
		forWhileInstructionIr,
		"",
	},
	{
		"endless loop with instruction",
		forEndlessInstruction,
		forEndlessInstructionIr,
		"",
	},
}

func TestFor(t *testing.T) {
	for _, test := range forTestCases {
		runTest(t, true, test)
	}
}
