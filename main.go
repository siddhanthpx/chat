package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (th *templateHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	th.once.Do(
		func() {
			th.templ = template.Must(template.ParseFiles(filepath.Join("templates", th.filename)))
		},
	)

	th.templ.Execute(rw, nil)
}

func main() {

	http.Handle("/", &templateHandler{filename: "chat.html"})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
