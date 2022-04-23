package dynamic

type Tape struct {
	cell []int
	pos  int
}

const chunkSize = 1024

func New() *Tape {
	return &Tape{
		cell: make([]int, chunkSize),
		pos:  0,
	}
}

func (t *Tape) grow(pos int) {
	if pos >= len(t.cell) {
		t.cell = append(t.cell, make([]int, chunkSize)...)
	}
	if pos < 0 {
		t.cell = append(make([]int, chunkSize), t.cell...)
		t.pos += chunkSize
	}
}

func (t *Tape) Move(off int) {
	t.pos += off
	t.grow(t.pos)
}

func (t *Tape) Read(off int) int {
	x := t.pos + off
	t.grow(x)
	return t.cell[x]
}

func (t *Tape) Write(off, v int) {
	x := t.pos + off
	t.grow(x)
	t.cell[x] = v
}
