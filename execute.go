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
	for _, cmd := range p {
		switch cmd.Op {
		case Add:
			t.Add(cmd.Arg, cmd.Off)
		case Move:
			t.Move(cmd.Arg)
		case Print:
			t.Print(cmd.Off)
		case Scan:
			t.Scan(cmd.Off)
		case Loop:
			for !t.IsZero() {
				cmd.Branch.Execute(t)
			}
		case Clear:
			t.Clear(cmd.Off)
		case Mult:
			t.Mult(cmd.Dst, cmd.Arg, cmd.Off)
		case Search:
			t.Search(cmd.Arg)
		default:
			panic("unknown opcode")
		}
	}
}
