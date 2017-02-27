package optimize

import . "github.com/dim13/golyb"

var defaultOpts = []func(Program) Program{
	Contract,
	Loops,
	Offset,
}

func All(p Program, opts ...func(Program) Program) Program {
	if opts == nil {
		opts = defaultOpts
	}
	for _, f := range opts {
		p = f(p)
	}
	return p
}
