package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"alukart32.com/tools"
)

const (
	SPACE       = " "
	COMMA_DELIM = ","
	BRACKETS    = "[]"
)

func FlushToFile(s string) {
	// WriteFile writes data to the named file, creating it if necessary.
	tools.CheckErr(os.WriteFile("../assets/wr0.txt", []byte(s), 0755))
}

func CreateAndFlushToFile(path string, data interface{}) {
	f, e := os.Create(path) // better this os.OpenFile("../assets/wr1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	tools.CheckErr(e)
	defer f.Close()

	f.Write(data.([]byte))
	f.Sync()
}

func WriteString(data []int) {
	f, e := os.OpenFile("../assets/wr2.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	tools.CheckErr(e)
	defer f.Close()

	// not effectively
	// [x y z] -> [x,y,z] -> x,y,z
	str := strings.Replace(fmt.Sprint(data), SPACE, COMMA_DELIM, -1)
	_, e = f.WriteString(str[1 : len(str)-1])
	tools.CheckErr(e)
}

func BufWrite(data string) {
	f, e := os.OpenFile("../assets/wr3.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	tools.CheckErr(e)
	defer f.Close()

	b := bufio.NewWriter(f)
	defer b.Flush()
	_, e = b.WriteString(data)
	tools.CheckErr(e)
}
