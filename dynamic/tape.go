package dynamic

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

type Tape struct {
	cell []int
	pos  int
	r    io.Reader
	w    io.Writer
}

const chunkSize = 1024

func New(r io.Reader, w io.Writer) *Tape {
	if r == nil {
		r = os.Stdin
	}
	if w == nil {
		w = os.Stdout
	}
	return &Tape{
		cell: make([]int, chunkSize),
		pos:  0,
		r:    r,
		w:    w,
	}
}

func (t *Tape) grow(pos int) {
	if pos >= len(t.cell) {
		t.cell = append(t.cell, make([]int, chunkSize)...)
	}
	if pos < 0 {
		t.cell = append(make([]int, chunkSize), t.cell...)
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
	t.cell[x] += n
}

func (t *Tape) Print(off int) {
	x := t.pos + off
	t.grow(x)
	if c := t.cell[x]; c > unicode.MaxASCII {
		fmt.Fprintf(t.w, "%d", c)
	} else {
		fmt.Fprintf(t.w, "%c", c)
	}
}

func (t *Tape) Scan(off int) {
	x := t.pos + off
	t.grow(x)
	fmt.Fscanf(t.r, "%c", &t.cell[x])
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
