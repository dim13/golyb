package golyb

import (
	"bytes"
	"io"
	"os"
	"strings"
)

//go:generate stringer -type=Opcode

type Opcode int

const (
	Add Opcode = iota
	Move
	Print
	Scan
	Loop
	Clear
	Mult
	Search
)

type Command struct {
	Op     Opcode
	Arg    int
	Dst    int
	Off    int
	Branch Program
}

type Program []Command

func ParseFile(fname string) (Program, error) {
	fd, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(fd)
	return parse(buf), nil
}

func ParseString(prog string) Program {
	return parse(strings.NewReader(prog))
}

func parse(r io.RuneReader) (p Program) {
	for {
		v, _, err := r.ReadRune()
		if err == io.EOF {
			return p
		}
		switch v {
		case '+':
			p = append(p, Command{Op: Add, Arg: 1})
		case '-':
			p = append(p, Command{Op: Add, Arg: -1})
		case '>':
			p = append(p, Command{Op: Move, Arg: 1})
		case '<':
			p = append(p, Command{Op: Move, Arg: -1})
		case '.':
			p = append(p, Command{Op: Print})
		case ',':
			p = append(p, Command{Op: Scan})
		case '[':
			p = append(p, Command{Op: Loop, Branch: parse(r)})
		case ']':
			return p
		}
	}
}
