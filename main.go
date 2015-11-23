package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	file  = flag.String("file", "", "Source file")
	out   = flag.String("out", "", "Output file")
	tape  = flag.String("tape", "finite", "Tape type: finite/infinite")
	noopt = flag.Bool("noopt", false, "Disable optimization")
	debug = flag.Bool("debug", false, "Enable debugging")
	dump  = flag.Bool("dump", false, "Dump AST")
)

func output(out string) *os.File {
	if out == "" {
		return os.Stdout
	}
	file, _ := os.Create(out)
	return file
}

var storage = map[string]func(io.ReadWriter) Storage{
	"finite":   NewFiniteTape,
	"infinite": NewInfiniteTape,
}

func main() {
	flag.Parse()

	if *file == "" {
		flag.Usage()
		return
	}

	program, err := ParseFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	if !*noopt {
		program = Optimize(program)
	}

	if *dump {
		fmt.Println(program)
		return
	}

	if st, ok := storage[*tape]; ok {
		o := output(*out)
		Execute(program, st(o))
	} else {
		flag.Usage()
		return
	}
}
