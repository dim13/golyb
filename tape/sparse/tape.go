package sparse

type Tape struct {
	data map[int]int
	pos  int
}

func New() *Tape {
	return &Tape{
		data: make(map[int]int),
	}
}

func (t *Tape) Move(off int) {
	t.pos += off
}

func (t *Tape) Read(off int) int {
	return t.data[t.pos+off]
}

func (t *Tape) Write(off, v int) {
	t.data[t.pos+off] = v
}
