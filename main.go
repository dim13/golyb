package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
)

var (
	file    = flag.String("file", "", "Source file (required)")
	out     = flag.String("out", "", "Output file or /dev/null")
	tape    = flag.String("tape", "static", "Tape type: static or dynamic")
	noopt   = flag.Bool("noopt", false, "Disable optimization")
	debug   = flag.Bool("debug", false, "Enable debugging")
	dump    = flag.Bool("dump", false, "Dump AST")
	profile = flag.String("profile", "", "Write CPU profile to file")
)

func output(out string) (io.ReadWriter, error) {
	if out == "" {
		return os.Stdout, nil
	}
	return os.Create(out)
}

var storage = map[string]func(io.ReadWriter) Storage{
	"static":  NewStaticTape,
	"dynamic": NewDynamicTape,
}

func main() {
	flag.Parse()

	if *profile != "" {
		f, err := os.Create(*profile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *file == "" {
		flag.Usage()
		return
	}

	program, err := ParseFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	if !*noopt {
		program = OptContract(program)
		program = OptLoops(program)
		program = OptOffset(program)
	}

	if *dump {
		fmt.Println(program)
		return
	}

	if st, ok := storage[*tape]; ok {
		o, err := output(*out)
		if err != nil {
			log.Fatal(err)
		}
		Execute(program, st(o))
	} else {
		flag.Usage()
		return
	}
}
