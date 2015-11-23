package main

import "log"

type Storage interface {
	Add(int)
	Move(int)
	Print()
	Scan()
	IsZero() bool
	Clear()
	Mult(int, int)
	Search(int)
}

func Execute(p Program, s Storage) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
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
			s.Mult(cmd.Dst, cmd.Arg)
		case Search:
			s.Search(cmd.Arg)
		}
	}
}
