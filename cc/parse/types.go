package cc

import (
	"fmt"

	. "github.com/abduld/castdiff/cc/ast"
)

type typeOp int

const (
	tChar typeOp = 1 << iota
	tShort
	tInt
	tLong
	tSigned
	tUnsigned
	tFloat
	tDouble
	tVoid
	tLonglong
)

var builtinTypes = map[typeOp]*Type{
	tChar:                     CharType,
	tChar | tSigned:           CharType,
	tChar | tUnsigned:         UcharType,
	tShort:                    ShortType,
	tShort | tSigned:          ShortType,
	tShort | tUnsigned:        UshortType,
	tShort | tInt:             ShortType,
	tShort | tSigned | tInt:   ShortType,
	tShort | tUnsigned | tInt: UshortType,
	tInt:                         IntType,
	tInt | tSigned:               IntType,
	tInt | tUnsigned:             UintType,
	tLong:                        LongType,
	tLong | tSigned:              LongType,
	tLong | tUnsigned:            UlongType,
	tLong | tInt:                 LongType,
	tLong | tSigned | tInt:       LongType,
	tLong | tUnsigned | tInt:     UlongType,
	tLonglong:                    LonglongType,
	tLonglong | tSigned:          LonglongType,
	tLonglong | tUnsigned:        UlonglongType,
	tLonglong | tInt:             LonglongType,
	tLonglong | tSigned | tInt:   LonglongType,
	tLonglong | tUnsigned | tInt: UlonglongType,
	tFloat:  FloatType,
	tDouble: DoubleType,
	tVoid:   VoidType,
}

func splitTypeWords(ws []Syntax) (c Storage, q TypeQual, ty *Type) {
	// Could check for doubled words in general,
	// like const const, but no one cares.
	var t typeOp
	var ts []Syntax
	for _, w := range ws {
		switch w.String() {
		case "const":
			q |= Const
		case "volatile":
			q |= Volatile
		case "auto":
			c |= Auto
		case "static":
			c |= Static
		case "extern":
			c |= Extern
		case "typedef":
			c |= Typedef
		case "register":
			c |= Register
		case "inline":
			c |= Inline
		case "char":
			t |= tChar
			ts = append(ts, w)
		case "short":
			t |= tShort
			ts = append(ts, w)
		case "int":
			t |= tInt
			ts = append(ts, w)
		case "long":
			if t&tLong != 0 {
				t ^= tLonglong | tLong
			} else {
				t |= tLong
			}
			ts = append(ts, w)
		case "signed":
			t |= tSigned
			ts = append(ts, w)
		case "unsigned":
			t |= tUnsigned
			ts = append(ts, w)
		case "float":
			t |= tFloat
			ts = append(ts, w)
		case "double":
			t |= tDouble
			ts = append(ts, w)
		case "void":
			t |= tVoid
			ts = append(ts, w)
		}
	}

	if t == 0 {
		t |= tInt
	}

	ty = builtinTypes[t]
	if ty == nil {
		tr := ""
		for _, t := range ts {
			tr += " " + t.String();
		}
		fmt.Printf("unsupported type %q\n", tr)
	}

	return c, q, builtinTypes[t]
}
