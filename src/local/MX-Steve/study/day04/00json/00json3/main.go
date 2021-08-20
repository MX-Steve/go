package main 

import (
	"fmt"
	"encoding/json"
)

func main(){
	// ColorGroup struct
	type ColorGroup struct {
		ID		int			`json:"id"`
		Name	string 		`json:"name"`
		Colors	[]string	`json:"color"`
	}

	group := ColorGroup{
		ID: 1,
		Name: "社长",
		Colors: []string{"red","white","blue"},
	}

	b,err := json.Marshal(group)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("This Marshal result is %+v\n",string(b))
}