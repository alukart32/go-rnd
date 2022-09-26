package strategy

import (
	"io"
)

type PrintStrategy interface {
	Print() error
	SetLog(io.Writer)    // to add a logger strategy to our types
	SetWriter(io.Writer) // to set the io.Writer strategy
}

type PrintOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (d *PrintOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}

func (d *PrintOutput) SetWriter(w io.Writer) {
	d.Writer = w
}
