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
			case len(b) == 1 && match(b, Program{Command{Op: Add}}):
				o = append(o, Command{Op: Clear})
			// [>] or [<]
			case len(b) == 1 && match(b, Program{Command{Op: Move}}):
				o = append(o, Command{Op: Search, Arg: b[0].Arg})
			// [->+<]
			case len(b) == 4 && match(b, Program{
				Command{Op: Add, Arg: -1},
				Command{Op: Move},
				Command{Op: Add},
				Command{Op: Move, Arg: -b[1].Arg}}):
				o = append(o, Command{
					Op:  Mult,
					Dst: b[1].Arg,
					Arg: b[2].Arg,
				})
				o = append(o, Command{Op: Clear})
			// [>+<-]
			case len(b) == 4 && match(b, Program{
				Command{Op: Move},
				Command{Op: Add},
				Command{Op: Move, Arg: -b[0].Arg},
				Command{Op: Add, Arg: -1}}):
				o = append(o, Command{
					Op:  Mult,
					Dst: b[0].Arg,
					Arg: b[1].Arg,
				})
				o = append(o, Command{Op: Clear})
			// [->+>+<<] or [->>+<+<]
			case len(b) == 6 && match(b, Program{
				Command{Op: Add, Arg: -1},
				Command{Op: Move},
				Command{Op: Add},
				Command{Op: Move},
				Command{Op: Add},
				Command{Op: Move, Arg: -b[1].Arg - b[3].Arg}}):
				o = append(o, Command{
					Op:  Mult,
					Dst: b[1].Arg,
					Arg: b[2].Arg,
				})
				o = append(o, Command{
					Op:  Mult,
					Dst: b[1].Arg + b[3].Arg,
					Arg: b[4].Arg,
				})
				o = append(o, Command{Op: Clear})
			// [>+>+<<-] or [>>+<+<-]
			case len(b) == 6 && match(b, Program{
				Command{Op: Move},
				Command{Op: Add},
				Command{Op: Move},
				Command{Op: Add},
				Command{Op: Move, Arg: -b[0].Arg - b[2].Arg},
				Command{Op: Add, Arg: -1}}):
				o = append(o, Command{
					Op:  Mult,
					Dst: b[0].Arg,
					Arg: b[1].Arg,
				})
				o = append(o, Command{
					Op:  Mult,
					Dst: b[0].Arg + b[2].Arg,
					Arg: b[3].Arg,
				})
				o = append(o, Command{Op: Clear})
			default:
				cmd.Branch = loops(cmd.Branch)
				o = append(o, cmd)
			}
		default:
			// passthrough
			o = append(o, cmd)
		}
	}
	return o
}

func offset(p Program) Program {
	var o Program
	var lastmove Command
	// [>>>?<<<] for Add, Print, Scan, Clear, Mult
	// not for Move, BNZ, Search
	for i := 0; i < len(p); i++ {
		switch b := p[i:]; {
		case len(b) >= 3 &&
			b[0].Op == Move &&
			(b[1].Op == Add || b[1].Op == Print || b[1].Op == Scan || b[1].Op == Clear) &&
			b[2].Op == Move:
			o = append(o, Command{
				Op:  b[1].Op,
				Arg: b[1].Arg,
				Off: b[0].Arg,
			})
			lastmove = Command{
				Op:  Move,
				Arg: b[0].Arg + b[2].Arg,
			}
			p[i+2] = lastmove
			i += 1
		case len(b) >= 4 &&
			b[0].Op == Move &&
			b[1].Op == Mult &&
			b[2].Op == Clear &&
			b[3].Op == Move:
			o = append(o, Command{
				Op:  b[1].Op,
				Dst: b[1].Dst,
				Arg: b[1].Arg,
				Off: b[0].Arg,
			})
			o = append(o, Command{
				Op:  b[2].Op,
				Arg: b[2].Arg,
				Off: b[0].Arg,
			})
			lastmove = Command{
				Op:  Move,
				Arg: b[0].Arg + b[3].Arg,
			}
			p[i+3] = lastmove
			i += 2
		default:
			p[i].Branch = offset(p[i].Branch)
			o = append(o, p[i])
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

func Optimize(p Program) Program {
	p = contract(p)
	p = loops(p)
	p = offset(p)
	return p
}
