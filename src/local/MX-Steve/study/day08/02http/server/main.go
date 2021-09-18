package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http server
func sayHello(w http.ResponseWriter, t *http.Request) {
	// w.Write([]byte("hello world"))
	data, err := ioutil.ReadFile("./hello.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(data)
}

// http server
func index(w http.ResponseWriter, t *http.Request) {
	// w.Write([]byte("hello world"))
	// fmt.Println(t.Method)
	t.ParseForm()
	// fmt.Printf("%#v\n", t.Form)
	uname := t.Form.Get("username")
	pwd := t.Form.Get("password")
	fmt.Println(uname)
	fmt.Println(pwd)

	w.Write([]byte("index22"))
}
func main() {
	http.HandleFunc("/demo", sayHello) // register url
	http.HandleFunc("/form", index)    // register url
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		panic(err)
	}
}
