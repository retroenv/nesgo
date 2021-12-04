package tests

import (
	"testing"
)

var singleConstValueDec = []byte(`
const bg_color = 1
`)
var singleConstValueDecIr = `
const, bg_color, 1
`

var singleConstValueHex = []byte(`
const bg_color = 0x1a
`)
var singleConstValueHexIr = `
const, bg_color, 26
`

var singleConstValueBin = []byte(`
const bg_color = 0b10000000
`)
var singleConstValueBinIr = `
const, bg_color, 128
`

var constTestCases = []testCase{
	{
		"single const declaration with binary value",
		singleConstValueBin,
		singleConstValueBinIr,
		"",
	},
	{
		"single const declaration with hex value",
		singleConstValueHex,
		singleConstValueHexIr,
		"",
	},
	{
		"single const declaration with decimal value",
		singleConstValueDec,
		singleConstValueDecIr,
		"",
	},
}

func TestConst(t *testing.T) {
	for _, test := range constTestCases {
		runTest(t, false, test)
	}
}
