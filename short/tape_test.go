package short

import (
	"bytes"
	"testing"

	"github.com/dim13/golyb"
)

func TestStatic(t *testing.T) {
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
			p.Execute(NewTape(buf))
			if buf.String() != tc.output {
				t.Errorf("got %q, want %q", buf.String(), tc.output)
			}
		})
	}
}
