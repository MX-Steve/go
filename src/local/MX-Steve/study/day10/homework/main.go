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
	// Get -> show page ; Post -> handle data
	if r.Method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		err := createUser(username, password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "https://www.baidu.com", 301)
	} else {
		t, err := template.ParseFiles("./register.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		t.Execute(w, nil)
	}
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(500)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		err = queryUser(username, password)
		fmt.Println(err)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		http.Redirect(w, r, "http://www.baidu.com", 301)
	} else {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			w.WriteHeader(500)
		}
		t.Execute(w, nil)
	}
}
func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	err = http.ListenAndServe(":10000", nil)
	if err != nil {
		fmt.Println("start server error, ", err)
	}
}
