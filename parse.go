package main

import (
	"bytes"
	"fmt"
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

func printArg(n int, more, less string) string {
	var s string
	for i := 0; i < n; i++ {
		s += more
	}
	for i := 0; i > n; i-- {
		s += less
	}
	return s
}

func (p Program) String() string {
	var s string
	for _, cmd := range p {
		s += fmt.Sprint(cmd)
	}
	return s
}

func (c Command) String() string {
	var s string
	switch c.Op {
	case Add:
		s = printArg(c.Arg, "+", "-")
	case Move:
		s = printArg(c.Arg, ">", "<")
	case Print:
		s = "."
	case Scan:
		s = ","
	case BNZ:
		s = "[" + fmt.Sprint(c.Branch) + "]"
	case Clear:
		s = "[-]"
	case Mult:
		s = "[-" + printArg(c.Off, ">", "<") +
			printArg(c.Arg, "+", "-") +
			printArg(-c.Off, ">", "<") + "]"
	}
	return s
}
