package main

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
