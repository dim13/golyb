// Go f*ck your self -- a BrainF*ck interpreter
package main

func match(p Program, q Program) bool {
	if len(p) != len(q) {
		return false
	}
	for i := range p {
		if p[i].Op != q[i].Op {
			return false
		}
		if q[i].Arg != 0 && p[i].Arg != q[i].Arg {
			return false
		}
	}
	return true
}

func clearLoop(p Program) Program {
	for i, cmd := range p {
		if cmd.Op == BNZ {
			b := cmd.Branch
			if match(b, Program{Command{Op: Add, Arg: 0}}) {
				p[i] = Command{Op: Clear}
			} else {
				p[i].Branch = clearLoop(b)
			}
		}
	}
	return p
}

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
		cmd := p[i]
		switch cmd.Op {
		case BNZ:
			cmd.Branch = contract(cmd.Branch)
		case Add, Move:
			var n int
			cmd, n = scan(p[i:])
			i += n - 1
		}
		o = append(o, cmd)
	}
	return o
}

func Optimize(p Program) Program {
	o := contract(p)
	return clearLoop(o)
}
