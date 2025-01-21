package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var templates = make(map[string]*template.Template)

func main() {
	for _, e := range unwrap(os.ReadDir("./templates")) {
		base := strings.TrimSuffix(e.Name(), ".html")
		file := "templates/" + e.Name()
		templates[base] = template.Must(template.ParseFiles(file))
	}

	http.HandleFunc("/", root)
	try(http.ListenAndServe(":6969", nil))
}

func try(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func unwrap[T any](x T, err error) T {
	try(err)
	return x
}

func root(w http.ResponseWriter, r *http.Request) {
	try(templates["index"].Execute(w, nil))
}