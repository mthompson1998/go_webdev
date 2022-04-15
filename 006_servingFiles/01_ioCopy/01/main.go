/*
ioCopy serves images and content to your server?
*/

package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpeg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<img src=toby.jpeg>
	`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("toby.jpeg")
	if err != nil {
		http.Error(w, "file not found", 404)
	}
	defer f.Close()

	io.Copy(w, f)
}