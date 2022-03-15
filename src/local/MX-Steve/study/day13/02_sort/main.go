package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type personSlice []person

func (p personSlice) Len() int { return len(p) }

func (p personSlice) Less(i, j int) bool { return p[i].Age < p[j].Age }

func (p personSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	a := personSlice{
		{Name: "AAA", Age: 84},
		{Name: "BBB", Age: 14},
		{Name: "CCC", Age: 34},
		{Name: "DDD", Age: 24},
		{Name: "EEE", Age: 94},
		{Name: "FFF", Age: 64},
	}
	sort.Sort(a)
	fmt.Println("Sort: ", a)
	sort.Stable(a)
	fmt.Println("Stable:", a)
}
