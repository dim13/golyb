package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
	"runtime/pprof"

	"github.com/dim13/golyb"
	"github.com/dim13/golyb/optimize"
)

var (
	file    = flag.String("file", "", "Source file (required)")
	in      = flag.String("in", "", "Input file")
	out     = flag.String("out", "", "Output file or /dev/null")
	profile = flag.String("profile", "", "Write CPU profile to file")
	dump    = flag.Bool("dump", false, "Dump AST and terminate")
	noop    = flag.Bool("noop", false, "Disable optimization")
	show    = flag.Bool("show", false, "Dump tape cells")
	tape    = Tape("static")
)

func main() {
	flag.Var(&tape, "tape", tape.Usage())
	flag.Parse()

	if *profile != "" {
		f, err := os.Create(*profile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	program, err := golyb.ParseFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	if !*noop {
		program = optimize.All(program)
	}

	if *dump {
		fmt.Print(program)
		return
	}

	var r io.Reader
	if *in != "" && *in != "-" {
		r, err = os.Open(*in)
		if err != nil {
			log.Fatal(err)
		}
	}
	if r == nil {
		r = os.Stdin
	}

	var w io.Writer
	if *out != "" && *out != "-" {
		w, err = os.Create(*out)
		if err != nil {
			log.Fatal(err)
		}
	}
	if w == nil {
		w = os.Stdout
	}

	mem := tape.New()
	defer stacktrace()
	golyb.Execute(w, r, program, mem)

	if *show {
		fmt.Println(mem)
	}
}

func stacktrace() {
	if r := recover(); r != nil {
		debug.PrintStack()
		log.Fatal(r)
	}
}
