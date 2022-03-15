package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 5, 4, -1, 11, 9, -14}
	sort.Ints(a)
	fmt.Println(a)
	ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad"}
	sort.Strings(ss)
	fmt.Println(ss)
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	fmt.Printf("After reverse: %v\n", ss)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Printf("After reverse: %v\n", a)
	b := []int{11, 2, 33, 23, -1, 24}
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	fmt.Println(b)
	sort.Ints(b)
	fmt.Println(b)

}
