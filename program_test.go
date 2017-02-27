package golyb

import (
	"os"
	"testing"
)

func exec(prog string) {
	ParseString(prog).Optimize().Execute(NewStaticTape(os.Stdout))
}

// Hello World!
const HelloWorld = `++++++++++[>+++++++>++++++++++>+++>+<<<<-]
>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.`

func ExampleHelloWorld() {
	exec(HelloWorld)
	// Output: Hello World!
}

// Prints 202
const Numeric = `>+>+>+>+>++<[>[<+++>-]<<]>.`

// Numeric output
func ExampleNumeric() {
	exec(Numeric)
	// Output: 202
}

// Goes to cell 30000 and reports from there with a #. (Verifies that the
// array is big enough.)
const FarAway = `++++[>++++++<-]>[>+++++>+++++++<<-]>>++++<
[[>[[>>+<<-]<]>>>-]>-[>+>+<<-]>]+++++[>+++++++<<++>-]>.<<.`

func ExampleFarAway() {
	exec(FarAway)
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

func BenchmarkHanoi(b *testing.B)         { bench(b, "testdata/hanoi.b", true) }
func BenchmarkHanoiRaw(b *testing.B)      { bench(b, "testdata/hanoi.b", false) }
func BenchmarkMandelbrot(b *testing.B)    { bench(b, "testdata/mandelbrot.b", true) }
func BenchmarkMandelbrotRaw(b *testing.B) { bench(b, "testdata/mandelbrot.b", false) }
func BenchmarkLong(b *testing.B)          { bench(b, "testdata/long.b", true) }
func BenchmarkLongRaw(b *testing.B)       { bench(b, "testdata/long.b", false) }
