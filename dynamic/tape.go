package dynamic

import (
	"fmt"
	"io"
	"unicode"

	"github.com/dim13/golyb"
)

type Cell int

type Tape struct {
	cell []Cell
	pos  int
	out  io.ReadWriter
}

const chunkSize = 1024

func NewTape(out io.ReadWriter) golyb.Storage {
	return &Tape{
		cell: make([]Cell, chunkSize),
		pos:  0,
		out:  out,
	}
}

func (t *Tape) grow(pos int) {
	if pos >= len(t.cell) {
		t.cell = append(t.cell, make([]Cell, chunkSize)...)
	}
	if pos < 0 {
		t.cell = append(make([]Cell, chunkSize), t.cell...)
		t.pos += chunkSize
	}
}

func (t *Tape) Move(n int) {
	t.pos += n
	t.grow(t.pos)
}

func (t *Tape) Add(n, off int) {
	x := t.pos + off
	t.grow(x)
	t.cell[x] += Cell(n)
}

func (t *Tape) Print(off int) {
	x := t.pos + off
	t.grow(x)
	if c := t.cell[x]; c > unicode.MaxASCII {
		fmt.Fprintf(t.out, "%d", c)
	} else {
		fmt.Fprintf(t.out, "%c", c)
	}
}

func (t *Tape) Scan(off int) {
	x := t.pos + off
	t.grow(x)
	fmt.Fscanf(t.out, "%c", &t.cell[x])
}

func (t *Tape) IsZero() bool {
	return t.cell[t.pos] == 0
}

func (t *Tape) Clear(off int) {
	x := t.pos + off
	t.grow(x)
	t.cell[x] = 0
}

func (t *Tape) Mult(dst, arg, off int) {
	x := t.pos + off
	t.grow(x)
	v := t.cell[x]
	t.Move(dst)
	t.Add(int(v)*arg, off)
	t.Move(-dst)
}

func (t *Tape) Search(n int) {
	for !t.IsZero() {
		t.Move(n)
	}
}

func (t *Tape) String() string {
	return fmt.Sprint(t.cell)
}
