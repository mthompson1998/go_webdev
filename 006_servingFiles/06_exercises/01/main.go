/*
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/" "/dog/" "/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.
*/

package main

import (
	"io"
	"net/http"
)

func main () {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", myself)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is the home page")
}

func bar(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "toby.jpeg")
}

func myself(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "This is Mike Thompson")
}