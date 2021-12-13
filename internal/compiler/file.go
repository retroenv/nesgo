package compiler

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/retroenv/nesgo/internal/ast"
	"github.com/retroenv/nesgo/internal/gocc/lexer"
	"github.com/retroenv/nesgo/internal/gocc/parser"
)

// File is a .go file.
type File struct {
	Path      string
	IsIgnored bool
	Package   string
	Imports   []*ast.Import
	Constants []*ast.Constant
	Variables []*ast.Variable
	Functions []*ast.Function
}

// parseFile parses the file using the lexer and parser and returns an AST
// presentation of the file.
func parseFile(fileName string, data []byte) (*File, error) {
	ignored, err := isFileIgnored(data)
	if err != nil {
		return nil, err
	}
	f := &File{
		Path: fileName,
	}
	if ignored {
		f.IsIgnored = true
		return f, nil
	}

	l := lexer.NewLexer(data)
	p := parser.NewParser()

	res, err := p.Parse(l)
	if err != nil {
		return nil, fmt.Errorf("parsing file: %w", err)
	}

	astFile, ok := res.(*ast.File)
	if !ok {
		return nil, fmt.Errorf("unexpected file parse type %T", res)
	}
	f.Package = astFile.Package.Name
	f.Imports = astFile.Imports
	f.Constants = astFile.Constants
	f.Variables = astFile.Variables
	f.Functions = astFile.Functions

	return f, nil
}

// isFileIgnored returns whether the file is to be ignored by the
// compiler due to a set built flag.
func isFileIgnored(data []byte) (bool, error) {
	reader := bytes.NewReader(data)
	buf := bufio.NewReader(reader)

	for i := 0; i < 2; i++ {
		b, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return false, nil
			}
			return false, fmt.Errorf("reading line: %w", err)
		}
		b = bytes.TrimSpace(b)

		var tags []string
		if bytes.HasPrefix(b, buildHeader1) {
			b = b[len(buildHeader1):]
			tags = strings.Split(string(b), ",")
		} else {
			if !bytes.HasPrefix(b, buildHeader2) {
				continue
			}
			b = b[len(buildHeader2):]
			tags = strings.Split(string(b), "&&")
		}

		for _, tag := range tags {
			if strings.TrimSpace(tag) == nesGoIgnoreTag {
				return true, nil
			}
		}
	}

	return false, nil
}
