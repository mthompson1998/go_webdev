package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mthompson1998/go_webdev/015_gomongo/001/models"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Index</title>
	</head>
	<body>
		<a href="/user/19823748/">GO TO: http://localhost:8080/user/19823748</a>
	</body>
	</html>`

		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name: "James Bond",
		Gender: "male",
		Age:	32,
		Id:		p.ByName("id"),
	}
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", uj)
}