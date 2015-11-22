// Go f*ck your self -- a BrainF*ck interpreter
package main

const helloWorld = `++++++++++        initializes cell zero to 10
	[
	   >+++++++>++++++++++>+++>+<<<<-
	]                 this loop sets the next four cells to 70/100/30/10 
	>++.              print   'H'
	>+.               print   'e'
	+++++++.                  'l'
	.                         'l'
	+++.                      'o'
	>++.                      space
	<<+++++++++++++++.        'W'
	>.                        'o'
	+++.                      'r'
	------.                   'l'
	--------.                 'd'
	>+.                       '!'
	>.                        newline`

func ExampleHelloWorld() {
	Execute(ParseString(helloWorld), NewTape())
	// Output: Hello World!
}

const numeric = `>+>+>+>+>++<[>[<+++>-]<<]>.`

// Numeric output
func ExampleNumeric() {
	Execute(ParseString(numeric), NewTape())
	// Output: 202
}

const faraway = `++++[>++++++<-]>[>+++++>+++++++<<-]>>++++<
	[[>[[>>+<<-]<]>>>-]>-[>+>+<<-]>]
	+++++[>+++++++<<++>-]>.<<.`

// Goes to cell 30000 and reports from there with a #. (Verifies that the
// array is big enough.)
func ExampleFarAway() {
	Execute(ParseString(faraway), NewTape())
	// Output: #
}
