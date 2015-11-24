package main

func scan(p Program) (Command, int) {
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

func OptContract(p Program) Program {
	var o Program
	for i := 0; i < len(p); i++ {
		switch cmd := p[i]; cmd.Op {
		case Add, Move:
			cmd, n := scan(p[i:])
			if cmd.Arg != 0 {
				o = append(o, cmd)
			}
			i += n
		default:
			cmd.Branch = OptContract(cmd.Branch)
			o = append(o, cmd)
		}
	}
	return o
}
