package ast

// Value is a value definition.
type Value struct {
	Value string
}

// NewValue returns a value.
func NewValue(val string) (Node, error) {
	return &Value{
		Value: val,
	}, nil
}

// String implement the fmt.Stringer interface.
func (v Value) String() string {
	return v.Value
}
