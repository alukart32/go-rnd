package main

import (
	"html/template"
	"log"
	"os"
)

const (
	helloTempPath   = "./assets/hello.gohtml"
	contextTempPath = "./assets/context.gohtml"
)

func main() {
	//helloTemplate()
	contextTemplate()
}

func helloTemplate() {
	t, err := template.ParseFiles(helloTempPath)
	if err != nil {
		log.Fatal(err)
	}

	tempData := &struct {
		Name string
	}{
		Name: "darkness",
	}
	err = t.Execute(os.Stdout, tempData)
	if err != nil {
		log.Fatal(err)
	}
}

type Dog struct {
	Name string
	Age  int
}

type Test struct {
	HTML     string
	SafeHTML template.HTML
	Title    string
	Path     string
	Dog      Dog
	Map      map[string]string
}

func contextTemplate() {
	t, err := template.ParseFiles(contextTempPath)
	if err != nil {
		log.Fatal(err)
	}

	data := &Test{
		HTML:     "<h1>A header!</h1>",
		SafeHTML: template.HTML("<h1>A Safe header</h1>"),
		Title:    "Backslash! An in depth look at the \"\\\" character.",
		Path:     "/dashboard/settings",
		Dog:      Dog{"Fido", 6},
		Map: map[string]string{
			"key":       "value",
			"other_key": "other_value",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
