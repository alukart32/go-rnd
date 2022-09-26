package files

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"alukart32.com/tools"
)

const BUFFER_SIZE = 512

func isEOF(e error) bool {
	if e == nil {
		return false
	}

	if errors.Is(e, io.EOF) {
		return true
	} else {
		panic(e)
	}
}

func ReadFile(path []byte) {
	buf := make([]byte, BUFFER_SIZE)
	f, e := os.Open(string(path)) // return *os.File and error
	tools.CheckErr(e)
	defer f.Close() // we defer until the function return
	for {
		n, e := f.Read(buf)
		isEOF(e)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func ReadBufferFile(path string) {
	buf := make([]byte, BUFFER_SIZE)
	f, e := os.Open(path) // return *os.File and error
	tools.CheckErr(e)
	defer f.Close() // we defer until the function return
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush() // to flush all output
	for {
		n, e := r.Read(buf)
		isEOF(e)
		if n == 0 {
			break
		}
		w.Write(buf[:n])
	}
	w.Write([]byte{'\n'}) // add \n for console output
}

func ReadSeekFile(path string) {
	f, e := os.Open(string(path))
	tools.CheckErr(e)
	defer f.Close()
	seek, e := f.Seek(4, 0)
	tools.CheckErr(e)
	buf2 := make([]byte, 3)
	readBytes, e := f.Read(buf2)
	tools.CheckErr(e)
	fmt.Printf("%d bytes @ position %d\n", readBytes, seek)
	fmt.Printf("%v\n", string(buf2[:readBytes]))
}

func ReadAtLeast(path string) {
	f, e := os.Open(string(path))
	tools.CheckErr(e)
	defer f.Close()
	seek, e := f.Seek(4, 0)
	tools.CheckErr(e)
	buf := make([]byte, 3)
	n, e := io.ReadAtLeast(f, buf, 2)
	tools.CheckErr(e)
	fmt.Printf("%d bytes @ position %d\n", n, seek)
	fmt.Printf("%v\n", string(buf[:n]))
}

func BufPeek(path string) {
	f, e := os.Open(path)
	tools.CheckErr(e)
	defer f.Close()
	r := bufio.NewReader(f)
	buf, e := r.Peek(5)
	tools.CheckErr(e)
	fmt.Printf("5 bytes: %s\n", string(buf))
}

func MkdirIfNotExists(path string) {
	if _, e := os.Open(path); errors.Is(e, os.ErrNotExist) {
		os.Mkdir(path, 0755)
		fmt.Printf("The directory %s was created\n", path)
	} else {
		fmt.Printf("The directory %s already exists!\n", path)
	}
}

// Read a single line from standard input
// https://gosamples.dev/read-user-input/

// Use the fmt.Scanln() function if you want to read each word of a line into a different variable,
// and there is a certain number of words to read, no less, no more.
func ReadSingleLineFmtScanln() {
	fmt.Println("input text:")
	var w1, w2, w3 string
	n, err := fmt.Scanln(&w1, &w2, &w3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("number of items read: %d\n", n)
	fmt.Printf("read line: %s %s %s-\n", w1, w2, w3)
}

// Use the bufio.Reader if you want to read a full line of text together with the newline character.
func ReadSingleLineBufioReader() {
	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read line: %s-\n", line)
}

// Use the bufio.Scanner to read a full line of text in a convenient way without the newline character.
func ReadSingleLineBufioScanner() {
	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read line: %s-\n", scanner.Text())
}

// Read multiple lines from console

// With the fmt.Scan() you can read multiple lines of text only if the line consists of a single word.
func ReadMultStringLinesFmtScan() {
	fmt.Println("input text:")
	var w1, w2, w3 string
	n, err := fmt.Scan(&w1, &w2, &w3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("number of items read: %d\n", n)
	fmt.Printf("read text: %s %s %s-\n", w1, w2, w3)
}

// Use the bufio.Reader if you want to read multiple lines together with the newline character at the end of each line.
//
// input text:
// ab cd ef gh
// hj kl mn op
//
// output:
// ab cd ef gh
//
// hj kl mn op
func ReadMultStringLinesBufioReader() {
	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)

	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
	}

	fmt.Println("output:")
	for _, l := range lines {
		fmt.Println(l)
	}
}

// The most recommended and universal method of reading multiple lines is to use the bufio.Scanner,
// which allows you to get a list of input lines without the newline character at the end of each line.
//
// input text:
// ab cd ef gh
// hj kl mn op
//
// output:
// ab cd ef gh
// hj kl mn op
func ReadMultStringLinesBufioScanner() {
	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()

		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}

	fmt.Println("output:")
	for _, l := range lines {
		fmt.Println(l)
	}
}

// Read a single character from terminal
func ReadSingleRuneScanf() {
	fmt.Println("input text:")
	var char rune
	_, err := fmt.Scanf("%c", &char)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read character: %c-\n", char)
}

// ReadRune() method, which returns the rune, the size of the rune, and the reading error
func ReadSingleRuneBufioReader() {
	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read character: %c-\n", char)
}

// Read formatted user input
//
// input text:
// Anna is born in Germany
// number of items read: 2
// Anna Germany
func ReadFormatedUserInput() {
	fmt.Println("input text:")
	var name string
	var country string
	n, err := fmt.Scanf("%s is born in %s", &name, &country)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("number of items read: %d\n", n)
	fmt.Println(name, country)
}

// Read numbers from user input
func ReadNumersFmtScanf() {
	fmt.Println("input number:")
	var number int64
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read number: %d\n", number)
}

func ReadNumersFmtScan() {
	fmt.Println("input number:")
	var number int64
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read number: %d\n", number)
}
