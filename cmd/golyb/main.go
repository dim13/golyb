package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"

	"github.com/dim13/golyb"
	"github.com/dim13/golyb/dynamic"
	"github.com/dim13/golyb/optimize"
	"github.com/dim13/golyb/static"
)

var (
	file    = flag.String("file", "", "Source file (required)")
	in      = flag.String("in", "", "Input file")
	out     = flag.String("out", "", "Output file or /dev/null")
	profile = flag.String("profile", "", "Write CPU profile to file")
	tape    = flag.String("tape", "static", "Tape type: static or dynamic")
	dump    = flag.Bool("dump", false, "Dump AST and terminate")
	noop    = flag.Bool("noop", false, "Disable optimization")
	show    = flag.Bool("show", false, "Dump tape cells")
	store   = map[string]func(io.Reader, io.Writer) golyb.Storage{
		"static":  static.NewTape,
		"dynamic": dynamic.NewTape,
	}
)

func main() {
	flag.Parse()

	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

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

	storage, ok := store[*tape]
	if !ok {
		flag.Usage()
		return
	}

	var r io.Reader
	if *in != "" {
		r, err = os.Open(*in)
		if err != nil {
			log.Fatal(err)
		}
	}

	var w io.Writer
	if *out != "" {
		w, err = os.Create(*out)
		if err != nil {
			log.Fatal(err)
		}
	}

	st := storage(r, w)
	program.Execute(st)

	if *show {
		fmt.Println(st)
	}
}
