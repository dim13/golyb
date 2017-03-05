package main

import (
	"errors"
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

type Storage string

func (s *Storage) Set(v string) error {
	if v == "static" || v == "dynamic" {
		*s = Storage(v)
		return nil
	}
	return errors.New("unknown tape type")
}

func (s Storage) String() string {
	return string(s)
}

func (s Storage) Usage() string {
	return "Tape type: static or dynamic"
}

func (s Storage) New(r io.Reader, w io.Writer) golyb.Storage {
	switch s {
	case "static":
		return static.New(r, w)
	case "dynamic":
		return dynamic.New(r, w)
	}
	return nil
}

var (
	file    = flag.String("file", "", "Source file (required)")
	in      = flag.String("in", "", "Input file")
	out     = flag.String("out", "", "Output file or /dev/null")
	profile = flag.String("profile", "", "Write CPU profile to file")
	dump    = flag.Bool("dump", false, "Dump AST and terminate")
	noop    = flag.Bool("noop", false, "Disable optimization")
	show    = flag.Bool("show", false, "Dump tape cells")
	storage = Storage("static")
)

func init() {
	flag.Var(&storage, "tape", storage.Usage())
	flag.Parse()
}

func main() {
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

	var w io.Writer
	if *out != "" && *out != "-" {
		w, err = os.Create(*out)
		if err != nil {
			log.Fatal(err)
		}
	}

	tape := storage.New(r, w)
	program.Execute(tape)

	if *show {
		fmt.Println(tape)
	}
}
