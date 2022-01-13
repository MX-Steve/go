package main

import "fmt"

func f1() {
	s := "abc"
	for i, n := 0, len(s); i < n; i++ {
		fmt.Println(s[i])
	}
}
func f2() {
	s := "abc"
	n := len(s)
	for n > 0 {
		n--
		fmt.Println(s[n])
	}
}
func f3() {
	s := "abc"
	for {
		fmt.Println(s)
	}
}
func f4() {
	numbers := [6]int{11, 22, 33}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值为 %d\n", i, x)
	}
}
func f5() {
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j < (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d 是素数\n", i)
		}
	}
}
func main() {
	f5()
	// f4()
	// f1()
	// f2()
	// f3()
}
