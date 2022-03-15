package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Human struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Lesson
}
type Lesson struct {
	Lessons []string `json:"lessons"`
}

func main() {
	jsonStr := `{"name":"Lisi","age":18,"gender":"Male","lessons":["Franch","English"]}`
	strR := strings.NewReader(jsonStr)
	h := &Human{}
	err := json.NewDecoder(strR).Decode(h)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(h)
	f, err := os.Create("./t.json")
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(f).Encode(h)
}
