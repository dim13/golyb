package sparse

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

type Tape struct {
	data map[int]int
	pos  int
	r    io.Reader
	w    io.Writer
}

func New(r io.Reader, w io.Writer) *Tape {
	if r == nil {
		r = os.Stdin
	}
	if w == nil {
		w = os.Stdout
	}
	return &Tape{
		data: make(map[int]int),
		r:    r,
		w:    w,
	}
}

func (t *Tape) Add(arg int, off int) {
	t.data[t.pos+off] += arg
}

func (t *Tape) Move(off int) {
	t.pos += off
}

func (t *Tape) Print(off int) {
	if c := t.data[t.pos+off]; c > unicode.MaxASCII {
		fmt.Fprintf(t.w, "%d", c)
	} else {
		fmt.Fprintf(t.w, "%c", c)
	}
}

func (t *Tape) Scan(off int) {
	var arg int
	fmt.Fscanf(t.r, "%c", &arg)
	t.data[t.pos+off] = arg
}

func (t *Tape) IsZero() bool {
	return t.data[t.pos] == 0
}

func (t *Tape) Clear(off int) {
	delete(t.data, t.pos+off)
}

func (t *Tape) Mult(arg int, off int, dst int) {
	t.data[t.pos+dst+off] += t.data[t.pos+off] * arg
}

func (t *Tape) Search(off int) {
	for t.data[t.pos] != 0 {
		t.pos += off
	}
}
