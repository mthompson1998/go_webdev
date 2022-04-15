package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/tm/", thompson)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bar bar")
}

func thompson(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	err = tpl.ExecuteTemplate(w, "something.gohtml", "Thompson")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

