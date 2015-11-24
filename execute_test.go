package main

import (
	"os"
	"testing"
)

func opt(p Program) Program {
	p = OptContract(p)
	p = OptLoops(p)
	p = OptOffset(p)
	return p
}

func exec(prog string) {
	t := NewStaticTape(os.Stdout)
	p := ParseString(prog)
	Execute(opt(p), t, false)
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

type devNull struct{}

func (devNull) Write(p []byte) (int, error) { return len(p), nil }
func (devNull) Read(p []byte) (int, error)  { return len(p), nil }

func bench(b *testing.B, fname string, optimize bool) {
	p, err := ParseFile(fname)
	if err != nil {
		b.Fatal(err)
	}
	if optimize {
		p = opt(p)
	}
	for i := 0; i < b.N; i++ {
		t := NewStaticTape(devNull{})
		Execute(p, t, false)
	}
}

func BenchmarkHanoi(b *testing.B)           { bench(b, "samples/hanoi.b", true) }
func BenchmarkHanoiNoopt(b *testing.B)      { bench(b, "samples/hanoi.b", false) }
func BenchmarkMandelbrot(b *testing.B)      { bench(b, "samples/mandelbrot.b", true) }
func BenchmarkMandelbrotNoopt(b *testing.B) { bench(b, "samples/mandelbrot.b", false) }
func BenchmarkLong(b *testing.B)            { bench(b, "samples/long.b", true) }
func BenchmarkLongNoopt(b *testing.B)       { bench(b, "samples/long.b", false) }
