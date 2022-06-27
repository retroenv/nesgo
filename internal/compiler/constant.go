package compiler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/retroenv/nesgo/internal/ast"
)

func (p *Package) addConstants(file *File) error {
	for _, con := range file.Constants {
		name := con.Name
		if _, ok := p.constants[name]; ok {
			return fmt.Errorf("constant '%s' is defined multiple times", name)
		}

		if con.AliasName != "" {
			imp := file.importLookup[con.AliasPackage]
			con.AliasPackage = imp.Path
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

	if con, ok := p.constants[constant]; ok {
		return con, nil
	}

	imports := p.functionFile[caller].Imports
	for _, imp := range imports {
		impPack := packages[imp.Path]
		if impPack == nil {
			continue
		}
		if con, ok := impPack.constants[constant]; ok {
			return con, nil
		}
	}

	return nil, fmt.Errorf("constant '%s' can not be found", constant)
}

func (p *Package) resolveConstantAlias(packages map[string]*Package,
	aliasCon *ast.Constant) error {
	pack, ok := packages[aliasCon.AliasPackage]
	if !ok {
		return fmt.Errorf("constant alias package '%s' can not be found", aliasCon.AliasPackage)
	}

	con, ok := pack.constants[aliasCon.AliasName]
	if !ok {
		return fmt.Errorf("constant alias '%s' in package '%s' can not be found",
			aliasCon.AliasName, aliasCon.AliasPackage)
	}

	aliasCon.Value = con.Value
	aliasCon.AliasName = ""
	aliasCon.AliasPackage = ""
	return nil
}
