package iocat

import (
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

type writer struct {
	pw   *io.PipeWriter
	done chan struct{}
}

func NewWriter(w io.Writer) io.WriteCloser {
	pr, pw := io.Pipe()

	wc := writer{pw, make(chan struct{})}
	go func() {
		defer close(wc.done)
		pr.CloseWithError(copy(w, pr))
	}()
	return &wc
}

func (w *writer) Write(data []byte) (int, error) {
	return w.pw.Write(data)
}

func (w *writer) Close() error {
	if err := w.pw.Close(); err != nil {
		return err
	}
	<-w.done
	return nil
}

// copy copies the given image reader and encodes it as an
// iTerm2 image into the writer.
func copy(w io.Writer, r io.Reader) error {
	header := strings.NewReader("\033]1337;File=inline=1:")
	footer := strings.NewReader("\a\n")

	pr, pw := io.Pipe()
	// encode body
	go func() {
		defer pw.Close()

		wc := base64.NewEncoder(base64.StdEncoding, pw)
		_, err := io.Copy(wc, r)
		if err != nil {
			pw.CloseWithError(fmt.Errorf("could not encode image: %v", err))
			return
		}

		if err = wc.Close(); err != nil {
			pw.CloseWithError(fmt.Errorf("could not close base64 encoder"))
		}
	}()

	_, err := io.Copy(w, io.MultiReader(header, pr, footer))
	return err
}
