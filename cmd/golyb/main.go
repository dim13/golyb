package main

import (
	"flag"
	"io"
	"log"
	"os"
	"runtime/pprof"

	"github.com/dim13/golyb"
)

func output(out, in string) (io.ReadWriter, error) {
	var err error
	var r io.Reader
	var w io.Writer

	if out != "" {
		w, err = os.Create(out)
		if err != nil {
			return nil, err
		}
	} else {
		w = os.Stdout
	}

	if in != "" {
		r, err = os.Open(in)
		if err != nil {
			return nil, err
		}
	} else {
		r = os.Stdin
	}
	return struct {
		io.Reader
		io.Writer
	}{r, w}, nil
}

var storage = map[string]func(io.ReadWriter) golyb.Storage{
	"static":  golyb.NewStaticTape,
	"dynamic": golyb.NewDynamicTape,
}

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
		program = program.Optimize()
	}

	if *dump {
		program.Dump()
		return
	}

	st, ok := storage[*tape]
	if !ok {
		flag.Usage()
		return
	}

	o, err := output(*out, *in)
	if err != nil {
		log.Fatal(err)
	}
	s := st(o)
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
