package main

import (
	"fmt"
	"io"
	"os"

	"alukart32.com/go/iocat/iocat"
)

// iocat filepath
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "no filepath")
		os.Exit(1)
	}

	for _, f := range os.Args[1:] {
		err := cat(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "can't copy %v: %v", f, err)
		}
	}
}

func cat(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("can't open file")
	}
	defer f.Close()

	wc := iocat.NewWriter(os.Stdout)
	if _, err := io.Copy(wc, f); err != nil {
		return err
	}

	return wc.Close()
}
