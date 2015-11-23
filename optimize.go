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

func loops(p Program) Program {
	var o Program
	for _, cmd := range p {
		switch cmd.Op {
		case BNZ:
			switch b := cmd.Branch; {
			// [-] or [+]
			case match(b, Program{Command{Op: Add}}):
				o = append(o, Command{Op: Clear})
			// [>] or [<]
			case match(b, Program{Command{Op: Move}}):
				o = append(o, Command{Op: Search, Arg: b[0].Arg})
			// [->+<]
			case len(b) == 4 && match(b, Program{
				Command{Op: Add, Arg: -1},
				Command{Op: Move},
				Command{Op: Add},
				Command{Op: Move, Arg: -b[1].Arg}}):
				o = append(o, Command{
					Op:  Mult,
					Off: b[1].Arg,
					Arg: b[2].Arg,
				})
			// [>+<-]
			case len(b) == 4 && match(b, Program{
				Command{Op: Move},
				Command{Op: Add},
				Command{Op: Move, Arg: -b[0].Arg},
				Command{Op: Add, Arg: -1}}):
				o = append(o, Command{
					Op:  Mult,
					Off: b[0].Arg,
					Arg: b[1].Arg,
				})
			// todo: [->+>+<<]
			// todo: [>+>+<<-]
			default:
				o = append(o, Command{Op: BNZ, Branch: loops(b)})
			}
		default:
			// passthrough
			o = append(o, cmd)
		}
	}
	return o
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
	return loops(o)
}
