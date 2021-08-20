package main
 
import (
	"fmt"
)

func main(){
	// var a int
	// a = 100
	// fmt.Println(a)
	// b := &a // 取变量 a 的内存地址
	// fmt.Println(b)
	// fmt.Printf("%T\n",b)
	// fmt.Println(*b) // * 取地址对应的值
	// fmt.Println(b == &a)

	// 指针应用
	a := [3]int{1,2,3}
	modifyArray(a) // 在函数中复制了一个新的数组，不会影响外层数组
	fmt.Println(a)
	modifyArray2(&a) // 在函数中复制了一个新的数组，不会影响外层数组
	fmt.Println(a)
}

// 接收一个数组的参数
func modifyArray2(a1 *[3]int){
	// 法 1
	// (*a1)[0]=100
	// 法 2
	a1[0]=100
}

// 内部修改，对外部不影响
func modifyArray(a1 [3]int){
	a1[0] = 100
}