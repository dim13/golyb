package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

const size = 4096

type Tape struct {
	cell []int
	pos  int
	out  io.ReadWriter
}

func output(out string) *os.File {
	if out == "stdout" {
		return os.Stdout
	}
	file, _ := os.Create(out)
	return file
}

func NewTape(out string) *Tape {
	return &Tape{
		cell: make([]int, size),
		pos:  0,
		out:  output(out),
	}
}

func (t *Tape) Move(n int) {
	t.pos += n
	if t.pos >= len(t.cell) {
		t.cell = append(t.cell, make([]int, size)...)
	} else if t.pos < 0 {
		t.cell = append(make([]int, size), t.cell...)
		t.pos += size
	}
}

func (t *Tape) Add(n int) {
	t.cell[t.pos] += n
}

func (t *Tape) Print() {
	if c := t.cell[t.pos]; c > unicode.MaxASCII {
		fmt.Fprintf(t.out, "%d", c)
	} else {
		fmt.Fprintf(t.out, "%c", c)
	}
}

func (t *Tape) Scan() {
	fmt.Fscanf(t.out, "%c", &t.cell[t.pos])
}

func (t *Tape) IsZero() bool {
	return t.cell[t.pos] == 0
}

func (t *Tape) Clear() {
	t.cell[t.pos] = 0
}

func (t *Tape) Mult(off, arg int) {
	v := t.cell[t.pos]
	t.Move(off)
	t.Add(v * arg)
	t.Move(-off)
	t.Clear()
}

func (t *Tape) Search(n int) {
	for !t.IsZero() {
		t.Move(n)
	}
}
