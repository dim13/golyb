package main

import "os"

// Hello World!
const helloWorld = `++++++++++[>+++++++>++++++++++>+++>+<<<<-]
>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.`

func ExampleHelloWorld() {
	Execute(ParseStringOptimized(helloWorld), NewFiniteTape(os.Stdout))
	// Output: Hello World!
}

// Prints 202
const numeric = `>+>+>+>+>++<[>[<+++>-]<<]>.`

// Numeric output
func ExampleNumeric() {
	Execute(ParseStringOptimized(numeric), NewFiniteTape(os.Stdout))
	// Output: 202
}

// Goes to cell 30000 and reports from there with a #. (Verifies that the
// array is big enough.)
const faraway = `++++[>++++++<-]>[>+++++>+++++++<<-]>>++++<
[[>[[>>+<<-]<]>>>-]>-[>+>+<<-]>]+++++[>+++++++<<++>-]>.<<.`

func ExampleFarAway() {
	Execute(ParseStringOptimized(faraway), NewFiniteTape(os.Stdout))
	// Output: #
}
