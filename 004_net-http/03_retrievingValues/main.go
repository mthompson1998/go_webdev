/*
Step 1: Handler interface
		type Handler interface {
			ServeHTTP(ResponseWriter, *Request)
		}
Step 2: Listen and Serve
		L&S takes a handler
Step 3: Retrieving Values (*http.Request)
*/

package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}