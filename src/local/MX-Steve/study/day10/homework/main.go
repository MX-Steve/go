package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
1. login
2. register
*/
func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	// Get -> show page ; Post -> handle data
	t, err := template.ParseFiles("./register.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
func main() {
	http.HandleFunc("/register", registerHandler)
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		fmt.Println("start server error, ", err)
	}
}
