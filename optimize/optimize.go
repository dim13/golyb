package optimize

import "github.com/dim13/golyb"

func Optimize(p golyb.Program) golyb.Program {
	p = Contract(p)
	p = Loops(p)
	p = Offset(p)
	return p
}
