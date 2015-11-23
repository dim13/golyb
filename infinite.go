package main

import (
	"fmt"
	"io"
	"unicode"
)

const size = 4096

type InfiniteTape struct {
	cell []int
	pos  int
	out  io.ReadWriter
}

func NewInfiniteTape(out io.ReadWriter) Storage {
	return &InfiniteTape{
		cell: make([]int, size),
		pos:  0,
		out:  out,
	}
}

func (t *InfiniteTape) Move(n int) {
	t.pos += n
	if t.pos >= len(t.cell) {
		t.cell = append(t.cell, make([]int, size)...)
	} else if t.pos < 0 {
		t.cell = append(make([]int, size), t.cell...)
		t.pos += size
	}
}

func (t *InfiniteTape) Add(n, off int) {
	t.cell[t.pos+off] += n
}

func (t *InfiniteTape) Print(off int) {
	if c := t.cell[t.pos+off]; c > unicode.MaxASCII {
		fmt.Fprintf(t.out, "%d", c)
	} else {
		fmt.Fprintf(t.out, "%c", c)
	}
}

func (t *InfiniteTape) Scan(off int) {
	fmt.Fscanf(t.out, "%c", &t.cell[t.pos+off])
}

func (t *InfiniteTape) IsZero() bool {
	return t.cell[t.pos] == 0
}

func (t *InfiniteTape) Clear(off int) {
	t.cell[t.pos+off] = 0
}

func (t *InfiniteTape) Mult(dst, arg, off int) {
	v := t.cell[t.pos+off]
	t.Move(dst)
	t.Add(v*arg, off)
	t.Move(-dst)
	//t.Clear() // inserted by optimization
}

func (t *InfiniteTape) Search(n int) {
	for !t.IsZero() {
		t.Move(n)
	}
}
