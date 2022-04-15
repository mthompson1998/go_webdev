package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "cat cat cat")
}


func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)

	// Use htt.HandleFunc so that you dont need to create types (hotdog, hotcat)

	http.ListenAndServe(":8080", nil)
}