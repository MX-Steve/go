package main 

import (
	"fmt"
	"strings"
)

func f1(){
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("hello world")
}

var (
	coins = 50
	users = []string{"Matthew", "Sarah","Augustus","Heidi","Emilie","Peter"}
	distribution = make(map[string]int,len(users))
)

func dispatchChoin() int {
	sum := 0
	for _, name := range users{
		if strings.Contains(name,"e") || strings.Contains(name,"E"){
			distribution[name] = distribution[name] + 1
			sum += 1
		}
		if strings.Contains(name,"i") || strings.Contains(name,"I"){
			distribution[name] = distribution[name] + 2
			sum += 2
		}
		if strings.Contains(name,"o") || strings.Contains(name,"O"){
			distribution[name] = distribution[name] + 3
			sum += 3
		}
		if strings.Contains(name,"u") || strings.Contains(name,"U"){
			distribution[name] = distribution[name] + 4
			sum += 4
		}
	}
	return coins - sum
}


func dispatchChoin2() int {
	sum := 0
	for _, name := range users {
		for _,char := range name {
			switch char {
			case 'e', 'E' :
				distribution[name] = distribution[name] + 1
				sum = sum + 1
			case 'i', 'I':
				distribution[name] = distribution[name] + 2
				sum = sum + 2
			case 'o', 'O':
				distribution[name] = distribution[name] + 3
				sum = sum + 3
			case 'u', 'U':
				distribution[name] = distribution[name] + 4
				sum = sum + 4
			}
		}
	}
	return coins - sum
}

func main(){
	// f1()

	// left := dispatchChoin()
	// fmt.Println("剩下: ",left)
	// for k, v := range distribution {
	// 	fmt.Println(k, v)
	// }

	left := dispatchChoin2()
	fmt.Println("剩下： ",left)
	for k,v := range distribution {
		fmt.Println(k, v)
	}
}