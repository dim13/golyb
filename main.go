package main

import (
	"flag"
	"log"
	"os"
)

var (
	file = flag.String("file", "", "Source file")
	out  = flag.String("out", "", "Output file")
	tape = flag.String("tape", "finite", "Tape type: finite/infinite")
	opt  = flag.Bool("opt", true, "Optimization")
)

func output(out string) *os.File {
	if out == "" {
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
	if *opt {
		program = Optimize(program)
	}
	o := output(*out)

	var st Storage
	switch *tape {
	case "finite":
		st = NewFiniteTape(o)
	case "infinite":
		st = NewInfiniteTape(o)
	default:
		flag.Usage()
		os.Exit(1)
	}

	Execute(program, st)
}
