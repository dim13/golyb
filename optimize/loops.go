package optimize

import . "github.com/dim13/golyb"

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
		if q[i].Off != 0 && p[i].Off != q[i].Off {
			return false
		}
	}
	return true
}

func Loops(p Program) (out Program) {
	for _, cmd := range p {
		if cmd.Op != Loop {
			// passthrough
			out = append(out, cmd)
			continue
		}
		switch b := cmd.Branch; {
		// [-] or [+]
		case len(b) == 1 && match(b, Program{Command{Op: Add}}):
			out = append(out, Command{Op: Clear})
		// [>] or [<]
		case len(b) == 1 && match(b, Program{Command{Op: Move}}):
			out = append(out, Command{Op: Search, Off: b[0].Off})
		// [->+<]
		case len(b) == 4 && match(b, Program{
			Command{Op: Add, Arg: -1},
			Command{Op: Move},
			Command{Op: Add},
			Command{Op: Move, Off: -b[1].Off}}):
			out = append(out, Command{
				Op:  Mult,
				Dst: b[1].Off,
				Arg: b[2].Arg,
			})
			out = append(out, Command{Op: Clear})
		// [>+<-]
		case len(b) == 4 && match(b, Program{
			Command{Op: Move},
			Command{Op: Add},
			Command{Op: Move, Off: -b[0].Off},
			Command{Op: Add, Arg: -1}}):
			out = append(out, Command{
				Op:  Mult,
				Dst: b[0].Off,
				Arg: b[1].Arg,
			})
			out = append(out, Command{Op: Clear})
		// [->+>+<<] or [->>+<+<]
		case len(b) == 6 && match(b, Program{
			Command{Op: Add, Arg: -1},
			Command{Op: Move},
			Command{Op: Add},
			Command{Op: Move},
			Command{Op: Add},
			Command{Op: Move, Off: -b[1].Off - b[3].Off}}):
			out = append(out, Command{
				Op:  Mult,
				Dst: b[1].Off,
				Arg: b[2].Arg,
			})
			out = append(out, Command{
				Op:  Mult,
				Dst: b[1].Off + b[3].Off,
				Arg: b[4].Arg,
			})
			out = append(out, Command{Op: Clear})
		// [>+>+<<-] or [>>+<+<-]
		case len(b) == 6 && match(b, Program{
			Command{Op: Move},
			Command{Op: Add},
			Command{Op: Move},
			Command{Op: Add},
			Command{Op: Move, Off: -b[0].Off - b[2].Off},
			Command{Op: Add, Arg: -1}}):
			out = append(out, Command{
				Op:  Mult,
				Dst: b[0].Off,
				Arg: b[1].Arg,
			})
			out = append(out, Command{
				Op:  Mult,
				Dst: b[0].Off + b[2].Off,
				Arg: b[3].Arg,
			})
			out = append(out, Command{Op: Clear})
		case len(b) == 0:
			continue
		default:
			cmd.Branch = Loops(cmd.Branch)
			out = append(out, cmd)
		}
	}
	return out
}
