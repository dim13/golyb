package golyb

import (
	"bytes"
	"fmt"
	"strings"
)

func (p Program) Dump() {
	p.dump(0)
}

func (p Program) dump(n int) {
	ind := strings.Repeat("    ", n)
	for _, c := range p {
		fmt.Printf("%s%v\n", ind, c)
		if c.Op == BNZ {
			c.Branch.dump(n + 1)
		}
	}
}

func (c Command) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s", c.Op)
	if c.Arg != 0 {
		fmt.Fprintf(buf, " %d", c.Arg)
	}
	if c.Off != 0 {
		fmt.Fprintf(buf, " @%d", c.Off)
	}
	if c.Dst != 0 {
		fmt.Fprintf(buf, " → @%d", c.Dst)
	}
	return buf.String()
}
