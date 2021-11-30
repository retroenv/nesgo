package ast

import "fmt"

// Package is a package declaration.
type Package struct {
	Name string
}

// NewPackage initializes a package declaration.
func NewPackage(name string) (Node, error) {
	return &Package{
		Name: name,
	}, nil
}

// String implement the fmt.Stringer interface.
func (p Package) String() string {
	return fmt.Sprintf("package, %s", p.Name)
}
