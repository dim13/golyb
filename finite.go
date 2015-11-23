package main

import (
	"fmt"
	"io"
	"unicode"
)

type FiniteTape struct {
	cell []int
	pos  int
	out  io.ReadWriter
}

func NewFiniteTape(out io.ReadWriter) Storage {
	return &FiniteTape{
		cell: make([]int, 65536),
		pos:  4096, // left some space on LHS
		out:  out,
	}
}

func (t *FiniteTape) Move(n int) {
	t.pos += n
}

func (t *FiniteTape) Add(n, off int) {
	t.cell[t.pos+off] += n
}

func (t *FiniteTape) Print(off int) {
	format := "%c"
	v := t.cell[t.pos+off]
	if v > unicode.MaxASCII {
		format = "%d"
	}
	fmt.Fprintf(t.out, format, v)
}

func (t *FiniteTape) Scan(off int) {
	fmt.Fscanf(t.out, "%c", &t.cell[t.pos+off])
}

func (t *FiniteTape) IsZero() bool {
	return t.cell[t.pos] == 0
}

func (t *FiniteTape) Clear(off int) {
	t.cell[t.pos+off] = 0
}

func (t *FiniteTape) Mult(dst, arg, off int) {
	t.cell[t.pos+dst+off] += t.cell[t.pos+off] * arg
}

func (t *FiniteTape) Search(n int) {
	for t.cell[t.pos] != 0 {
		t.pos += n
	}
}
