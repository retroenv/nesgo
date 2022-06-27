package compiler

import (
	"bytes"
	"strings"
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

var testFileHeader = []byte(`package main

import . "github.com/retroenv/nesgo/pkg/nes"

func main() {
  Start(test)
}
`)

type testCase struct {
	name             string
	input            []byte
	expectedAssembly string
}

func runCompileTest(t *testing.T, test testCase) {
	t.Helper()

	cfg := &Config{
		DisableComments: true,
	}
	c, err := New(cfg)
	assert.NoError(t, err)

	buf := bytes.Buffer{}
	buf.Write(testFileHeader)
	buf.Write(test.input)

	assert.NoError(t, c.Parse("main.go", buf.Bytes()))

	assert.NoError(t, c.optimize())

	for _, fun := range c.functions {
		assert.NoError(t, c.outputFunction(fun))
	}

	s := strings.Join(c.output, "")
	s = strings.TrimSpace(s)
	expected := strings.TrimSpace(test.expectedAssembly)
	assert.Equal(t, expected, s, test.name)
}
