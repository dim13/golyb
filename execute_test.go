package main

import "os"

func optParse(prog string) Program {
	p := ParseString(prog)
	p = OptContract(p)
	p = OptLoops(p)
	p = OptOffset(p)
	return p
}

// Hello World!
const helloWorld = `++++++++++[>+++++++>++++++++++>+++>+<<<<-]
>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.`

func ExampleHelloWorld() {
	Execute(optParse(helloWorld), NewStaticTape(os.Stdout))
	// Output: Hello World!
}

// Prints 202
const numeric = `>+>+>+>+>++<[>[<+++>-]<<]>.`

// Numeric output
func ExampleNumeric() {
	Execute(optParse(numeric), NewStaticTape(os.Stdout))
	// Output: 202
}

// Goes to cell 30000 and reports from there with a #. (Verifies that the
// array is big enough.)
const faraway = `++++[>++++++<-]>[>+++++>+++++++<<-]>>++++<
[[>[[>>+<<-]<]>>>-]>-[>+>+<<-]>]+++++[>+++++++<<++>-]>.<<.`

func ExampleFarAway() {
	Execute(optParse(faraway), NewStaticTape(os.Stdout))
	// Output: #
}
