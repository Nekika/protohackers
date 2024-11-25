package main

import (
	"embed"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*
var f embed.FS

func main() {
	directory := os.Args[1]

	path := filepath.Join(directory, "main.go")

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		panic(err)
	}

	t, err := template.ParseFS(f, "templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := t.ExecuteTemplate(file, "main", nil); err != nil {
		panic(err)
	}
}
