package tests

import (
	"testing"
)

var emptyFunction = []byte(`
func test() {
}
`)
var emptyFunctionIr = `
func, test
`

var inlineFunction = []byte(`
func test(_ ...Inline) {
}
`)
var inlineFunctionIr = `
func, inline, test
`

var inlineFunctionSingleParam = []byte(`
func test(data *uint8, _ ...Inline) {
}
`)
var inlineFunctionSingleParamIr = `
func, inline, (data), test
`

var inlineFunctionMultipleParams = []byte(`
func test(data *uint8, value uint8, _ ...Inline) {
}
`)
var inlineFunctionMultipleParamsIr = `
func, inline, (data, value), test
`

var functionWithBody = []byte(`
func test() {
  Dex()
}
`)
var functionWithBodyIr = `
func, test
inst, dex
`

var functionWithReturn = []byte(`
func test() {
  return
}
`)
var functionWithReturnIr = `
func, test
inst, rts
`

var functionInlineRegisterParam = []byte(`
func test(X, _ ...Inline) {
}
`)
var functionInlineRegisterParamIr = `
func, inline, test
`

var functionRegisterParam = []byte(`
func test(index uint8) {
  Ldy(index)
  Dey()
}
`)
var functionRegisterParamIr = `
func, (index), test
inst, dey
`

var functionTestCases = []testCase{
	{
		"function with register as param",
		functionRegisterParam,
		functionRegisterParamIr,
		"",
	},
	{
		"inline function with register as param",
		functionInlineRegisterParam,
		functionInlineRegisterParamIr,
		"",
	},
	{
		"function with return in body",
		functionWithReturn,
		functionWithReturnIr,
		"",
	},
	{
		"function with body",
		functionWithBody,
		functionWithBodyIr,
		"",
	},
	{
		"function inlined with single param",
		inlineFunctionSingleParam,
		inlineFunctionSingleParamIr,
		"",
	},
	{
		"function inlined with multiple params",
		inlineFunctionMultipleParams,
		inlineFunctionMultipleParamsIr,
		"",
	},
	{
		"function inlined",
		inlineFunction,
		inlineFunctionIr,
		"",
	},
	{
		"empty function",
		emptyFunction,
		emptyFunctionIr,
		"",
	},
}

func TestFunction(t *testing.T) {
	for _, test := range functionTestCases {
		runTest(t, false, test)
	}
}
