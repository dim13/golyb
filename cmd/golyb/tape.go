package main

import (
	"errors"
	"io"

	"github.com/dim13/golyb"
	"github.com/dim13/golyb/dynamic"
	"github.com/dim13/golyb/sparse"
	"github.com/dim13/golyb/static"
)

type Tape string

func (t Tape) String() string {
	return string(t)
}

func (t *Tape) Set(v string) error {
	switch v {
	case "static", "dynamic", "sparse":
		*t = Tape(v)
	default:
		return errors.New("unknown tape type")
	}
	return nil
}

func (t Tape) New(r io.Reader, w io.Writer) golyb.Tape {
	switch t {
	case "static":
		return static.New(r, w)
	case "dynamic":
		return dynamic.New(r, w)
	case "sparse":
		return sparse.New(r, w)
	}
	return nil
}

func (_ Tape) Usage() string {
	return "Tape type: static, dynamic of sparse"
}
