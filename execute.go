package golyb

import (
	"fmt"
	"io"
	"unicode"
)

type Tape interface {
	Move(off int)
	Read(off int) int
	Write(off, v int)
}

func (p Program) Execute(w io.Writer, r io.Reader, t Tape) {
	for i := range p {
		switch p[i].Op {
		case Add:
			x := t.Read(p[i].Off) + p[i].Arg
			t.Write(p[i].Off, x)
		case Move:
			t.Move(p[i].Off)
		case Print:
			if x := t.Read(p[i].Off); x > unicode.MaxASCII {
				fmt.Fprintf(w, "%d", x)
			} else {
				fmt.Fprintf(w, "%c", x)
			}
		case Scan:
			var x int
			fmt.Fscanf(r, "%c", &x)
			t.Write(p[i].Off, x)
		case Loop:
			for t.Read(0) != 0 {
				p[i].Branch.Execute(w, r, t)
			}
		case Clear:
			t.Write(p[i].Off, 0)
		case Mult:
			x := t.Read(p[i].Off) * p[i].Arg
			x += t.Read(p[i].Off + p[i].Dst)
			t.Write(p[i].Off+p[i].Dst, x)
		case Search:
			for t.Read(0) != 0 {
				t.Move(p[i].Off)
			}
		default:
			panic("unknown opcode")
		}
	}
}
