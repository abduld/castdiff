package ast

type TypeQual int

const (
	Const TypeQual = 1 << iota
	Volatile
	Restrict
)

func (q TypeQual) String() string {
	s := ""
	if q&Const != 0 {
		s += "const "
	}
	if q&Volatile != 0 {
		s += "volatile "
	}
	if q&Restrict != 0 {
		s += "restrict "
	}
	if s == "" {
		return ""
	}
	return s[:len(s)-1]
}
