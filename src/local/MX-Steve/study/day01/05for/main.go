package main

import (
	"fmt"
)

func swi() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("不在合理范围")
	}
}

func swi2() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("odd")
	case 2, 4, 6, 8:
		fmt.Println("even")
	default:
		fmt.Println(n)
	}
}
func swi3() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

func gototest() {
	flag := false
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 2 {
				flag = true
				break
			}
			fmt.Printf("%d--%d\n", i, j)
		}
		if flag {
			break
		}
	}
	fmt.Println("two laywers")

}

func ninenine() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dX%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func main() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// swi()
	// swi2()
	// swi3()
	// gototest()
	ninenine()
}
