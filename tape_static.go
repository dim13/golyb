package main

import (
	"fmt"
	"io"
	"unicode"
)

type StaticTape struct {
	cell []int
	pos  int
	out  io.ReadWriter
}

func NewStaticTape(out io.ReadWriter) Storage {
	return &StaticTape{
		cell: make([]int, 32768),
		pos:  1024, // left some space on LHS
		out:  out,
	}
}

func (t *StaticTape) Move(n int) {
	t.pos += n
}

func (t *StaticTape) Add(n, off int) {
	t.cell[t.pos+off] += n
}

func (t *StaticTape) Print(off int) {
	format := "%c"
	v := t.cell[t.pos+off]
	if v > unicode.MaxASCII {
		format = "%d"
	}
	fmt.Fprintf(t.out, format, v)
}

func (t *StaticTape) Scan(off int) {
	fmt.Fscanf(t.out, "%c", &t.cell[t.pos+off])
}

func (t *StaticTape) IsZero() bool {
	return t.cell[t.pos] == 0
}

func (t *StaticTape) Clear(off int) {
	t.cell[t.pos+off] = 0
}

func (t *StaticTape) Mult(dst, arg, off int) {
	t.cell[t.pos+dst+off] += t.cell[t.pos+off] * arg
}

func (t *StaticTape) Search(n int) {
	for t.cell[t.pos] != 0 {
		t.pos += n
	}
}

func (t *StaticTape) Dump() ([]int, int) {
	return t.cell, t.pos
}
