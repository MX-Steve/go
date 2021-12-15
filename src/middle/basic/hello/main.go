package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func ar() {
	a := [3]int{1, 2}
	fmt.Println(a)
	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	c := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}
	fmt.Println(c)
	d := [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(d)
}

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}
func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func sumArr(a [10]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func test_sumArr() {
	rand.Seed(1)
	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	fmt.Printf("sum=%d\n", sum)
}

func main() {
	test_sumArr()
	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [5]int{1, 2, 3, 4, 5}
	printArr(&arr2)
	fmt.Println(arr2)
	a1 := [2]int{1, 2}
	test(a1)
	fmt.Printf("a1: %p\n", &a1)
	fmt.Println(len(a1), cap(a1))
	ar()
	fmt.Println("Hello world!")
	b := strings.Contains("hello", "he")
	fmt.Println(b)
	str := "hello, 中国"
	for _, r := range str {
		fmt.Printf("%c", r)
	}
	s2 := []rune(str)
	s2[8] = '华'
	fmt.Println(string(s2))
}
