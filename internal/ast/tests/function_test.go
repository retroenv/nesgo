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
func test(_ ...inline) {
}
`)
var inlineFunctionIr = `
func, inline, test
`

var inlineFunctionParam = []byte(`
func test(data *uint8, _ ...inline) {
}
`)
var inlineFunctionParamIr = `
func, inline, (data), test
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

var functionTestCases = []struct {
	name          string
	input         []byte
	expectedIr    string
	expectedError string
}{
	{
		"function with body",
		functionWithBody,
		functionWithBodyIr,
		"",
	},
	{
		"function inlined with param",
		inlineFunctionParam,
		inlineFunctionParamIr,
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
		runTest(t, false, test.input, test.expectedIr, test.expectedError, test.name)
	}
}
