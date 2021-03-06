// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cc implements parsing, type checking, and printing of C programs.
package cc

import (
	_ "encoding/json"
	"fmt"
	"log"
	"os"

	parse "github.com/abduld/castdiff/cc/parse"
)

func Test(file string) {

	r, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	prog, err := parse.Read(file, r)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prog.String())

	/*
		b, err := json.Marshal(prog)
		if err == nil {
			os.Stdout.Write(b)
		}
	*/
}
