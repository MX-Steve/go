package main 

import (
	"fmt"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println("PATH: ", r.URL.Path)
	fmt.Println("SCHEME: ", r.URL.Scheme)
	fmt.Println("METHOD: ", r.Method)
	fmt.Println()
	fmt.Fprintf(w, "<h1>Index Page</h1>")
	fmt.Println(w)
}

func main(){
	fmt.Println("Start")
	http.HandleFunc("/", HandleIndex)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
// http 测试 
// ]# go run main.go
// ]# curl http://192.168.226.130:8000/a
// ]# go run main.go 
// PATH: /a
// SCHEME:
// METHOD: GET
// &{0xc000074000 0xc000076000 {} 0x4cf640 true false false false 0 {0 0} 0xc000062100 {0xc00007c000 map[] false false} map[] false 19 -1 200 false false [] 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0] 0xc00002a0e0 0}
