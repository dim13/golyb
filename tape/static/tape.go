package static

type Tape struct {
	cell []int
	pos  int
}

const (
	tapeSize = 30 * 1024
	margin   = 1024
)

func New() *Tape {
	return &Tape{
		cell: make([]int, tapeSize+2*margin),
		pos:  margin,
	}
}

func (t *Tape) Move(off int) {
	t.pos += off
}

func (t *Tape) Read(off int) int {
	return t.cell[t.pos+off]
}

func (t *Tape) Write(off, v int) {
	t.cell[t.pos+off] = v
}
