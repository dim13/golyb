package golyb

type Tape interface {
	Add(int, int)
	Move(int)
	Print(int)
	Scan(int)
	IsZero() bool
	Clear(int)
	Mult(int, int, int)
	Search(int)
}

func (p Program) Execute(t Tape) {
	for i := 0; i < len(p); i++ {
		switch p[i].Op {
		case Add:
			t.Add(p[i].Arg, p[i].Off)
		case Move:
			t.Move(p[i].Arg)
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
			t.Search(p[i].Arg)
		default:
			panic("unknown opcode")
		}
	}
}
