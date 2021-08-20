package main 

import (
	"fmt"
	"strconv"
)

func nPlus(a string) string {
	num, err := strconv.Atoi(string(a[len(a)-1]))
	if err != nil {
		fmt.Println("error")
		return ""
	}
	card := fmt.Sprintf("card%d",(num + 1))
	return card
}

func main(){
	a := "card7"
	b := nPlus(a)
	fmt.Println(b)
}