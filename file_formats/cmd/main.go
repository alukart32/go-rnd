package main

import (
	"fmt"
	"log"

	"alukart32.com/usage/formats/internal/json"
	"alukart32.com/usage/formats/internal/yaml"
)

func main() {
	fmt.Printf("JSON\n\nBasicMarshal::\n\t")
	json.MarshalBasic()

	fmt.Printf("\nMarshalIndent::\n")
	json.MarshalIndent()

	fmt.Printf("\nPretty print by marshaling value::\n")
	json.PrettyPrintByMarshalingValue()

	fmt.Printf("\nPretty print by encoding value::\n")
	json.PrettyPrintByEncodingValue()

	fmt.Printf("Pretty print by indent value::\n")
	json.PrettyPrintByIndent()

	fmt.Printf("Convert JSON to CSV\n\n")
	if err := json.ConvertJSONToCsv("../assets/data.json", "../assets/data.csv"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Convert CSV to JSON::\n")
	if str, err := json.ConvertCSVtoJSON("../assets/data.csv"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(str)
	}

	fmt.Printf("\nYAML\n\nBasicYaml::\n")
	err := yaml.BasicYaml()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Embeded struct::\n")
	yaml.EmbededStruct()
}
