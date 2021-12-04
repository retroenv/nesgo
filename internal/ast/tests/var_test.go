package tests

import (
	"testing"

	"github.com/retroenv/nesgo/internal/ast"
)

var singleVarType = []byte(`
var i int8
`)
var singleVarTypeIr = `
var, i, int8
`

var multipleVarsType = []byte(`
var i, j int8
`)
var multipleVarsTypeIr = `
var, i, int8
var, j, int8
`

var multipleVarsTypeInitializer = []byte(`
var i, j = NewUint8(2)
`)
var multipleVarsTypeInitializerIr = `
var, i, uint8, 2
var, j, uint8, 2
`

var singleVarInvalidType = []byte(`
var i int128
`)
var singleVarTypeInvalidInitializer = []byte(`
var i int8 = 1
`)
var singleVarTypeInvalidName = []byte(`
var x int8
`)

var singleVarTypeValidInitializer = []byte(`
var i = NewUint8(1)
`)
var singleVarTypeValidInitializerIr = `
var, i, uint8, 1
`

var varGroupSingleType = []byte(`
var (
  i int8
)
`)
var varGroupSingleTypeIr = `
var, i, int8
`

var varGroupMultipleTypes = []byte(`
var (
  i int8
  j int8
)
`)
var varGroupMultipleTypesIr = `
var, i, int8
var, j, int8
`

var varTestCases = []testCase{
	{
		"var list declaration with type and valid initializer",
		multipleVarsTypeInitializer,
		multipleVarsTypeInitializerIr,
		"",
	},
	{
		"var list declaration with type",
		multipleVarsType,
		multipleVarsTypeIr,
		"",
	},
	{
		"var group declaration with multiple types",
		varGroupMultipleTypes,
		varGroupMultipleTypesIr,
		"",
	},
	{
		"var group declaration with single type",
		varGroupSingleType,
		varGroupSingleTypeIr,
		"",
	},
	{
		"single var declaration with reserve name",
		singleVarTypeInvalidName,
		"",
		ast.ErrInvalidVariableName.Error(),
	},
	{
		"single var declaration with type and invalid initializer",
		singleVarTypeInvalidInitializer,
		"",
		ast.ErrInvalidInitializer.Error(),
	},
	{
		"single var declaration with type and valid initializer",
		singleVarTypeValidInitializer,
		singleVarTypeValidInitializerIr,
		"",
	},
	{
		"single var declaration with invalid type",
		singleVarInvalidType,
		"",
		"Expected one of:",
	},
	{
		"single var declaration with type",
		singleVarType,
		singleVarTypeIr,
		"",
	},
}

func TestVar(t *testing.T) {
	for _, test := range varTestCases {
		runTest(t, false, test)
	}
}
