package main

import (
	"net/http"
	"text/template"
)

func tHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("Layout.html", "Index.html")
	t.ExecuteTemplate(w, "layout", "Hello World!")
}

func main() {
	http.HandleFunc("/", tHandler)
	http.ListenAndServe(":8080", nil)
}
