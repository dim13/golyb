package main

import (
	"bytes"
	"io"
	"io/ioutil"
)

type Opcode int

const (
	Add Opcode = iota
	Move
	Print
	Scan
	BNZ
	Clear
	Mult
	Search
)

type Command struct {
	Op     Opcode
	Arg    int
	Off    int
	Branch Program
}

type Program []Command

func parse(buf *bytes.Buffer) Program {
	var p Program
	for {
		r, _, err := buf.ReadRune()
		if err == io.EOF {
			return p
		}
		var cmd Command
		switch r {
		case '+':
			cmd = Command{Op: Add, Arg: 1}
		case '-':
			cmd = Command{Op: Add, Arg: -1}
		case '>':
			cmd = Command{Op: Move, Arg: 1}
		case '<':
			cmd = Command{Op: Move, Arg: -1}
		case '.':
			cmd = Command{Op: Print}
		case ',':
			cmd = Command{Op: Scan}
		case '[':
			cmd = Command{Op: BNZ, Branch: parse(buf)}
		case ']':
			return p
		default:
			continue
		}
		p = append(p, cmd)
	}
}

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

func ParseStringOptimized(prog string) Program {
	return Optimize(ParseString(prog))
}
