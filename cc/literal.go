package cc

import (
	"strconv"
)

type BoolLiteral struct {
	SyntaxInfo
	Id    int
	Value bool
}

func (x *BoolLiteral) GetId() int {
	return x.Id
}

func (x *BoolLiteral) String() string {
	if x.Value == true {
		return "true"
	} else {
		return "false"
	}
}

func (x *BoolLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

type IntegerLiteral struct {
	SyntaxInfo
	Id    int
	Value int
}

func (x *IntegerLiteral) GetId() int {
	return x.Id
}

func (x *IntegerLiteral) String() string {
	return strconv.Itoa(x.Value)
}

func (x *IntegerLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

type RealLiteral struct {
	SyntaxInfo
	Id    int
	Value float64
}

func (x *RealLiteral) GetId() int {
	return x.Id
}

func (x *RealLiteral) String() string {
	return strconv.FormatFloat(x.Value, 'f', 6, 64)
}

func (x *RealLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

type StringLiteral struct {
	SyntaxInfo
	Id    int
	Value string
}

func (x *StringLiteral) GetId() int {
	return x.Id
}

func (x *StringLiteral) String() string {
	return x.Value
}

func (x *StringLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

type SymbolLiteral struct {
	SyntaxInfo
	Id    int
	Value string
}

func (x *SymbolLiteral) GetId() int {
	return x.Id
}

func (x *SymbolLiteral) String() string {
	return x.Value
}

func (x *SymbolLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

type Keyword struct {
	SyntaxInfo
	Id    int
	Value string
}

func (x *Keyword) GetId() int {
	return x.Id
}

func (x *Keyword) String() string {
	return x.Value
}

func (x *Keyword) GetChildren() []Syntax {
	return []Syntax{}
}
