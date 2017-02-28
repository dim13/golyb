package golyb

import (
	"bytes"
	"fmt"
	"io"
)

const indLevel = 4

func (p Program) String() string {
	buf := new(bytes.Buffer)
	p.dump(buf, 0)
	return buf.String()
}

func (p Program) dump(w io.Writer, n int) {
	for _, c := range p {
		fmt.Fprintf(w, "%*s%v\n", n*indLevel, "", c)
		if c.Op == Loop {
			c.Branch.dump(w, n+1)
		}
	}
}

func (c Command) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s", c.Op)
	if c.Arg != 0 {
		fmt.Fprintf(buf, " %+d", c.Arg)
	}
	if c.Off != 0 {
		fmt.Fprintf(buf, " @%+d", c.Off)
	}
	if c.Dst != 0 {
		fmt.Fprintf(buf, " → @%+d", c.Dst)
	}
	return buf.String()
}
