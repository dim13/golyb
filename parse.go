package golyb

import (
	"bytes"
	"io"
	"io/ioutil"
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
	prog, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return parse(bytes.NewBuffer(prog)), nil
}

func ParseString(prog string) Program {
	return parse(bytes.NewBufferString(prog))
}

func parse(buf *bytes.Buffer) Program {
	var p Program
	for {
		r, _, err := buf.ReadRune()
		if err == io.EOF {
			return p
		}
		switch r {
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
			p = append(p, Command{Op: Loop, Branch: parse(buf)})
		case ']':
			return p
		}
	}
}
