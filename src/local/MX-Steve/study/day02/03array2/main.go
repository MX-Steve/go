package main

import (
	"fmt"
)

func main(){
	// var a [3]int 
	// a = [3]int{1,2,3}
	// var b [3][2]int 
	// b = [3][2]int{
	// 	[2]int{1,2},
	// 	[2]int{3,4},
	// }
	// fmt.Println(a)
	// fmt.Println(b)

	// var c = [3][2]int{
	// 	{1,2},
	// 	{3,4},
	// 	{5,6},
	// }
	// fmt.Println(c)

	// var d = [...][2]int{
	// 	{1,2},
	// 	{3,4},
	// }
	// fmt.Println(d)
	// fmt.Println(d[1][1])
	// for i:=0;i<len(d);i++{
	// 	for j:=0;j<len(d[i]);j++{
	// 		fmt.Println(d[i][j])
	// 	}
	// }
	// for _,v1 := range d{
	// 	fmt.Println(v1)
	// 	for _,v2 := range v1 {
	// 		fmt.Println(v2)
	// 	}
	// }

	// a :=[2]int{1,2}
	// b := a 
	// b[0]=100
	// fmt.Println(a)
	// fmt.Println(b)

	c :=[3][2]int{
		{1,2},
		{3,4},
		{5,6},
	}
	d := c
	d[0][0]=100
	fmt.Println(c)
	fmt.Println(d)
}