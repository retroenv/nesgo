package tests

import (
	"testing"
)

var instructionNoParam = []byte(`
Lda()
`)
var instructionNoParamIr = `
inst, lda
`

var instructionImmediate = []byte(`
Lda(0x12)
`)
var instructionImmediateIr = `
inst, lda, immediate, 0x12
`

var instructionAbsolute = []byte(`
Lda(0x1234)
`)
var instructionAbsoluteIr = `
inst, lda, absolute, 0x1234
`

var instructionAbsoluteConst = []byte(`
Sta(JOYPAD1)
`)
var instructionAbsoluteConstIr = `
inst, sta, absolute, JOYPAD1
`

var instructionZeroPage = []byte(`
Lda(ZeroPage(0x12))
`)
var instructionZeroPageIr = `
inst, lda, zeropage, 0x12
`

var instructionZeroPageX = []byte(`
Lda(ZeroPage(0x12), X)
`)
var instructionZeroPageXIr = `
inst, lda, zeropage x, 0x12
`

var instructionTestCases = []testCase{
	// {
	// 	"instruction with zeropage and x param",
	// 	instructionZeroPageX,
	// 	instructionZeroPageXIr,
	// 	"",
	// },
	// {
	// 	"instruction with zeropage param",
	// 	instructionZeroPage,
	// 	instructionZeroPageIr,
	// 	"",
	// },
	{
		"instruction with absolute param",
		instructionAbsolute,
		instructionAbsoluteIr,
		"",
	},
	{
		"instruction with absolute const param",
		instructionAbsoluteConst,
		instructionAbsoluteConstIr,
		"",
	},
	{
		"instruction with immediate param",
		instructionImmediate,
		instructionImmediateIr,
		"",
	},
	{
		"instruction without parameters",
		instructionNoParam,
		instructionNoParamIr,
		"",
	},
}

func TestInstruction(t *testing.T) {
	for _, test := range instructionTestCases {
		runTest(t, true, test)
	}
}
