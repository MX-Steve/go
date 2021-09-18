package main

import (
	"fmt"
	"net/http"
)

// http server
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1 style='color:green'>hello world!</h1>")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed,err:%v\n", err)
		return
	}
}
