package tests

import (
	"testing"
)

var callNoParam = []byte(`
Init()
`)
var callNoParamIr = `
call, Init
`

var callSingleParam = []byte(`
StartPPUTransfer(PALETTE_START)
`)
var callSingleParamIr = `
call, StartPPUTransfer, PALETTE_START
`

var callSingleParamExpression = []byte(`
PPUMask(MASK_BG_CLIP | MASK_SPR_CLIP | MASK_BG | MASK_SPR)
`)
var callSingleParamExpressionIr = `
call, PPUMask, MASK_BG_CLIP
op, |, 
MASK_SPR_CLIP
op, |, 
MASK_BG
op, |, 
MASK_SPR
`

var callFmtDebugPrint = []byte(`
Dex()
fmt.Println("debug")
`)
var callFmtDebugPrintIr = `
inst, dex
`

var callTestCases = []struct {
	name          string
	input         []byte
	expectedIr    string
	expectedError string
}{
	{
		"fmt debug print call",
		callFmtDebugPrint,
		callFmtDebugPrintIr,
		"",
	},
	{
		"call without parameters",
		callNoParam,
		callNoParamIr,
		"",
	},
	{
		"call with 1 parameter",
		callSingleParam,
		callSingleParamIr,
		"",
	},
	{
		"call with 1 parameter expression",
		callSingleParamExpression,
		callSingleParamExpressionIr,
		"",
	},
}

func TestCall(t *testing.T) {
	for _, test := range callTestCases {
		runTest(t, true, test)
	}
}
