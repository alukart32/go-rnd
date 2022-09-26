// https://pkg.go.dev/encoding/csv
package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	"alukart32.com/tools"
)

func ReadCustomCSVFile(path string) {
	f, e := os.Open(path)
	tools.CheckErr(e)
	defer f.Close()

	// format
	// header_1, header_2, header_3
	// # comment
	//   a ;  b  ;  c
	r := csv.NewReader(f)
	r.Comma = ';'
	r.Comment = '#'

	records, e := r.ReadAll()
	tools.CheckErr(e)

	for i := range records {
		fmt.Println(records[i])
	}
}

func WriteCustomCSVFile(path string) {
	f, e := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	tools.CheckErr(e)
	defer f.Close()

	records := [][]string{
		{"header_1", "header_2", "header_3"},
		{"a", "b", "c"},
		{"12", " r45t", "ooow"},
	}

	w := csv.NewWriter(f)
	w.UseCRLF = true
	w.Comma = ';'
	w.WriteAll(records)

	tools.CheckErr(w.Error())
}
