package ast

// Identifier is an identifier declaration.
type Identifier struct {
	Name string
}

// NewIdentifier returns a new identifier.
func NewIdentifier(name string) (Node, error) {
	return NewIdentifierNoError(name), nil
}

// NewIdentifierNoError returns a new identifier.
func NewIdentifierNoError(name string) Node {
	return &Identifier{
		Name: name,
	}
}

// String implement the fmt.Stringer interface.
func (i Identifier) String() string {
	return i.Name
}
