package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html", "./ul.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}
func main() {
	http.HandleFunc("/index", index)
	panic(http.ListenAndServe("0.0.0.0:8081", nil))
}
