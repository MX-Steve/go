package main

import "fmt"

// 如果切片的容量小于 1024 个元素， 于是扩容的时候就翻倍增加容量。
// 如果元素个数超过 1024 个元素，增长因子就变成 1.25， 即每次增加原来的四分之一

func f1() {
	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(data)
	fmt.Println(s)
}
func f2() {
	s1 := []int{1, 2, 3, 4, 8: 100}
	fmt.Println(s1, len(s1), cap(s1))
	s2 := make([]int, 6, 8)
	fmt.Println(s2, len(s2), cap(s2))
	s3 := make([]int, 6)
	fmt.Println(s3, len(s3), cap(s3))
}
func f3() {
	d := [5]struct {
		x int
	}{}
	s := d[:]
	d[1].x = 10
	s[2].x = 20
	fmt.Println(d)
	fmt.Printf("%p, %p\n", &d, &d[0])
}
func f4() {
	var a = []int{1, 2, 3}
	var b = []int{4, 5, 6}
	c := append(a, b...)
	fmt.Printf("c: %v\n", c)
	d := append(c, 7)
	fmt.Println(d)
	e := append(d, 8, 9, 10)
	fmt.Println(e)
}
func f6() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("array data: %v\n", data)
	s1 := data[8:]
	s2 := data[:5]
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied s1: %v\n", s1)
	fmt.Printf("copied s2: %v\n", s2)
	fmt.Printf("last data: %v\n", data)
}
func f5() {
	f6()
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1: %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2: %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied s1: %v\n", s1)
	fmt.Printf("copied s2: %v\n", s2)
	s3 := []int{1, 2, 3}
	s3 = append(s3, s2...)
	fmt.Printf("append s3: %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last s3: %v\n", s3)

}
func f7() {
	data := [...]int{11, 22, 33, 44, 55, 66}
	s := data[:]
	for index, value := range s {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}
}
func f8() {
	array := []int{1, 2, 3, 5}
	slice := make([]int, 6)
	n := copy(slice, array)
	fmt.Println(n, slice)
}
func main() {
	f8()
	f7()
	f5()
	f4()
	f3()
	f2()
	f1()
	data2 := [][]int{
		{1, 2, 3},
		{100, 200},
		{11, 22, 33, 44},
	}
	fmt.Println(data2)
	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(data)
	fmt.Println(s)
	//1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	// 2.:=
	s2 := []int{}
	// 3.make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	// 4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包
	s6 = arr[1:4]
	fmt.Println(s6)
}
