// Package compiler provides functionality for parsing Golang code.
package compiler

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/retroenv/nesgo/internal/ast"
)

const cpuRegisterSize = 16

var (
	buildHeader1      = []byte("// +build ")
	buildHeader2      = []byte("//go:build ")
	mainContextPrefix = "main."
	nesGoIgnoreTag    = "!nesgo"
	testFileSuffix    = "_test.go"
)

// Compiler defines a new compiler.
type Compiler struct {
	cfg *Config

	packages       map[string]*Package
	packagesToLoad map[string]struct{}

	functionsAdded   map[string]*Function
	functionsToParse map[string]*Function

	// info to output
	functions            []*Function
	variables            map[string]*ast.Variable
	variablesInitialized map[string]*ast.Variable
	resetHandler         string
	nmiHandler           string
	irqHandler           string
	output               []string
}

// New returns a new compiler.
func New(cfg *Config) (*Compiler, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &Compiler{
		cfg: cfg,

		packages:       map[string]*Package{},
		packagesToLoad: map[string]struct{}{},

		functionsAdded:   map[string]*Function{},
		functionsToParse: map[string]*Function{},

		variables:            map[string]*ast.Variable{},
		variablesInitialized: map[string]*ast.Variable{},
	}, nil
}

// Parse parses a file content and it's imports.
func (c *Compiler) Parse(fileName string, data []byte) error {
	file, err := parseFile(fileName, data)
	if err != nil {
		return fmt.Errorf("parsing file: %w", err)
	}
	if file.IsIgnored {
		return fmt.Errorf("file '%s' has ignore header set", fileName)
	}

	pack := newPackage("main")
	if err = pack.addFile(fileName, file); err != nil {
		return fmt.Errorf("processing file '%s': %w", fileName, err)
	}
	c.packages[file.Package] = pack
	c.updatePackagesToLoad(file.Imports)

	for len(c.packagesToLoad) > 0 {
		for packageName := range c.packagesToLoad {
			if pack, err = parsePackage(packageName); err != nil {
				return fmt.Errorf("parsing package '%s': %w", fileName, err)
			}
			c.packages[packageName] = pack
			delete(c.packagesToLoad, packageName)

			for _, file = range pack.files {
				c.updatePackagesToLoad(file.Imports)
			}
		}
	}

	if err := c.resolveTypeAliases(); err != nil {
		return fmt.Errorf("resolving type aliases': %w", err)
	}

	return nil
}

func (c *Compiler) resolveTypeAliases() error {
	for _, pack := range c.packages {
		for _, con := range pack.constants {
			if con.AliasName == "" {
				continue
			}

			if err := pack.resolveConstantAlias(c.packages, con); err != nil {
				return err
			}
		}
	}
	return nil
}

// OutputAsmFile creates an .asm file based on the given .nes file name as base.
func (c *Compiler) OutputAsmFile(fileName string) (string, string, error) {
	if err := c.optimize(); err != nil {
		return "", "", fmt.Errorf("optimizing AST: %w", err)
	}

	if err := c.generateProgramOutput(); err != nil {
		return "", "", fmt.Errorf("generating program output: %w", err)
	}

	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	asmFile := baseName + ".asm"
	objectFile := baseName + ".o"
	f, err := os.Create(asmFile)
	if err != nil {
		return "", "", fmt.Errorf("creating file '%s': %w", baseName, err)
	}

	for _, l := range c.output {
		if _, err = f.Write([]byte(l)); err != nil {
			_ = f.Close()
			return "", "", fmt.Errorf("writing file '%s': %w", baseName, err)
		}
	}
	return asmFile, objectFile, f.Close()
}

// optimize the AST.
func (c *Compiler) optimize() error {
	mainPackage := c.packages["main"]
	if mainPackage == nil {
		return errors.New("no main package found")
	}

	if err := c.addHandlersToParse(mainPackage); err != nil {
		return err
	}

	for len(c.functionsToParse) > 0 {
		for name, fun := range c.functionsToParse {
			if err := c.resolveFunctionNodes(fun); err != nil {
				return fmt.Errorf("processing function '%s': %w", name, err)
			}
			if err := c.addFunction(name, fun); err != nil {
				return err
			}
		}
	}

	c.processIrqHandlers()
	for _, fun := range c.functionsAdded {
		fun.addFunctionReturn()
	}

	if len(c.variablesInitialized) > 0 {
		if err := c.createVariableInitializations(mainPackage); err != nil {
			return fmt.Errorf("creating variable initialization function: %w", err)
		}
	}

	return c.inlineFunctions()
}

// addHandlersToParse parses the main function to get the entrypoints
// for the NES handlers.
func (c *Compiler) addHandlersToParse(mainPackage *Package) error {
	mainFunc := mainPackage.functions["main"]
	if mainFunc == nil {
		return errors.New("no main function found")
	}
	if err := c.resolveFunctionNodes(mainFunc); err != nil {
		return fmt.Errorf("processing main: %w", err)
	}
	if c.resetHandler == "" {
		return errors.New("no reset handler set, use Start() in main function")
	}

	resetHandler := mainPackage.functions[c.resetHandler]
	if resetHandler == nil {
		return fmt.Errorf("reset handler function '%s' not found", c.resetHandler)
	}
	c.functionsToParse[mainContextPrefix+resetHandler.Definition.Name] = resetHandler

	if c.nmiHandler != "" {
		nmiHandler := mainPackage.functions[c.nmiHandler]
		if nmiHandler == nil {
			return fmt.Errorf("nmi handler function '%s' not found", c.nmiHandler)
		}
		c.functionsToParse[mainContextPrefix+nmiHandler.Definition.Name] = nmiHandler
	}
	if c.irqHandler != "" {
		irqHandler := mainPackage.functions[c.irqHandler]
		if irqHandler == nil {
			return fmt.Errorf("irq handler function '%s' not found", c.irqHandler)
		}
		c.functionsToParse[mainContextPrefix+irqHandler.Definition.Name] = irqHandler
	}
	return nil
}

// updatePackagesToLoad adds imports that are not loaded yet to the list
// of packages to load.
func (c *Compiler) updatePackagesToLoad(imports []*ast.Import) {
	for _, imp := range imports {
		if _, ok := c.packages[imp.Path]; !ok {
			c.packagesToLoad[imp.Path] = struct{}{}
		}
	}
}
