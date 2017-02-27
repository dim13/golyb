package static

import (
	"fmt"
	"io"
	"unicode"

	"github.com/dim13/golyb"
)

type Tape struct {
	cell []int
	pos  int
	out  io.ReadWriter
}

const (
	tapeSize = 30 * 1024
	margin   = 1024
)

func NewTape(out io.ReadWriter) golyb.Storage {
	return &Tape{
		cell: make([]int, tapeSize+2*margin),
		pos:  margin, // left some space on LHS
		out:  out,
	}
}

func (t *Tape) Move(n int) {
	t.pos += n
}

func (t *Tape) Add(n, off int) {
	t.cell[t.pos+off] += n
}

func (t *Tape) Print(off int) {
	if c := t.cell[t.pos+off]; c > unicode.MaxASCII {
		fmt.Fprintf(t.out, "%d", c)
	} else {
		fmt.Fprintf(t.out, "%c", c)
	}
}

func (t *Tape) Scan(off int) {
	fmt.Fscanf(t.out, "%c", &t.cell[t.pos+off])
}

func (t *Tape) IsZero() bool {
	return t.cell[t.pos] == 0
}

func (t *Tape) Clear(off int) {
	t.cell[t.pos+off] = 0
}

func (t *Tape) Mult(dst, arg, off int) {
	t.cell[t.pos+dst+off] += t.cell[t.pos+off] * arg
}

func (t *Tape) Search(n int) {
	for t.cell[t.pos] != 0 {
		t.pos += n
	}
}

func (t *Tape) String() string {
	return fmt.Sprint(t.cell)
}
