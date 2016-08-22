package main

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
	BNZ
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

type Storage interface {
	Add(int, int)
	Move(int)
	Print(int)
	Scan(int)
	IsZero() bool
	Clear(int)
	Mult(int, int, int)
	Search(int)
	Dump() ([]int, int)
}

func (p Program) Execute(s Storage) {
	for _, cmd := range p {
		switch cmd.Op {
		case Add:
			s.Add(cmd.Arg, cmd.Off)
		case Move:
			s.Move(cmd.Arg)
		case Print:
			s.Print(cmd.Off)
		case Scan:
			s.Scan(cmd.Off)
		case BNZ:
			for !s.IsZero() {
				cmd.Branch.Execute(s)
			}
		case Clear:
			s.Clear(cmd.Off)
		case Mult:
			s.Mult(cmd.Dst, cmd.Arg, cmd.Off)
		case Search:
			s.Search(cmd.Arg)
		default:
			panic("unknown opcode")
		}
	}
}

func (p Program) Optimize() Program {
	return p.Contract().Loops().Offset()
}
