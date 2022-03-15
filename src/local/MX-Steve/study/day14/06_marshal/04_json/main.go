package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name   json.RawMessage `json:"name"`
	Gender string          `json:"gender"`
	Age    int             `json:"age"`
	Lesson
}

type Lesson struct {
	Lessons []string `json:"lessons"`
}

func main() {
	jsonStr := `{"name":"Lisi","age":19,"gender":"Male","lessons":["English","History"]}`
	var hu Human
	if err := json.Unmarshal([]byte(jsonStr), &hu); err == nil {
		fmt.Printf("\n 结构体Human \n")
		fmt.Printf("%+v \n", hu)
		fmt.Println(hu.Lessons)
	}
	var UName string
	if err := json.Unmarshal(hu.Name, &UName); err == nil {
		fmt.Printf("\n Human.Name: %s \n", UName)
	}
}
