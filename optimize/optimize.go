package optimize

import (
	. "github.com/dim13/golyb"
	"github.com/dim13/golyb/optimize/contract"
	"github.com/dim13/golyb/optimize/loops"
	"github.com/dim13/golyb/optimize/offset"
)

var defaultOpts = []func(Program) Program{
	contract.Optimize,
	loops.Optimize,
	offset.Optimize,
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
