package main

func Optimize(p Program) Program {
	p = OptContract(p)
	p = OptLoops(p)
	p = OptOffset(p)
	return p
}
