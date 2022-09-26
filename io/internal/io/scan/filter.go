package scan

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScanFromConsole() {
	var in string
	fmt.Scanln(&in)
	for in != "\\q" {
		fmt.Println("Read: ", in)
		fmt.Scanln(&in)
	}
}

func ConsoleLineFilter() {
	// Wrapping the unbuffered os.Stdin with a buffered scanner gives us a convenient Scan method
	// that advances the scanner to the next token; which is the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "\\q" {
			return
		}
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// Check for errors during Scan. End of file is expected and not reported by Scan as an error.
	if e := scanner.Err(); e != nil {
		fmt.Fprintln(os.Stderr, "error:", e)
		os.Exit(1)
	}
}
