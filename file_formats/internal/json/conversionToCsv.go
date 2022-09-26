// https://gosamples.dev/json-to-csv/
package json

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type FruitAndVegetableRank struct {
	Vegetable string `json:"vegetable"`
	Fruit     string `json:"fruit"`
	Rank      int64  `json:"rank"`
}

func ConvertJSONToCsv(src, dest string) error {
	// 1. read source file (json file)
	srcFile, e := os.Open(src)
	if e != nil {
		return e
	}
	defer srcFile.Close()

	var jsonData []FruitAndVegetableRank
	if e := json.NewDecoder(srcFile).Decode(&jsonData); e != nil {
		return e
	}

	// 2. open a file for csv format
	dstFile, e := os.OpenFile(dest, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if e != nil {
		return e
	}
	defer dstFile.Close()

	writer := csv.NewWriter(dstFile)
	defer writer.Flush()

	// 3. write to the buffer headers
	headers := []string{"vegetable", "fruit", "rank"}
	if e := writer.Write(headers); e != nil {
		return e
	}

	// 4. write all jsonData to csv format
	for _, r := range jsonData {
		csvRow := append([]string{}, r.Vegetable, r.Fruit, fmt.Sprint(r.Rank))
		if e := writer.Write(csvRow); e != nil {
			return e
		}
	}
	return nil
}

type ShoppingRecord struct {
	// 1. Create a struct for storing CSV lines and annotate it with JSON struct field tags
	Vegetable string `json:"vegetable"`
	Fruit     string `json:"fruit"`
	Rank      int    `json:"rank"`
}

func convertCSVtoJSON(data [][]string) []ShoppingRecord {
	var list []ShoppingRecord
	for i, line := range data {
		// skip headers
		if i == 0 {
			continue
		}
		var rec ShoppingRecord
		// walk through fields
		for j, field := range line {
			switch j {
			case 0:
				rec.Vegetable = field
			case 1:
				rec.Fruit = field
			case 2:
				var e error
				rec.Rank, e = strconv.Atoi(field)
				if e != nil {
					continue
				}
			}
		}
		list = append(list, rec)
	}
	return list
}

func ConvertCSVtoJSON(src string) (string, error) {
	// 1. open src csv file
	srcFile, e := os.Open(src)
	if e != nil {
		return "", e
	}
	defer srcFile.Close()

	// 2. read records from csv file
	reader := csv.NewReader(srcFile)
	csvRecords, e := reader.ReadAll()
	if e != nil {
		return "", e
	}

	// 3. convert records to specific type
	data := convertCSVtoJSON(csvRecords)

	// 4. marshal
	jsonData, e := json.MarshalIndent(data, "", "  ")
	if e != nil {
		return "", e
	}
	return string(jsonData), nil
}
