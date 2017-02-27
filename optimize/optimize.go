package optimize

import . "github.com/dim13/golyb"

func Optimize(p Program) Program {
	p = Contract(p)
	p = Loops(p)
	p = Offset(p)
	return p
}
