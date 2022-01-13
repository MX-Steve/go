package main

import "fmt"

func s1() {
	var grade string = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B":
		fmt.Println("良好")
	case grade == "C":
		fmt.Println("及格")
	case grade == "D":
		fmt.Println("加油")
	}
}
func s2() {
	var x interface{}
	switch x.(type) {
	case int:
		fmt.Println("int 类型")
	case float64:
		fmt.Println("float64类型")
	case nil:
		fmt.Println("nil 类型")
	default:
		fmt.Println("未知类型")
	}
}
func s3() {
	var n = 0
	switch {
	case n > 0 && n <= 10:
		fmt.Println("0<n<=10")
	case n > 10:
		fmt.Println("n>10")
	default:
		fmt.Println("000")
	}
}
func main() {
	s3()
}
