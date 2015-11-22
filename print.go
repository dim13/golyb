package main

import "fmt"

func printArg(n int, more, less string) string {
	var s string
	for i := 0; i < n; i++ {
		s += more
	}
	for i := 0; i > n; i-- {
		s += less
	}
	return s
}

func (p Program) String() string {
	var s string
	for _, cmd := range p {
		s += fmt.Sprint(cmd)
	}
	return s
}

func (c Command) String() string {
	var s string
	switch c.Op {
	case Add:
		s = printArg(c.Arg, "+", "-")
	case Move:
		s = printArg(c.Arg, ">", "<")
	case Print:
		s = "."
	case Scan:
		s = ","
	case BNZ:
		s = "[" + fmt.Sprint(c.Branch) + "]"
	case Clear:
		s = "[-]"
	case Mult:
		s = "[-" + printArg(c.Off, ">", "<") +
			printArg(c.Arg, "+", "-") +
			printArg(-c.Off, ">", "<") + "]"
	}
	return s
}
