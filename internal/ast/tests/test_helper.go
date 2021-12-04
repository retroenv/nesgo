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

type testCase struct {
	name          string
	input         []byte
	expectedIr    string
	expectedError string
}

func runTest(t *testing.T, useMainFunc bool, test testCase) {
	t.Helper()

	buf := bytes.Buffer{}
	buf.Write(header)
	if useMainFunc {
		buf.Write(headerMain)
	}
	buf.Write(test.input)
	if useMainFunc {
		buf.Write(footer)
	}

	l := lexer.NewLexer(buf.Bytes())
	p := parser.NewParser()
	res, err := p.Parse(l)
	if test.expectedError == "" {
		assert.NoError(t, err, test.name)
	} else {
		e := &goccErrors.Error{}
		if errors.As(err, &e) {
			assert.True(t, errors.As(err, &e), test.name)
			s := e.String()
			assert.True(t, strings.Contains(s, test.expectedError),
				fmt.Sprintf("%s:\n%s", test.name, s))
		} else {
			assert.Error(t, err, test.expectedError, test.name)
		}
		return
	}

	s := fmt.Sprint(res)
	s = strings.TrimPrefix(s, headerIr)
	if useMainFunc {
		s = strings.TrimPrefix(s, headerMainIr)
	}

	s = strings.TrimSpace(s)
	expectedIr := strings.TrimSpace(test.expectedIr)
	assert.Equal(t, expectedIr, s, test.name)
}
