package compiler

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/retroenv/nesgo/internal/ast"
)

// Package represents a code package.
type Package struct {
	name string

	// key is the filename
	files map[string]*File
	// key is the constant name
	constants map[string]*ast.Constant
	// key is the variable name
	variables map[string]*ast.Variable
	// key is the function name
	functions map[string]*Function
	// key is the function name, maps back to file for import parsing
	functionFile map[string]*File
}

func parsePackage(name string) (*Package, []string, error) {
	packName, directory, err := currentPackage()
	if err != nil {
		return nil, nil, fmt.Errorf("getting current package: %w", err)
	}

	pack := newPackage(packName)
	// TODO support external packages
	if !strings.HasPrefix(name, packName) {
		if name == "fmt" {
			return pack, nil, nil
		}
		return nil, nil, errors.New("external packages are not support yet")
	}

	dir := path.Join(directory, strings.TrimPrefix(name, packName))
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, nil, fmt.Errorf("reading package '%s' directory: %w", name, err)
	}

	var subPackages []string
	for _, entry := range files {
		entryName := entry.Name()
		if entry.Type() == os.ModeDir {
			subDir := path.Join(name, entryName)
			subPackages = append(subPackages, subDir)
			continue
		}
		if strings.HasSuffix(strings.ToLower(entryName), testFileSuffix) {
			continue
		}

		fullPath := path.Join(dir, entryName)
		data, err := os.ReadFile(fullPath)
		if err != nil {
			return nil, nil, fmt.Errorf("reading file: %w", err)
		}
		file, err := parseFile(fullPath, data)
		if err != nil {
			return nil, nil, fmt.Errorf("parsing file '%s': %w", fullPath, err)
		}
		if file.IsIgnored {
			continue
		}
		pack.name = file.Package // update package name once a file is parsed
		if err = pack.addFile(entryName, file); err != nil {
			return nil, nil, fmt.Errorf("processing file '%s': %w", fullPath, err)
		}
	}
	return pack, subPackages, nil
}

func newPackage(name string) *Package {
	return &Package{
		name:         name,
		files:        map[string]*File{},
		constants:    map[string]*ast.Constant{},
		variables:    map[string]*ast.Variable{},
		functions:    map[string]*Function{},
		functionFile: map[string]*File{},
	}
}

func (p *Package) addFile(fileName string, file *File) error {
	p.files[fileName] = file

	for _, fun := range file.Functions {
		if err := p.addFunction(fun, file); err != nil {
			return err
		}
	}

	if err := p.addConstants(file); err != nil {
		return fmt.Errorf("adding constants: %w", err)
	}
	if err := p.addVariables(file.Variables); err != nil {
		return fmt.Errorf("adding variables: %w", err)
	}
	return nil
}

func (p *Package) addFunction(astFun *ast.Function, file *File) error {
	s := astFun.Definition.Name
	if _, ok := p.functions[s]; ok {
		return fmt.Errorf("function '%s' defined multiple times", s)
	}

	fun := &Function{
		Definition: astFun.Definition,
		Body:       astFun.Body,
		Package:    p,
		Labels:     map[string]*ast.Label{},
	}

	if err := collectAndLinkAllLabels(fun); err != nil {
		return err
	}

	p.functions[s] = fun
	p.functionFile[s] = file
	return nil
}

func (p *Package) findFunction(packages map[string]*Package,
	caller, function string) (string, *Function, error) {
	sl := strings.Split(function, ".")
	if len(sl) > 1 {
		// TODO support non dot imported functions
		return "", nil, errors.New("non dot imports from external packages are not support yet")
	}

	// check if it's part of the main package
	if f, ok := p.functions[function]; ok {
		return function, f, nil
	}

	imports := p.functionFile[caller].Imports
	for _, imp := range imports {
		impPack := packages[imp.Path]
		if impPack == nil {
			continue
		}
		if f, ok := impPack.functions[function]; ok {
			s := fmt.Sprintf("%s.%s", imp.Path, function)
			return s, f, nil
		}
	}

	return "", nil, fmt.Errorf("function '%s' not found", function)
}
