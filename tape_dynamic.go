package main

import (
	"fmt"
	"io"
	"unicode"
)

type DynamicTape struct {
	cell []int
	pos  int
	out  io.ReadWriter
}

const chunkSize = 1024

func NewDynamicTape(out io.ReadWriter) Storage {
	return &DynamicTape{
		cell: make([]int, chunkSize),
		pos:  0,
		out:  out,
	}
}

func (t *DynamicTape) grow(pos int) {
	if pos >= len(t.cell) {
		t.cell = append(t.cell, make([]int, chunkSize)...)
	}
	if pos < 0 {
		t.cell = append(make([]int, chunkSize), t.cell...)
		t.pos += chunkSize
	}
}

func (t *DynamicTape) Move(n int) {
	t.pos += n
	t.grow(t.pos)
}

func (t *DynamicTape) Add(n, off int) {
	x := t.pos + off
	t.grow(x)
	t.cell[x] += n
}

func (t *DynamicTape) Print(off int) {
	x := t.pos + off
	t.grow(x)
	if c := t.cell[x]; c > unicode.MaxASCII {
		fmt.Fprintf(t.out, "%d", c)
	} else {
		fmt.Fprintf(t.out, "%c", c)
	}
}

func (t *DynamicTape) Scan(off int) {
	x := t.pos + off
	t.grow(x)
	fmt.Fscanf(t.out, "%c", &t.cell[x])
}

func (t *DynamicTape) IsZero() bool {
	return t.cell[t.pos] == 0
}

func (t *DynamicTape) Clear(off int) {
	x := t.pos + off
	t.grow(x)
	t.cell[x] = 0
}

func (t *DynamicTape) Mult(dst, arg, off int) {
	x := t.pos + off
	t.grow(x)
	v := t.cell[x]
	t.Move(dst)
	t.Add(v*arg, off)
	t.Move(-dst)
	//t.Clear() // inserted by optimization
}

func (t *DynamicTape) Search(n int) {
	for !t.IsZero() {
		t.Move(n)
	}
}

func (t *DynamicTape) Dump() ([]int, int) {
	return t.cell, t.pos
}
