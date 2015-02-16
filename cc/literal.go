package cc

import (
	"strconv"
)

type EmptyLiteral struct {
	SyntaxInfo
	Id int
}

func (x *EmptyLiteral) GetId() int {
	return x.Id
}

func (x *EmptyLiteral) String() string {
	return ""
}

func (x *EmptyLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

type BooleanLiteral struct {
	SyntaxInfo
	Id    int
	Value bool
}

func (x *BooleanLiteral) GetId() int {
	return x.Id
}

func (x *BooleanLiteral) String() string {
	if x.Value == true {
		return "true"
	} else {
		return "false"
	}
}

func (x *BooleanLiteral) GetChildren() []Syntax {
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

type CharLiteral struct {
	SyntaxInfo
	Id    int
	Value byte
}

func (x *CharLiteral) GetId() int {
	return x.Id
}

func (x *CharLiteral) String() string {
	return string([]byte{x.Value})
}

func (x *CharLiteral) GetChildren() []Syntax {
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

func (x *StringLiteral) ToSymbolLiteral() *SymbolLiteral {
	return &SymbolLiteral{
		SyntaxInfo: x.SyntaxInfo,
		Value:      x.Value,
		Id:         x.Id,
	}
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

func (x *SymbolLiteral) ToSymbolLiteral() *SymbolLiteral {
	return x
}

type LanguageKeyword struct {
	SyntaxInfo
	Id    int
	Value string
}

func (x *LanguageKeyword) GetId() int {
	return x.Id
}

func (x *LanguageKeyword) String() string {
	return x.Value
}

func (x *LanguageKeyword) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *LanguageKeyword) ToSymbolLiteral() *SymbolLiteral {
	return &SymbolLiteral{
		SyntaxInfo: x.SyntaxInfo,
		Value:      x.Value,
		Id:         x.Id,
	}
}
