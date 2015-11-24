package main

type Storage interface {
	Add(int, int)
	Move(int)
	Print(int)
	Scan(int)
	IsZero() bool
	Clear(int)
	Mult(int, int, int)
	Search(int)
	Dump() ([]int, int)
}

func Execute(p Program, s Storage) {
	for _, cmd := range p {
		switch cmd.Op {
		case Add:
			s.Add(cmd.Arg, cmd.Off)
		case Move:
			s.Move(cmd.Arg)
		case Print:
			s.Print(cmd.Off)
		case Scan:
			s.Scan(cmd.Off)
		case BNZ:
			for !s.IsZero() {
				Execute(cmd.Branch, s)
			}
		case Clear:
			s.Clear(cmd.Off)
		case Mult:
			s.Mult(cmd.Dst, cmd.Arg, cmd.Off)
		case Search:
			s.Search(cmd.Arg)
		default:
			panic("unknown opcode")
		}
	}
}
