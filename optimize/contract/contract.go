package contract

import . "github.com/dim13/golyb"

func scan(p Program) (Command, int) {
	n := 0
	c := p[0]
	for _, cmd := range p[1:] {
		if cmd.Op == c.Op {
			c.Arg += cmd.Arg
			c.Off += cmd.Off
			n++
		} else {
			break
		}
	}
	return c, n
}

func Optimize(p Program) (out Program) {
	for i := 0; i < len(p); i++ {
		switch cmd := p[i]; cmd.Op {
		case Add, Move:
			cmd, n := scan(p[i:])
			if cmd.Arg != 0 || cmd.Off != 0 {
				out = append(out, cmd)
			}
			i += n
		default:
			cmd.Branch = Optimize(cmd.Branch)
			out = append(out, cmd)
		}
	}
	return out
}
