// Code generated by gocc; DO NOT EDIT.

package token

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset  int
	Line    int
	Column  int
	Context Context
}

func (p Pos) String() string {
	// If the context provides a filename, provide a human-readable File:Line:Column representation.
	switch src := p.Context.(type) {
	case Sourcer:
		return fmt.Sprintf("%s:%d:%d", src.Source(), p.Line, p.Column)
	default:
		return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
	}
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	return fmt.Sprintf("%s(%d,%s)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

// Equals returns returns true if the token Type and Lit are matches.
func (t *Token) Equals(rhs interface{}) bool {
	switch rhsT := rhs.(type) {
	case *Token:
		return t == rhsT || (t.Type == rhsT.Type && bytes.Equal(t.Lit, rhsT.Lit))
	default:
		return false
	}
}

// CharLiteralValue returns the string value of the char literal.
func (t *Token) CharLiteralValue() string {
	return string(t.Lit[1 : len(t.Lit)-1])
}

// Float32Value returns the float32 value of the token or an error if the token literal does not
// denote a valid float32.
func (t *Token) Float32Value() (float32, error) {
	if v, err := strconv.ParseFloat(string(t.Lit), 32); err != nil {
		return 0, err
	} else {
		return float32(v), nil
	}
}

// Float64Value returns the float64 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Float64Value() (float64, error) {
	return strconv.ParseFloat(string(t.Lit), 64)
}

// IDValue returns the string representation of an identifier token.
func (t *Token) IDValue() string {
	return string(t.Lit)
}

// Int32Value returns the int32 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Int32Value() (int32, error) {
	if v, err := strconv.ParseInt(string(t.Lit), 10, 64); err != nil {
		return 0, err
	} else {
		return int32(v), nil
	}
}

// Int64Value returns the int64 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Int64Value() (int64, error) {
	return strconv.ParseInt(string(t.Lit), 10, 64)
}

// UTF8Rune decodes the UTF8 rune in the token literal. It returns utf8.RuneError if
// the token literal contains an invalid rune.
func (t *Token) UTF8Rune() (rune, error) {
	r, _ := utf8.DecodeRune(t.Lit)
	if r == utf8.RuneError {
		err := fmt.Errorf("Invalid rune")
		return r, err
	}
	return r, nil
}

// StringValue returns the string value of the token literal.
func (t *Token) StringValue() string {
	return string(t.Lit[1 : len(t.Lit)-1])
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"terminator",
		"kwdPackage",
		"identifier",
		"kwdImport",
		"(",
		")",
		".",
		"stringLit",
		"empty",
		"kwdVar",
		"=",
		"kwdType",
		"kwdInline",
		"kwdConst",
		"singleOperators",
		"operators",
		"relOp",
		"*",
		"intLit",
		",",
		"kwdFunc",
		"kwdVariadic",
		"type",
		"typeConstructor",
		"kwdInterface",
		":",
		"kwdBreak",
		"kwdRet",
		"kwdGoto",
		"{",
		"}",
		"kwdIf",
		"not",
		"kwdFor",
	},

	idMap: map[string]Type{
		"INVALID":         0,
		"$":               1,
		"terminator":      2,
		"kwdPackage":      3,
		"identifier":      4,
		"kwdImport":       5,
		"(":               6,
		")":               7,
		".":               8,
		"stringLit":       9,
		"empty":           10,
		"kwdVar":          11,
		"=":               12,
		"kwdType":         13,
		"kwdInline":       14,
		"kwdConst":        15,
		"singleOperators": 16,
		"operators":       17,
		"relOp":           18,
		"*":               19,
		"intLit":          20,
		",":               21,
		"kwdFunc":         22,
		"kwdVariadic":     23,
		"type":            24,
		"typeConstructor": 25,
		"kwdInterface":    26,
		":":               27,
		"kwdBreak":        28,
		"kwdRet":          29,
		"kwdGoto":         30,
		"{":               31,
		"}":               32,
		"kwdIf":           33,
		"not":             34,
		"kwdFor":          35,
	},
}
