package optimize

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/dim13/golyb"
	"github.com/dim13/golyb/static"
)

func TestOptimize(t *testing.T) {
	testCases := []struct {
		fname  string
		output string
	}{
		{"../testdata/helloworld.b", "Hello World!\n"},
		{"../testdata/202.b", "202"},
		{"../testdata/faraway.b", "#\n"},
	}

	for _, tc := range testCases {
		t.Run(tc.output, func(t *testing.T) {
			p, err := golyb.ParseFile(tc.fname)
			if err != nil {
				t.Fatal(err)
			}
			buf := new(bytes.Buffer)
			All(p).Execute(static.New(nil, buf))
			if buf.String() != tc.output {
				t.Errorf("got %q, want %q", buf.String(), tc.output)
			}
		})
	}
}

func bench(b *testing.B, fname string, optimize bool) {
	p, err := golyb.ParseFile(fname)
	if err != nil {
		b.Fatal(err)
	}
	if optimize {
		p = All(p)
	}
	for i := 0; i < b.N; i++ {
		p.Execute(static.New(nil, ioutil.Discard))
	}
}

func BenchmarkHanoi(b *testing.B) {
	bench(b, "../testdata/hanoi.b", true)
}

func BenchmarkHanoiNoopt(b *testing.B) {
	bench(b, "../testdata/hanoi.b", false)
}

func BenchmarkMandelbrot(b *testing.B) {
	bench(b, "../testdata/mandelbrot.b", true)
}

func BenchmarkMandelbrotNoopt(b *testing.B) {
	bench(b, "../testdata/mandelbrot.b", false)
}

func BenchmarkLong(b *testing.B) {
	bench(b, "../testdata/long.b", true)
}

func BenchmarkLongNoopt(b *testing.B) {
	bench(b, "../testdata/long.b", false)
}
