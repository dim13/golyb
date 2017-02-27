package optimize

import "github.com/dim13/golyb"

func match(p golyb.Program, q golyb.Program) bool {
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

func Loops(p golyb.Program) golyb.Program {
	var o golyb.Program
	for _, cmd := range p {
		switch cmd.Op {
		case golyb.Branch:
			switch b := cmd.Branch; {
			// [-] or [+]
			case len(b) == 1 && match(b, golyb.Program{
				golyb.Command{Op: golyb.Add}}):
				o = append(o, golyb.Command{Op: golyb.Clear})
			// [>] or [<]
			case len(b) == 1 && match(b, golyb.Program{
				golyb.Command{Op: golyb.Move}}):
				o = append(o, golyb.Command{Op: golyb.Search, Arg: b[0].Arg})
			// [->+<]
			case len(b) == 4 && match(b, golyb.Program{
				golyb.Command{Op: golyb.Add, Arg: -1},
				golyb.Command{Op: golyb.Move},
				golyb.Command{Op: golyb.Add},
				golyb.Command{Op: golyb.Move, Arg: -b[1].Arg}}):
				o = append(o, golyb.Command{
					Op:  golyb.Mult,
					Dst: b[1].Arg,
					Arg: b[2].Arg,
				})
				o = append(o, golyb.Command{Op: golyb.Clear})
			// [>+<-]
			case len(b) == 4 && match(b, golyb.Program{
				golyb.Command{Op: golyb.Move},
				golyb.Command{Op: golyb.Add},
				golyb.Command{Op: golyb.Move, Arg: -b[0].Arg},
				golyb.Command{Op: golyb.Add, Arg: -1}}):
				o = append(o, golyb.Command{
					Op:  golyb.Mult,
					Dst: b[0].Arg,
					Arg: b[1].Arg,
				})
				o = append(o, golyb.Command{Op: golyb.Clear})
			// [->+>+<<] or [->>+<+<]
			case len(b) == 6 && match(b, golyb.Program{
				golyb.Command{Op: golyb.Add, Arg: -1},
				golyb.Command{Op: golyb.Move},
				golyb.Command{Op: golyb.Add},
				golyb.Command{Op: golyb.Move},
				golyb.Command{Op: golyb.Add},
				golyb.Command{Op: golyb.Move, Arg: -b[1].Arg - b[3].Arg}}):
				o = append(o, golyb.Command{
					Op:  golyb.Mult,
					Dst: b[1].Arg,
					Arg: b[2].Arg,
				})
				o = append(o, golyb.Command{
					Op:  golyb.Mult,
					Dst: b[1].Arg + b[3].Arg,
					Arg: b[4].Arg,
				})
				o = append(o, golyb.Command{Op: golyb.Clear})
			// [>+>+<<-] or [>>+<+<-]
			case len(b) == 6 && match(b, golyb.Program{
				golyb.Command{Op: golyb.Move},
				golyb.Command{Op: golyb.Add},
				golyb.Command{Op: golyb.Move},
				golyb.Command{Op: golyb.Add},
				golyb.Command{Op: golyb.Move, Arg: -b[0].Arg - b[2].Arg},
				golyb.Command{Op: golyb.Add, Arg: -1}}):
				o = append(o, golyb.Command{
					Op:  golyb.Mult,
					Dst: b[0].Arg,
					Arg: b[1].Arg,
				})
				o = append(o, golyb.Command{
					Op:  golyb.Mult,
					Dst: b[0].Arg + b[2].Arg,
					Arg: b[3].Arg,
				})
				o = append(o, golyb.Command{Op: golyb.Clear})
			case len(b) == 0:
				continue
			default:
				cmd.Branch = Loops(cmd.Branch)
				o = append(o, cmd)
			}
		default:
			// passthrough
			o = append(o, cmd)
		}
	}
	return o
}
