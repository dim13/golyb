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

func NewFiniteTape(out io.ReadWriter) *FiniteTape {
	return &FiniteTape{
		cell: make([]int, 32768),
		pos:  1024,
		out:  out,
	}
}

func (t *FiniteTape) Move(n int) {
	t.pos += n
}

func (t *FiniteTape) Add(n int) {
	t.cell[t.pos] += n
}

func (t *FiniteTape) Print() {
	format := "%c"
	if t.cell[t.pos] > unicode.MaxASCII {
		format = "%d"
	}
	fmt.Fprintf(t.out, format, t.cell[t.pos])
}

func (t *FiniteTape) Scan() {
	fmt.Fscanf(t.out, "%c", &t.cell[t.pos])
}

func (t *FiniteTape) IsZero() bool {
	return t.cell[t.pos] == 0
}

func (t *FiniteTape) Clear() {
	t.cell[t.pos] = 0
}

func (t *FiniteTape) Mult(off, arg int) {
	t.cell[t.pos+off] += t.cell[t.pos] * arg
}

func (t *FiniteTape) Search(n int) {
	for !t.IsZero() {
		t.Move(n)
	}
}
