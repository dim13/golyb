package main

import (
	"fmt"
	"unicode"
)

const size = 4096

type Tape struct {
	cell []int
	pos  int
}

func NewTape() *Tape {
	return &Tape{
		cell: make([]int, size),
		pos:  0,
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
		fmt.Printf("%d", c)
	} else {
		fmt.Printf("%c", c)
	}
}

func (t *Tape) Scan() {
	fmt.Scanf("%c", &t.cell[t.pos])
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
