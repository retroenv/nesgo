package ast

import (
	"fmt"
	"strings"
)

// NewFile initializes file relevant data structures.
func NewFile(pkg *Package, content interface{}) (Node, error) {
	f := &File{
		Package: pkg,
	}
	list, ok := content.(*NodeList)
	if !ok {
		// file does not contain any code
		return f, nil
	}

	var err error
	for _, node := range list.Nodes {
		if l, ok := node.(*NodeList); ok {
			for _, n := range l.Nodes {
				if err = f.indexNode(n); err != nil {
					return nil, err
				}
			}
			continue
		}

		if err = f.indexNode(node); err != nil {
			return nil, err
		}
	}

	return f, nil
}

// File is a .go file.
type File struct {
	Package   *Package
	Imports   []*Import
	Constants []*Constant
	Variables []*Variable
	Functions []*Function
}

// String implement the fmt.Stringer interface.
func (f File) String() string {
	b := &strings.Builder{}
	_, _ = fmt.Fprintln(b, f.Package)
	for _, v := range f.Imports {
		_, _ = fmt.Fprintln(b, v)
	}
	for _, v := range f.Constants {
		_, _ = fmt.Fprintln(b, v)
	}
	for _, v := range f.Variables {
		_, _ = fmt.Fprintln(b, v)
	}
	for _, v := range f.Functions {
		_, _ = fmt.Fprintln(b, v)
	}
	return b.String()
}

func (f *File) indexNode(node Node) error {
	switch n := node.(type) {
	case *Constant:
		f.Constants = append(f.Constants, n)
	case *Function:
		f.Functions = append(f.Functions, n)
	case *Import:
		f.Imports = append(f.Imports, n)
	case *Variable:
		f.Variables = append(f.Variables, n)
	default:
		return fmt.Errorf("type %T is not supported as top file declaration", node)
	}
	return nil
}
