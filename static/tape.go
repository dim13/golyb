package static

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

const (
	tapeSize = 30 * 1024
	margin   = 1024
)

func New(r io.Reader, w io.Writer) *Tape {
	if r == nil {
		r = os.Stdin
	}
	if w == nil {
		w = os.Stdout
	}
	return &Tape{
		cell: make([]int, tapeSize+2*margin),
		pos:  margin,
		r:    r,
		w:    w,
	}
}

func (t *Tape) Move(off int) {
	t.pos += off
}

func (t *Tape) Add(arg, off int) {
	t.cell[t.pos+off] += arg
}

func (t *Tape) Print(off int) {
	if c := t.cell[t.pos+off]; c > unicode.MaxASCII {
		fmt.Fprintf(t.w, "%d", c)
	} else {
		fmt.Fprintf(t.w, "%c", c)
	}
}

func (t *Tape) Scan(off int) {
	fmt.Fscanf(t.r, "%c", &t.cell[t.pos+off])
}

func (t *Tape) IsZero() bool {
	return t.cell[t.pos] == 0
}

func (t *Tape) Clear(off int) {
	t.cell[t.pos+off] = 0
}

func (t *Tape) Mult(arg, off, dst int) {
	t.cell[t.pos+dst+off] += t.cell[t.pos+off] * arg
}

func (t *Tape) Search(off int) {
	for t.cell[t.pos] != 0 {
		t.pos += off
	}
}

func (t *Tape) String() string {
	return fmt.Sprint(t.cell)
}
