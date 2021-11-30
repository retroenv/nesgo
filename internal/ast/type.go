package ast

var typeInitializer = map[string]string{
	"NewInt8":   "int8",
	"NewUint8":  "uint8",
	"NewUint16": "uint16",
}

// Type is a type declaration.
type Type struct {
	Name            string
	InitializerUsed bool
}

// String implement the fmt.Stringer interface.
func (t Type) String() string {
	return t.Name
}

// NewType returns a type.
func NewType(name string) (Node, error) {
	t := &Type{}
	var ok bool
	t.Name, ok = typeInitializer[name]
	if ok {
		t.InitializerUsed = true
	} else {
		t.Name = name
	}
	return t, nil
}
