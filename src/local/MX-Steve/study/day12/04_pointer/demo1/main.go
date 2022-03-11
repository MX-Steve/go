package main

import (
	"fmt"
	"runtime"
)

func main() {
	var a, b int = 20, 30
	var ptra *int
	var ptrb *int = &b
	ptra = &a
	fmt.Printf("a 地址: %x\n", &a)
	fmt.Printf("b 地址: %x\n", &b)
	fmt.Printf("ptra 的值: %x\n", ptra)
	fmt.Printf("ptrb 的值：%x\n", ptrb)
	fmt.Printf("*ptra 的值：%d\n", *ptra)
	fmt.Printf("*ptrb 的值：%d\n", *ptrb)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}
