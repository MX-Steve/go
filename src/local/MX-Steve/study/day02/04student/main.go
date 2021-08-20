package main 

import (
	"fmt"
	// "strings"
)

func main(){
	// words := "how are you you are a student"
	// wordSlice := strings.Split(words," ")
	// wordMap := make(map[string]int, len(wordSlice))
	// for _, word := range wordSlice{
	// 	v, ok := wordMap[word]
	// 	if ok {
	// 		wordMap[word] = v + 1
	// 	} else {
	// 		wordMap[word] = 1
	// 	}
	// }
	// for i,v := range wordMap {
	// 	fmt.Println(i,v)
	// }

	studentMap := make(map[string]map[string]int, 10)
	fmt.Println(studentMap)
	fmt.Printf("studentMap 的指针%p\n",studentMap)
	studentMap["zhangsan"] = make(map[string]int,3)
	studentMap["zhangsan"]["id"] = 3
	studentMap["zhangsan"]["score"] = 79
	studentMap["zhangsan"]["age"] = 24
	for i,v := range studentMap{
		fmt.Println("name: ",i)
		for i2,v2 := range v {
			fmt.Println(i2,v2)
		}
	}


}