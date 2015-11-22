package main

type Storage interface {
	Add(int)
	Move(int)
	Scan()
	Print()
	IsZero() bool
	Clear()
	Mult(int, int)
}

func Execute(p Program, s Storage) {
	for _, cmd := range p {
		switch cmd.Op {
		case Add:
			s.Add(cmd.Arg)
		case Move:
			s.Move(cmd.Arg)
		case Print:
			s.Print()
		case Scan:
			s.Scan()
		case BNZ:
			for !s.IsZero() {
				Execute(cmd.Branch, s)
			}
		case Clear:
			s.Clear()
		case Mult:
			s.Mult(cmd.Off, cmd.Arg)
		}
	}
}
