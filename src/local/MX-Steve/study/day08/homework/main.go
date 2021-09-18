package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./login.html")
		if err != nil {
			fmt.Println(err)
		}
		w.Write(data)
	} else if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if username == "lisi" && password == "123456" {
			http.Redirect(w, r, "/index", http.StatusFound)
		} else {
			w.Write([]byte("login err"))
		}
	}
}

type bu struct {
	Id    int
	Name  string
	Price float32
}

func index(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./index.html")
	if err != nil {
		fmt.Println(err)
	}
	t, err := template.New("index").Parse(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	bus := map[int]bu{
		1: {10, "洗面奶", 23.45},
		2: {11, "洗发水", 24.45},
		3: {12, "牙膏", 2.45},
		4: {13, "牙刷", 3.45},
	}
	profile := map[string]string{
		"name":  "里斯",
		"hobby": "basketball",
	}
	s := map[string]interface{}{
		"bus":     bus,
		"profile": profile,
	}
	t.Execute(w, s)
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/index", index)
	panic(http.ListenAndServe("0.0.0.0:8081", nil))
}
