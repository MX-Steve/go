package main

import (
	"fmt"
	"encoding/json"
)

func main(){
	var jsonBlob = []byte(`[
		{"name": "slice", "order": "mon", "ID": 1},
		{"name": "qual", "order": "tus", "ID": 2}
	]`)
	type Animal struct {
		ID 		int 	`json:"id"`
		Name	string	`json:"name"`
		Order 	string	`json:"order"`
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n",animals)
}