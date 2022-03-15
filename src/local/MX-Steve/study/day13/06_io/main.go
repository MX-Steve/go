package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fileObj, err := os.Open("../05_io/tt.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer fileObj.Close()
	// contents, err := ioutil.ReadAll(fileObj)
	// fmt.Println(string(contents))
	// if contents, _ := ioutil.ReadFile("../05_io/tt.txt"); err == nil {
	// 	fmt.Println(string(contents))
	// }
	// ioutil.WriteFile("./t3.txt", contents, 0666)
	fileObj, _ := os.OpenFile("./t3.txt", os.O_RDWR|os.O_CREATE, 0666)
	defer fileObj.Close()
	Rd := bufio.NewReader(fileObj)
	cont, _ := Rd.ReadSlice('#')
	fmt.Println(string(cont))
	Wr := bufio.NewWriter(fileObj)
	Wr.WriteString("WriteString writes a ## string.")
	Wr.Flush()
}
