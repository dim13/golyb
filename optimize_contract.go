package main

func scan(p Program) (Command, int) {
	n := 1
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

func contract(p Program) Program {
	var o Program
	for i := 0; i < len(p); i++ {
		switch cmd := p[i]; cmd.Op {
		case Add, Move:
			cmd, n := scan(p[i:])
			o = append(o, cmd)
			i += n - 1
		default:
			cmd.Branch = contract(cmd.Branch)
			o = append(o, cmd)
		}
	}
	return o
}
