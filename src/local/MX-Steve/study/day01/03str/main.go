package main

import (
	"fmt"
)

func changeString() {
	// 强制类型转换
	s5 := "big"
	byterArray := []byte(s5)
	fmt.Println(byterArray)
	byterArray[0] = 'B'
	s5 = string(byterArray)
	fmt.Println(s5)
}
func strDesc() {
	s6 := "hello"
	// for _, item := range s6 {
	// 	fmt.Println(item)
	// }
	// lastStr := ""
	// for i := len(s6) - 1; i >= 0; i-- {
	// 	lastStr = lastStr + string(s6[i])
	// }
	// fmt.Println(lastStr)
	byteArr := []byte(s6)
	// lastStr := ""
	// for i := len(byteArr) - 1; i >= 0; i-- {
	// 	lastStr = lastStr + string(byteArr[i])
	// }
	// fmt.Println(lastStr)
	for i := 0; i < len(byteArr)/2; i++ {
		byteArr[i], byteArr[len(byteArr)-1-i] = byteArr[len(byteArr)-1-i], byteArr[i]
	}
	fmt.Println(string(byteArr))
}

func main() {
	// s := "hello中国"
	// fmt.Println("abcd")
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// 	fmt.Printf("%c\n", s[i])
	// }
	// for i, item := range s {
	// 	fmt.Println(i, string(item))
	// }
	// changeString()
	strDesc()
}
