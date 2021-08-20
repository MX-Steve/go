package main

import (
	"fmt"
	"strings"
)

func main() {
	// // map
	// var m1 map[string]int
	// m1 = make(map[string]int)
	// m1["name"] = 20
	// fmt.Println(m1)
	// // get k-v from map
	// v, ok := m1["name"]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("not exists")
	// }
	// s := "how do you do"
	// getCount(s)
	// student()
	left := countCoin()
	fmt.Println(left)
	fmt.Println(distribution)
}

var (
	coins        = 50
	users        = []string{"zes", "lis", "ww", "zl"}
	distribution = make(map[string]int, len(users))
)

// dispatch coin
func countCoin() int {
	sum := 0
	for _, name := range users {
		if strings.Contains(name, "e") || strings.Contains(name, "E") {
			distribution[name] += 1
			sum++
		} else if strings.Contains(name, "i") || strings.Contains(name, "I") {
			distribution[name] += 2
			sum += 2
		}
	}
	return coins - sum
}

func student() {
	studentMap := make(map[string]map[string]int, 10)
	studentMap["lisi"] = make(map[string]int, 3)
	studentMap["lisi"]["id"] = 1
	studentMap["lisi"]["age"] = 19
	studentMap["lisi"]["score"] = 90
	studentMap["zs"] = make(map[string]int, 3)
	studentMap["zs"]["id"] = 2
	studentMap["zs"]["age"] = 18
	studentMap["zs"]["score"] = 95

	for k, v := range studentMap {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println(k2, v2)
		}
	}
}

func getCount(str string) {
	wordSlice := strings.Split(str, " ")
	wordMap := make(map[string]int, len(wordSlice))
	for _, v := range wordSlice {
		_, ok := wordMap[v]
		if ok {
			wordMap[v] += 1
		} else {
			wordMap[v] = 1
		}
	}
	for k, v := range wordMap {
		fmt.Println(k, v)
	}
}
