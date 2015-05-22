package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	cc "github.com/abduld/castdiff/syntax"
)

var (
	cfgFile = flag.String("c", "", "config file")
	inc     = flag.String("I", "", "include directory")
)

func main() {
	log.SetFlags(0)
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: castdiff [options] *.c\n")
		flag.PrintDefaults()
		os.Exit(2)
	}

	if *inc != "" {
		cc.AddInclude(*inc)
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
	}

	var r []io.Reader
	files := args
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		r = append(r, f)
		defer f.Close()
	}
	prog, err := cc.ReadMany(files, r)
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(prog)
	if err == nil {
		os.Stdout.Write(b)
	}
	//fmt.Println(prog)
}
