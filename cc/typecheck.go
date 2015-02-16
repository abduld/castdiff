// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Scoping and type checking.
// C99 standard: http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1124.pdf

package cc

type Scope struct {
	Decl map[string]*Decl
	Tag  map[string]*Type
	Next *Scope
}

func (lx *lexer) pushDecl(decl *Decl) {
	sc := lx.scope
	if sc == nil {
		panic("no scope")
	}
	if decl.Name.String() == "" {
		return
	}
	if sc.Decl == nil {
		sc.Decl = make(map[string]*Decl)
	}
	sc.Decl[decl.Name.String()] = decl
	if hdr := lx.declSave; hdr != nil && sc.Next == nil {
		hdr.decls = append(hdr.decls, decl)
	}
}

func (lx *lexer) lookupDecl(name string) *Decl {
	for sc := lx.scope; sc != nil; sc = sc.Next {
		decl := sc.Decl[name]
		if decl != nil {
			return decl
		}
	}
	return nil
}

func (lx *lexer) pushType(typ *Type) *Type {
	sc := lx.scope
	if sc == nil {
		panic("no scope")
	}

	if typ.Kind == Enum && typ.Decls != nil {
		for _, decl := range typ.Decls {
			lx.pushDecl(decl)
		}
	}

	if typ.Tag.String() == "" {
		return typ
	}

	old := lx.lookupTag(typ.Tag.String())
	if old == nil {
		if sc.Tag == nil {
			sc.Tag = make(map[string]*Type)
		}
		sc.Tag[typ.Tag.String()] = typ
		if hdr := lx.declSave; hdr != nil && sc.Next == nil {
			hdr.types = append(hdr.types, typ)
		}
		return typ
	}

	// merge typ into old
	if old.Kind != typ.Kind {
		lx.Errorf("conflicting tags: %s %s and %s %s", old.Kind, old.Tag, typ.Kind, typ.Tag)
		return typ
	}
	if typ.Decls != nil {
		if old.Decls != nil {
			lx.Errorf("multiple definitions for %s %s", old.Kind, old.Tag)
		}
		old.SyntaxInfo = typ.SyntaxInfo
		old.Decls = typ.Decls
	}
	return old
}

func (lx *lexer) lookupTag(name string) *Type {
	for sc := lx.scope; sc != nil; sc = sc.Next {
		typ := sc.Tag[name]
		if typ != nil {
			return typ
		}
	}
	return nil
}
