package optimize

import "github.com/dim13/golyb"

func scan(p golyb.Program) (golyb.Command, int) {
	n := 0
	c := p[0]
	for _, cmd := range p[1:] {
		if cmd.Op == c.Op {
			c.Arg += cmd.Arg
			n++
		} else {
			break
		}
	}
	return c, n
}

func Contract(p golyb.Program) golyb.Program {
	var o golyb.Program
	for i := 0; i < len(p); i++ {
		switch cmd := p[i]; cmd.Op {
		case golyb.Add, golyb.Move:
			cmd, n := scan(p[i:])
			if cmd.Arg != 0 {
				o = append(o, cmd)
			}
			i += n
		default:
			cmd.Branch = Contract(cmd.Branch)
			o = append(o, cmd)
		}
	}
	return o
}