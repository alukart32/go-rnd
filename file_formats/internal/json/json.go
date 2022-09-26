package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func MarshalBasic() {
	type Object struct {
		ID     int
		Name   string
		Colors []string
	}
	group := &Object{
		ID:     1,
		Name:   "object",
		Colors: []string{"red", "purple"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(b)
	fmt.Println()
}

func MarshalIndent() {
	type color struct {
		Name    string `json:"name"`
		HexCode string `json:"hex_code"`
	}
	data := make(map[string]interface{})
	data = map[string]interface{}{
		"best_fruit":      "apple",                                 // string value
		"best_vegetables": []string{"potato", "carrot", "cabbage"}, // array value
		"best_websites": map[int]interface{}{ // map value
			1: "gosamples.dev", // integer key, string value
		},
		"best_color": color{
			Name:    "red",
			HexCode: "#FF0000",
		}, // struct value
	}

	jsonBytes, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))
}

// json.MarshalIndent generates JSON encoding of the value with indentation.
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func PrettyPrintByMarshalingValue() {
	fruit := struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}{
		Name:  "Strawberry",
		Color: "red",
	}
	res, err := PrettyStruct(fruit)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

// generates JSON encoding of the value with indentation.
func PrettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "	 ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

func PrettyPrintByEncodingValue() {
	fruit := struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}{
		Name:  "Strawberry",
		Color: "red",
	}
	var buf bytes.Buffer
	err := PrettyEncode(fruit, &buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

// don't change original JSON data
func PrettyString(data string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(data), "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func PrettyPrintByIndent() {
	fruitJSON := `{"name": "Strawberry", "color": "red"}`
	res, err := PrettyString(fruitJSON)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
