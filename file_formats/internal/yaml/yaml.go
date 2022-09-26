package yaml

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
a: string_value_a
b:
  c: 2
  d: [3, 4]
`

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func BasicYaml() error {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		return err
	}
	fmt.Printf("--- t:\n%+v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		return err
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), m)
	if err != nil {
		return err
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		return err
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))

	return nil
}

type StructA struct {
	A string `yaml:"a" json:"a"`
}

type StructB struct {
	// Embedded structs are not treated as embedded in YAML by default. To do that,
	// add the ",inline" annotation below
	StructA `yaml:",inline"`
	B       string `yaml:"b"`
}

var data2 = `
a: string from field "a"
b: string from field "b"
`

func EmbededStruct() {
	data := StructB{}

	e := yaml.Unmarshal([]byte(data2), &data)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("--- data:\n%+v\n\n", data)

	res, e := yaml.Marshal(&data)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("--- data dump:\n%s\n\n", string(res))
}
