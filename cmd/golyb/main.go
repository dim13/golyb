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

func main() {
	var (
		file    = flag.String("file", "", "Source file (required)")
		in      = flag.String("in", "", "Input file")
		out     = flag.String("out", "", "Output file or /dev/null")
		profile = flag.String("profile", "", "Write CPU profile to file")
		tape    = flag.String("tape", "static", "Tape type: static or dynamic")
		dump    = flag.Bool("dump", false, "Dump AST and terminate")
		noop    = flag.Bool("noop", false, "Disable optimization")
		show    = flag.Int("show", 0, "Dump # tape cells around last position")
		store   = map[string]func(io.ReadWriter) golyb.Storage{
			"static":  static.NewTape,
			"dynamic": dynamic.NewTape,
		}
	)
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

	if *file == "" {
		flag.Usage()
		return
	}

	program, err := golyb.ParseFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	if !*noop {
		program = optimize.Contract(program)
		program = optimize.Loops(program)
		program = optimize.Offset(program)
	}

	if *dump {
		fmt.Print(program)
		return
	}

	o, err := output(*out, *in)
	if err != nil {
		log.Fatal(err)
	}

	storage, ok := store[*tape]
	if !ok {
		flag.Usage()
		return
	}

	s := storage(o)
	program.Execute(s)

	if *show > 0 {
		cels, pos := s.Dump()
		from := pos - *show/2
		if from < 0 {
			from = 0
		}
		to := pos + *show/2
		if to > len(cels) {
			to = len(cels)
		}
		log.Println("From", from, "to", to, cels[from:to])
	}
}
