package compiler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/retroenv/nesgo/internal/ast"
)

func (p *Package) addConstants(constants []*ast.Constant) error {
	for _, con := range constants {
		name := con.Name
		if _, ok := p.constants[name]; ok {
			return fmt.Errorf("constant '%s' is defined multiple times", name)
		}
		p.constants[name] = con
	}
	return nil
}

func (p *Package) findConstant(packages map[string]*Package,
	caller, constant string) (*ast.Constant, error) {
	sl := strings.Split(constant, ".")
	if len(sl) > 1 {
		// TODO support non dot imported functions
		return nil, errors.New("non dot imports from external packages are not support yet")
	}

	if c, ok := p.constants[constant]; ok {
		return c, nil
	}

	imports := p.functionFile[caller].Imports
	for _, imp := range imports {
		impPack := packages[imp.Path]
		if c, ok := impPack.constants[constant]; ok {
			return c, nil
		}
	}

	return nil, fmt.Errorf("constant '%s' can not be found", constant)
}
