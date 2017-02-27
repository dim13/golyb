package optimize

import "github.com/dim13/golyb"

func Offset(p golyb.Program) golyb.Program {
	var o golyb.Program
	// [>>>?<<<] for Add, Print, Scan, Clear, Mult
	// not for Move, BNZ, Search
	for i := 0; i < len(p); i++ {
		switch b := p[i:]; {
		case len(b) >= 3 &&
			b[0].Op == golyb.Move &&
			(b[1].Op == golyb.Add || b[1].Op == golyb.Print || b[1].Op == golyb.Scan || b[1].Op == golyb.Clear) &&
			b[2].Op == golyb.Move:
			o = append(o, golyb.Command{
				Op:  b[1].Op,
				Arg: b[1].Arg,
				Off: b[0].Arg,
			})
			// push back combined move
			m := b[0].Arg + b[2].Arg
			p[i+2] = golyb.Command{
				Op:  golyb.Move,
				Arg: m,
			}
			if m == 0 {
				i += 2
			} else {
				i += 1
			}
		case len(b) >= 4 &&
			b[0].Op == golyb.Move &&
			b[1].Op == golyb.Mult &&
			b[2].Op == golyb.Clear &&
			b[3].Op == golyb.Move:
			o = append(o, golyb.Command{
				Op:  b[1].Op,
				Dst: b[1].Dst,
				Arg: b[1].Arg,
				Off: b[0].Arg,
			})
			o = append(o, golyb.Command{
				Op:  b[2].Op,
				Arg: b[2].Arg,
				Off: b[0].Arg,
			})
			// push back combined move
			m := b[0].Arg + b[3].Arg
			p[i+3] = golyb.Command{
				Op:  golyb.Move,
				Arg: b[0].Arg + b[3].Arg,
			}
			if m == 0 {
				i += 3
			} else {
				i += 2
			}
		default:
			p[i].Branch = Offset(p[i].Branch)
			o = append(o, p[i])
		}
	}
	return o
}
