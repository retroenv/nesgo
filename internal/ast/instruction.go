package ast

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Instruction is an instruction declaration.
type Instruction struct {
	Name      string
	Arguments Arguments
	Comment   string

	Addressing AddressingMode
}

// newInstruction creates an instruction specification.
func newInstruction(name string, arg interface{}) (*Instruction, error) {
	i := &Instruction{
		Name: name,
	}
	if arg != nil {
		info := CPUInstructions[name]
		if info == nil {
			return nil, fmt.Errorf("missing instruction info for '%s'", name)
		}
		if err := i.addArgument(info, arg); err != nil {
			return nil, err
		}
	}
	return i, nil
}

func (i *Instruction) addArgument(info *CPUInstruction, arg interface{}) error {
	switch val := arg.(type) {
	case *Identifier:
		if err := i.addIdentifierArgument(val); err != nil {
			return err
		}

	case *Value:
		i.addValueArgument(info, val)

	case *NodeList:
		for _, node := range val.Nodes {
			if err := i.addArgument(info, node); err != nil {
				return err
			}
		}

	case *ExpressionList:
		i.Arguments = append(i.Arguments, val)

	case *Call:
		if err := i.addCallArgument(info, val); err != nil {
			return err
		}

	default:
		return fmt.Errorf("type %T is not supported as instruction argument", arg)
	}
	return nil
}

func (i *Instruction) addIdentifierArgument(arg *Identifier) error {
	switch strings.ToUpper(arg.Name) {
	case "X":
		switch i.Addressing {
		case ZeroPageAddressing:
			i.Addressing = ZeroPageXAddressing
			return nil
		case AbsoluteAddressing:
			i.Addressing = AbsoluteXAddressing
			return nil
		case IndirectAddressing:
			i.Addressing = IndirectXAddressing
			return nil
		default:
			return errors.New("invalid instruction addressing mode used with X register")
		}

	case "Y":
		switch i.Addressing {
		case ZeroPageAddressing:
			i.Addressing = ZeroPageYAddressing
			return nil
		case AbsoluteAddressing:
			i.Addressing = AbsoluteYAddressing
			return nil
		case IndirectAddressing:
			i.Addressing = IndirectYAddressing
			return nil
		default:
			return errors.New("invalid instruction addressing mode used with Y register")
		}
	}

	i.Addressing = AbsoluteAddressing
	i.Arguments = append(i.Arguments, &ArgumentValue{Value: arg.Name})
	return nil
}

func (i *Instruction) addCallArgument(info *CPUInstruction, val *Call) error {
	switch val.Function {
	case "ZeroPage":
		i.Addressing = ZeroPageAddressing
	case "Absolute":
		i.Addressing = AbsoluteAddressing
	case "Indirect":
		i.Addressing = IndirectAddressing
	}

	for _, node := range val.Parameter {
		if err := i.addArgument(info, node); err != nil {
			return err
		}
	}
	return nil
}

func (i *Instruction) addValueArgument(info *CPUInstruction, val *Value) {
	i.Arguments = append(i.Arguments, &ArgumentValue{Value: val.Value})
	if i.Addressing != NoAddressing {
		return
	}
	if info.HasAddressing(ImmediateAddressing) {
		_, err := strconv.ParseUint(val.Value, 0, 8)
		if err == nil {
			i.Addressing = ImmediateAddressing
			return
		}
	}
	if info.HasAddressing(ZeroPageAddressing) {
		_, err := strconv.ParseUint(val.Value, 0, 8)
		if err == nil {
			i.Addressing = ZeroPageAddressing
			return
		}
	}
	i.Addressing = AbsoluteAddressing
}

// String implement the fmt.Stringer interface.
func (i Instruction) String() string {
	b := &strings.Builder{}
	_, _ = fmt.Fprintf(b, "inst, %s", i.Name)
	switch i.Addressing {
	case ImmediateAddressing:
		_, _ = fmt.Fprint(b, ", immediate")
	case AbsoluteAddressing:
		_, _ = fmt.Fprint(b, ", absolute")
	case AbsoluteXAddressing:
		_, _ = fmt.Fprint(b, ", absolute x")
	case AbsoluteYAddressing:
		_, _ = fmt.Fprint(b, ", absolute y")
	case ZeroPageAddressing:
		_, _ = fmt.Fprint(b, ", zeropage")
	case ZeroPageXAddressing:
		_, _ = fmt.Fprint(b, ", zeropage x")
	case ZeroPageYAddressing:
		_, _ = fmt.Fprint(b, ", zeropage y")
	case IndirectXAddressing:
		_, _ = fmt.Fprint(b, ", indirect x")
	case IndirectYAddressing:
		_, _ = fmt.Fprint(b, ", indirect y")
	}

	if len(i.Arguments) > 0 {
		_, _ = fmt.Fprintf(b, ", %s", i.Arguments)
	}
	s := b.String()
	return s
}
