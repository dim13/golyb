package main

func Optimize(p Program) Program {
	p = contract(p)
	p = loops(p)
	p = offset(p)
	return p
}
