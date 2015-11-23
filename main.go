package main

import (
	"flag"
	"log"
	"os"
)

var (
	file  = flag.String("file", "", "Source file")
	out   = flag.String("out", "stdout", "Output")
	debug = flag.Bool("debug", false, "Debug")
	opt   = flag.Bool("opt", true, "Optimize")
)

func output(out string) *os.File {
	if out == "stdout" {
		return os.Stdout
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
	if *debug {
		log.Println(program)
	}
	if *opt {
		program = Optimize(program)
	}
	if *debug {
		log.Println(program)
	}
	Execute(program, NewTape(output(*out)))
}
