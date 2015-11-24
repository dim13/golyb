package main

import "os"

func exec(prog string) {
	t := NewStaticTape(os.Stdout)
	p := ParseString(prog)
	p = OptContract(p)
	p = OptLoops(p)
	p = OptOffset(p)
	Execute(p, t, false)
}

// Hello World!
const helloWorld = `++++++++++[>+++++++>++++++++++>+++>+<<<<-]
>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.`

func ExampleHelloWorld() {
	exec(helloWorld)
	// Output: Hello World!
}

// Prints 202
const numeric = `>+>+>+>+>++<[>[<+++>-]<<]>.`

// Numeric output
func ExampleNumeric() {
	exec(numeric)
	// Output: 202
}

// Goes to cell 30000 and reports from there with a #. (Verifies that the
// array is big enough.)
const faraway = `++++[>++++++<-]>[>+++++>+++++++<<-]>>++++<
[[>[[>>+<<-]<]>>>-]>-[>+>+<<-]>]+++++[>+++++++<<++>-]>.<<.`

func ExampleFarAway() {
	exec(faraway)
	// Output: #
}
