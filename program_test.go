package main

import (
	"os"
	"testing"
)

func exec(prog string) {
	ParseString(prog).Optimize().Execute(NewStaticTape(os.Stdout))
}

// Hello World!
const helloWorld = `++++++++++[>+++++++>++++++++++>+++>+<<<<-]
>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.`

func ExamplehelloWorld() {
	exec(helloWorld)
	// Output: Hello World!
}

// Prints 202
const numeric = `>+>+>+>+>++<[>[<+++>-]<<]>.`

// Numeric output
func Examplenumeric() {
	exec(numeric)
	// Output: 202
}

// Goes to cell 30000 and reports from there with a #. (Verifies that the
// array is big enough.)
const faraway = `++++[>++++++<-]>[>+++++>+++++++<<-]>>++++<
[[>[[>>+<<-]<]>>>-]>-[>+>+<<-]>]+++++[>+++++++<<++>-]>.<<.`

func Examplefaraway() {
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
		p = p.Optimize()
	}
	for i := 0; i < b.N; i++ {
		p.Execute(NewStaticTape(devNull{}))
	}
}

func BenchmarkHanoi(b *testing.B)         { bench(b, "samples/hanoi.b", true) }
func BenchmarkHanoiRaw(b *testing.B)      { bench(b, "samples/hanoi.b", false) }
func BenchmarkMandelbrot(b *testing.B)    { bench(b, "samples/mandelbrot.b", true) }
func BenchmarkMandelbrotRaw(b *testing.B) { bench(b, "samples/mandelbrot.b", false) }
func BenchmarkLong(b *testing.B)          { bench(b, "samples/long.b", true) }
func BenchmarkLongRaw(b *testing.B)       { bench(b, "samples/long.b", false) }
