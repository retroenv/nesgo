package ast

import (
	"fmt"
	"strings"
)

// Import is an import declaration.
type Import struct {
	Alias string
	Path  string
}

// NewImport an import declaration.
func NewImport(alias string, lib string) (Node, error) {
	lib = strings.Trim(lib, `"`)

	if alias != "." {
		sl := strings.Split(lib, "/")
		name := sl[len(sl)-1]
		alias = name
	}

	return &Import{
		Alias: alias,
		Path:  lib,
	}, nil
}

// String implement the fmt.Stringer interface.
func (i Import) String() string {
	return fmt.Sprintf("import, %s, %s", i.Alias, i.Path)
}
