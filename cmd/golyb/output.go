package main

import (
	"io"
	"os"
)

func output(out, in string) (io.ReadWriter, error) {
	var err error
	var r io.Reader
	var w io.Writer

	if out != "" {
		w, err = os.Create(out)
		if err != nil {
			return nil, err
		}
	} else {
		w = os.Stdout
	}

	if in != "" {
		r, err = os.Open(in)
		if err != nil {
			return nil, err
		}
	} else {
		r = os.Stdin
	}
	return struct {
		io.Reader
		io.Writer
	}{r, w}, nil
}
