// Package tests contains AST tests.
package tests

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
	goccErrors "github.com/retroenv/nesgo/internal/gocc/errors"
	"github.com/retroenv/nesgo/internal/gocc/lexer"
	"github.com/retroenv/nesgo/internal/gocc/parser"
)

var header = []byte(`package main

import . "github.com/retroenv/nesgo/pkg/nes"

`)
var headerMain = []byte(`func main() {
`)

var headerIr = `package, main
import, ., github.com/retroenv/nesgo/pkg/nes
`
var headerMainIr = `func, main
`

var footer = []byte(`}`)

func runTest(t *testing.T, useMainFunc bool, input []byte,
	expectedIr, expectedError, testDescription string) {
	t.Helper()

	buf := bytes.Buffer{}
	buf.Write(header)
	if useMainFunc {
		buf.Write(headerMain)
	}
	buf.Write(input)
	if useMainFunc {
		buf.Write(footer)
	}

	l := lexer.NewLexer(buf.Bytes())
	p := parser.NewParser()
	res, err := p.Parse(l)
	if expectedError == "" {
		assert.NoError(t, err, testDescription)
	} else {
		var e *goccErrors.Error
		assert.True(t, errors.As(err, &e), testDescription)
		assert.Error(t, e.Err, expectedError, testDescription)
		return
	}

	s := fmt.Sprint(res)
	s = strings.TrimPrefix(s, headerIr)
	if useMainFunc {
		s = strings.TrimPrefix(s, headerMainIr)
	}

	s = strings.TrimSpace(s)
	expectedIr = strings.TrimSpace(expectedIr)
	assert.Equal(t, expectedIr, s, testDescription)
}
