package main

import (
	"flag"
	"log"
	"os"
)

var (
	file  = flag.String("file", "", "Source file")
	opt   = flag.Bool("opt", true, "Optimize")
	debug = flag.Bool("debug", false, "Debug")
)

func main() {
	flag.Parse()
	if *file == "" {
		flag.Usage()
		os.Exit(1)
	}
	program, err := ParseFile(*file)
	if err != nil {
		log.Fatal(err)
	}
	if *debug {
		log.Println(program)
	}
	if *opt {
		program = Optimize(program)
	}
	if *debug {
		log.Println(program)
	}
	Execute(program, NewTape())
}
