package compiler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/retroenv/nesgo/internal/ast"
)

var header = `.segment "HEADER"
.byte "NES", $1a ; Magic string that always begins an iNES header
.byte $02        ; Number of 16KB PRG-ROM banks
.byte $01        ; Number of 8KB CHR-ROM banks
.byte %00000001  ; Vertical mirroring, no save RAM, no mapper
.byte %00000000  ; No special-case flags set, no mapper
.byte $00        ; No PRG-RAM present
.byte $00        ; NTSC format

.segment "CODE"
`

var variableHeader = `.segment "HEADER"`

var footer = `.segment "VECTORS"
.addr %s, %s, %s

.segment "CHARS"
.res 8192
.segment "STARTUP"`

func (c *Compiler) generateProgramOutput() error {
	c.output = []string{header}

	for _, fun := range c.functions {
		if err := c.outputFunction(fun); err != nil {
			return err
		}
	}

	nmiHandler := "0"
	if c.nmiHandler != "" {
		nmiHandler = c.nmiHandler
	}
	irqHandler := "0"
	if c.irqHandler != "" {
		irqHandler = c.irqHandler
	}

	if err := c.outputVariables(); err != nil {
		return err
	}

	c.outputLine(footer, nmiHandler, c.resetHandler, irqHandler)
	return nil
}

func (c *Compiler) outputFunction(fun *Function) error {
	c.outputLine(".proc %s", fun.Definition.Name)

	for _, node := range fun.Body.Nodes {
		switch n := node.(type) {
		case *ast.Call:
			i := strings.LastIndex(n.Function, ".")
			label := n.Function[i+1:]
			c.outputLine("  jsr %s", label)

		case *ast.Instruction:
			if err := c.outputInstruction(n); err != nil {
				return fmt.Errorf("outputting instruction '%s': %w", n, err)
			}

		case *ast.Branching:
			ins := n.Instruction
			if ins == ast.GotoInstruction {
				ins = ast.JmpInstruction
			}
			c.outputLine("  %s %s", ins, n.DestinationName)

		case *ast.Label:
			c.outputLine("%s:", n.Name)

		case *ast.Statement:
			if n.Op == ast.NotOperator {
				continue
			}
		default:
			return fmt.Errorf("type %T is not supported as top file declaration", node)
		}
	}

	c.outputLine(".endproc\n")
	return nil
}

func (c *Compiler) outputLine(format string, a ...interface{}) {
	s := fmt.Sprintf(format+"\n", a...)
	c.output = append(c.output, s)
}

func (c *Compiler) outputLineWithComment(comment, format string, a ...interface{}) {
	if comment == "" {
		c.outputLine(format, a...)
		return
	}

	s := fmt.Sprintf(format, a...)
	s = fmt.Sprintf("%s  ; %s\n", s, comment)
	c.output = append(c.output, s)
}

func (c *Compiler) outputInstruction(ins *ast.Instruction) error {
	info := ast.CPUInstructions[ins.Name]

	switch len(ins.Arguments) {
	case 0:
		if !info.HasAddressing(ast.ImpliedAddressing | ast.AccumulatorAddressing) {
			return fmt.Errorf("instruction '%s' is missing a parameter", ins.Name)
		}
		c.outputLine("  %s", ins.Name)
		return nil

	case 1:
		return c.outputInstruction1Arg(ins, info)
	case 2:
		return c.outputInstruction2Args(ins, info)
	}

	return fmt.Errorf("instruction '%s' has unsupported parameters '%s'", ins.Name, ins.Arguments)
}

func (c *Compiler) outputInstruction1Arg(ins *ast.Instruction, info *ast.CPUInstruction) error {
	arg := ins.Arguments[0]
	node, ok := arg.(*ast.ArgumentValue)
	if !ok {
		return fmt.Errorf("wrong argument type %T for instruction with 1 arg", arg)
	}

	if info.HasAddressing(ast.RelativeAddressing) {
		c.outputLineWithComment(ins.Comment, "  %s %s", info.Alias, arg)
		return nil
	}
	if info.HasAddressing(ast.ImmediateAddressing) {
		val, err := strconv.ParseUint(node.Value, 0, 8)
		if err != nil {
			return fmt.Errorf("instruction '%s' has an invalid immediate parameter '%s'", ins.Name, arg)
		}
		c.outputLineWithComment(ins.Comment, "  %s #$%02x", info.Alias, val)
		return nil
	}
	if info.HasAddressing(ast.ZeroPageAddressing) {
		if val, err := strconv.ParseUint(node.Value, 0, 8); err == nil {
			c.outputLineWithComment(ins.Comment, "  %s $%02x", info.Alias, val)
			return nil
		}
	}
	if info.HasAddressing(ast.AbsoluteAddressing) {
		if val, err := strconv.ParseUint(node.Value, 0, 16); err == nil {
			c.outputLineWithComment(ins.Comment, "  %s $%04x", info.Alias, val)
			return nil
		}
		if _, ok := c.variables[node.Value]; ok {
			c.outputLineWithComment(ins.Comment, "  %s %s", info.Alias, node.Value)
			return nil
		}
	}
	return fmt.Errorf("instruction '%s' with 1 argument "+
		"has an unexpected parameter '%s'", ins.Name, arg)
}

func (c *Compiler) outputInstruction2Args(ins *ast.Instruction, info *ast.CPUInstruction) error {
	arg1 := ins.Arguments[0]
	node1, ok := arg1.(*ast.ArgumentValue)
	if !ok {
		return fmt.Errorf("wrong argument type %T for instruction", arg1)
	}
	arg2 := ins.Arguments[1]
	node2, ok := arg2.(*ast.ArgumentValue)
	if !ok {
		return fmt.Errorf("wrong argument type %T for instruction", arg2)
	}

	if node2.Value == "X" {
		if info.HasAddressing(ast.ZeroPageXAddressing) {
			if val, err := strconv.ParseUint(node1.Value, 0, 8); err == nil {
				c.outputLineWithComment(ins.Comment, "  %s $%02x, %s", info.Alias, val, arg2)
				return nil
			}
		}
		if info.HasAddressing(ast.AbsoluteXAddressing) {
			if val, err := strconv.ParseUint(node1.Value, 0, 16); err == nil {
				c.outputLineWithComment(ins.Comment, "  %s $%04x, %s", info.Alias, val, arg2)
				return nil
			}
		}
	}
	if node2.Value == "Y" {
		if info.HasAddressing(ast.ZeroPageYAddressing) {
			if val, err := strconv.ParseUint(node1.Value, 0, 8); err == nil {
				c.outputLineWithComment(ins.Comment, "  %s $%02x, %s", info.Alias, val, arg2)
				return nil
			}
		}
		if info.HasAddressing(ast.AbsoluteYAddressing) {
			if val, err := strconv.ParseUint(node1.Value, 0, 16); err == nil {
				c.outputLineWithComment(ins.Comment, "  %s $%04x, %s", info.Alias, val, arg2)
				return nil
			}
		}
	}

	return fmt.Errorf("instruction '%s' with 2 arguments "+
		"has an unexpected parameters '%s' and '%s'", ins.Name, arg1, arg2)
}

func (c *Compiler) outputVariables() error {
	if len(c.variables) == 0 {
		return nil
	}

	c.outputLine(variableHeader)

	for _, v := range c.variables {
		switch v.Type {
		case "int8", "uint8":
			c.outputLine("  %s: .res 1", v.Name)
		case "uint16":
			c.outputLine("  %s: .res 2", v.Name)
		default:
			return fmt.Errorf("variable type '%s' is not supported", v.Type)
		}
	}

	c.outputLine("")
	return nil
}

func formatNodeListComment(nodes []ast.Node) string {
	b := &strings.Builder{}
	for _, n := range nodes {
		switch t := n.(type) {
		case *ast.Statement:
			if t.Op == "cast" {
				continue
			}

			_, _ = fmt.Fprintf(b, "%s ", t.Op)

		case *ast.Type:
			_, _ = fmt.Fprintf(b, "%s ", t.Name)

		case *ast.Identifier:
			_, _ = fmt.Fprintf(b, "%s ", t.Name)

		case *ast.Value:
			_, _ = fmt.Fprintf(b, "%s ", t.Value)

		default:
			_, _ = fmt.Fprintf(b, "%s ", n)
		}
	}
	return b.String()
}
