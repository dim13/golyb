package golyb

type Tape interface {
	Add(arg int, off int)
	Move(off int)
	Print(off int)
	Scan(off int)
	IsZero() bool
	Clear(off int)
	Mult(arg int, off int, dst int)
	Search(off int)
}

func (p Program) Execute(t Tape) {
	for i := range p {
		switch p[i].Op {
		case Add:
			t.Add(p[i].Arg, p[i].Off)
		case Move:
			t.Move(p[i].Off)
		case Print:
			t.Print(p[i].Off)
		case Scan:
			t.Scan(p[i].Off)
		case Loop:
			for !t.IsZero() {
				p[i].Branch.Execute(t)
			}
		case Clear:
			t.Clear(p[i].Off)
		case Mult:
			t.Mult(p[i].Arg, p[i].Off, p[i].Dst)
		case Search:
			t.Search(p[i].Off)
		default:
			panic("unknown opcode")
		}
	}
}
