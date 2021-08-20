package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("sbc")
	// fmt.Printf("|%08d|\n", 800)
	// f1 := 3.1415926535
	// fmt.Printf("%.2f\n", f1)

	// s1 := "这是一个字符串\""
	// fmt.Printf("%s\n", s1)
	// fmt.Printf("%q\n", s1)
	// fmt.Printf("%20s\n", s1)

	// // 数组
	// var a [3]int // Define an array of int type with a length of 3;
	// a = [3]int{3, 5, 8}
	// fmt.Println(a)
	// var c = [3]string{"Beijing", "Shanghai"}
	// fmt.Println(c)
	// var d = [...]int{11, 22, 33, 22, 3333, 4445, 66}
	// fmt.Println(d)
	// var e [20]int
	// e = [20]int{19: 5}
	// fmt.Println(e)
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	// for i, item := range a {
	// 	fmt.Println(i, item)
	// }
	// f := [...]int{1, 3, 5, 7, 8}
	// sum := 0
	// for _, item := range f {
	// 	sum += item
	// }
	// fmt.Println(sum)
	// g := [...]int{1, 2, 3, 4, 5, 6, 7}
	// for k, v := range g {
	// 	other := 8 - v
	// 	for i := k + 1; i < len(g); i++ {
	// 		if g[i] == other {
	// 			fmt.Println(k, i)
	// 		}
	// 	}
	// 	// for i, j := range g {
	// 	// 	if j == other {
	// 	// 		fmt.Println(k, i)
	// 	// 	}
	// 	// }
	// }

	// // Multidimensional array
	// var a [3]int
	// a = [3]int{11, 22, 33}
	// var b [3][2]int
	// b = [3][2]int{
	// 	[2]int{11, 22},
	// 	[2]int{33, 44},
	// }
	// fmt.Println(a)
	// fmt.Println(b)
	// var c = [3][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// }
	// fmt.Println(c)
	// fmt.Println(c[1][1])
	// for i := 0; i < len(c); i++ {
	// 	for j := 0; j < len(c[i]); j++ {
	// 		fmt.Println(i, c[i][j])
	// 	}
	// }

	// // Array is a value type
	// a := [2]int{1, 2}
	// b := a
	// b[0] = 100
	// fmt.Println(a)
	// fmt.Println(b)

	// // Slice
	// // declare slice 1: direct declaration
	// var a []int
	// fmt.Printf("%T\n", a)
	// var b = []int{1, 2, 3}
	// fmt.Printf("%T\n", b)
	// // declare slice 2: declare by array
	// a1 := [3]int{1, 2, 3}
	// var c []int
	// c = a1[0:1]
	// fmt.Println(c)
	// d := a1[0:3]
	// fmt.Println(d)

	// // size of slice
	// b := []int{33, 22, 334}
	// fmt.Println(len(b))
	// // capacity of slice
	// fmt.Println(cap(b))
	// x := [...]string{"Beijing", "Shanghai", "Nanning", "Chongqi", "Sichuan"}
	// y := x[1:4]
	// fmt.Println(y)
	// fmt.Printf("length of slice y: %d\n", len(y))
	// fmt.Printf("capacity of slice y: %d\n", cap(y))

	// // empty slice
	// var a = []int{}
	// fmt.Println(a, len(a), cap(a))
	// a = append(a, 1)
	// fmt.Println(a, len(a), cap(a))
	// a = append(a, 1)
	// fmt.Println(a, len(a), cap(a))
	// a = append(a, 1)
	// fmt.Println(a, len(a), cap(a))
	// fmt.Printf("ptr:%p\n", a)
	// a = append(a, 1)
	// fmt.Println(a, len(a), cap(a))
	// fmt.Printf("ptr:%p\n", a)
	// a = append(a, 1)
	// fmt.Println(a, len(a), cap(a))
	// fmt.Printf("ptr:%p\n", a)

	// // slice is a reference type
	// a := []int{1, 2, 3}
	// b := a
	// b[0] = 100
	// fmt.Println(a)
	// fmt.Println(b)
	// var c []int // memory hasn't been requested
	// c = make([]int, 3, 3)
	// copy(c, a)
	// fmt.Println(c)

	// // delete some element from the slice
	// a := []string{"Shanghai", "Beijing", "Guangzhou"}
	// a = append(a[:1], a[2:]...)
	// fmt.Println(a)

	// a := [...]int{1, 3, 5, 7, 9, 11, 13}
	// b := a[:]
	// b[0] = 100
	// fmt.Println(a[0])
	// c := a[:5]
	// fmt.Println(c)
	// d := c[:6]
	// fmt.Println(d)

	//
	// a := [...]int{1, 2}
	// b := a[:]

	// fmt.Printf("a: %p\n", &a)
	// fmt.Printf("b: %p\n", b)
	// b[0] = 100
	// fmt.Println(a, b)
	// make: used to initialize reference types
	// new: used to make value types
	// The slice is declared but can't be used without initialization
	// Initialize with make

	// // map
	// var m1 map[string]int
	// m1 = make(map[string]int, 10)
	// m1["age"] = 19
	// fmt.Println(m1)
	// m2 := map[string]string{
	// 	"name":   "lisi",
	// 	"school": "Beijing",
	// }
	// fmt.Println(m2)
	// value, ok := m2["name"]
	// if !ok {
	// 	fmt.Println("not exists")
	// } else {
	// 	fmt.Println(value)
	// }
	// for k, v := range m2 {
	// 	fmt.Println(k, v)
	// }
	// // delete k-v
	// delete(m2, "name")
	// fmt.Println(m2)

	// m2["home"] = "Shanghai"
	// m2["hobby"] = "basketball"
	// fmt.Println(m2)
	// var keys = make([]string, 0, 10)
	// for key := range m2 {
	// 	keys = append(keys, key)
	// }
	// fmt.Println(keys)
	// sort.Strings(keys)
	// fmt.Println(keys)
	// for _, k := range keys {
	// 	fmt.Println(k, m2[k])
	// }

	// mapSlice := make([]map[string]int, 3, 10)
	// fmt.Println(mapSlice)
	// mapSlice = append(mapSlice, map[string]int{"aaa": 100})
	// mapSlice = append(mapSlice, map[string]int{"aaa": 1000})
	// fmt.Println(mapSlice)
	// mapSlice[2] = map[string]int{"ccc": 220}
	// fmt.Println(mapSlice)

	// sliceMap := make(map[string][]int, 10)
	// sliceMap["lisi"] = make([]int, 3, 10)
	// sliceMap["lisi"][0] = 2
	// fmt.Println(sliceMap)
	// countNum("how are you doing ? you are ok ?")
	// sayHello()
	// total := add(24, 56)
	// fmt.Println(total)
	// add2(1)
	// add2(1, 2)
	// add2(1, 3, 4)
	// add2(1, 2, 3, 5)
	testDefer()
}

//defer
func testDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("finished")
}

func countNum(str string) {
	strSlice := strings.Split(str, " ")
	fmt.Println(strSlice)
	var keys map[string]int
	keys = make(map[string]int, 10)
	for _, item := range strSlice {
		_, ok := keys[item]
		if ok {
			keys[item] += 1
		} else {
			keys[item] = 1
		}
	}
	fmt.Println(keys)
	for k, v := range keys {
		fmt.Printf("%s:%d\n", k, v)
	}
}

func sayHello() {
	fmt.Println("hello world")
}

func sayHi(name string) {
	fmt.Printf("Hi,%s", name)
}

func add(a, b int) int {
	return a + b
}

//variable parameters
//There are no default parameters in go language
func add2(a int, b ...int) int {
	fmt.Println(b)
	return 0
}
