package ast

type Storage int

const (
	Auto Storage = 1 << iota
	Static
	Extern
	Typedef
	Register
	Inline
	Global
	Device
	Host
)

func (c Storage) String() string {
	s := ""
	if c&Auto != 0 {
		s += "auto "
	}
	if c&Static != 0 {
		s += "static "
	}
	if c&Extern != 0 {
		s += "extern "
	}
	if c&Typedef != 0 {
		s += "typedef "
	}
	if c&Register != 0 {
		s += "register "
	}
	if c&Inline != 0 {
		s += "inline "
	}
	if c&Global != 0 {
		s += "__global__ "
	}
	if c&Device != 0 {
		s += "__device__ "
	}
	if c&Host != 0 {
		s += "__host__ "
	}
	if s == "" {
		return ""
	}
	return s[:len(s)-1]
}
