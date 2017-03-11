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
	var cmd Command
	for i := range p {
		cmd = p[i]
		switch cmd.Op {
		case Add:
			t.Add(cmd.Arg, cmd.Off)
		case Move:
			t.Move(cmd.Off)
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
			t.Mult(cmd.Arg, cmd.Off, cmd.Dst)
		case Search:
			t.Search(cmd.Off)
		default:
			panic("unknown opcode")
		}
	}
}
