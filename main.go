package main

import (
	"flag"
	"log"
	"os"
)

var (
	file = flag.String("file", "", "Source file")
	out  = flag.String("out", "", "Output")
	opt  = flag.Bool("opt", true, "Optimize")
)

func output(out string) *os.File {
	if out == "" {
		return nil
	}
	file, _ := os.Create(out)
	return file
}

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
	if *opt {
		program = Optimize(program)
	}
	Execute(program, NewTape(output(*out)))
}
