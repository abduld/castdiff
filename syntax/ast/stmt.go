// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

type Stmt interface {
	Syntax
	IsStmt() bool
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
