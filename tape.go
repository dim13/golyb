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
