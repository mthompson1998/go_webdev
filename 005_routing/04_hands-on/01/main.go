package main

import (
	"io"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo foo")
}

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog dog")
}

func m(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Mr. Mike")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)
	http.ListenAndServe(":8080", nil)
}