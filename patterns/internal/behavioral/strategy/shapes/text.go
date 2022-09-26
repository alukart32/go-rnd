package shapes

import (
	"bytes"
	"io"

	"alukart32.com/usage/patterns/internal/behavioral/strategy"
)

type TextSquare struct {
	strategy.PrintOutput
}

func (t *TextSquare) Print() error {
	r := bytes.NewReader([]byte("square\n"))
	io.Copy(t.Writer, r)
	return nil
}
