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

var instructionAbsoluteCastConst = []byte(`
Sta(Absolute(JOYPAD1))
`)
var instructionAbsoluteCastConstIr = `
inst, sta, absolute, JOYPAD1
`

var instructionAbsoluteX = []byte(`
Lda(Absolute(0x1234), X)
`)
var instructionAbsoluteXIr = `
inst, lda, absolute x, 0x1234
`

var instructionAbsoluteY = []byte(`
Lda(Absolute(0x1234), Y)
`)
var instructionAbsoluteYIr = `
inst, lda, absolute y, 0x1234
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
	{
		"instruction with absolute and y param",
		instructionAbsoluteY,
		instructionAbsoluteYIr,
		"",
	},
	{
		"instruction with absolute and x param",
		instructionAbsoluteX,
		instructionAbsoluteXIr,
		"",
	},
	{
		"instruction with absolute cast const param",
		instructionAbsoluteCastConst,
		instructionAbsoluteCastConstIr,
		"",
	},
	{
		"instruction with zeropage and x param",
		instructionZeroPageX,
		instructionZeroPageXIr,
		"",
	},
	{
		"instruction with zeropage param",
		instructionZeroPage,
		instructionZeroPageIr,
		"",
	},
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
