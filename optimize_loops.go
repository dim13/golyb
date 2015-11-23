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
