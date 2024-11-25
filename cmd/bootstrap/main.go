package main

import (
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	t := template.Must(template.ParseGlob("templates/*.gotmpl"))

	directory := os.Args[1]

	create(directory, "main.go", t.Lookup("main"))
	create(directory, "main_test.go", t.Lookup("main_test"))
}

func create(directory, name string, t *template.Template) {
	path := filepath.Join(directory, name)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := t.Execute(file, nil); err != nil {
		panic(err)
	}
}
