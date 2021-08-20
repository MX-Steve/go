package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"io/ioutil"
)

// bufio 读取数据
func readBufline(){
	f , err := os.Open("../01file_opens/xx.txt")
	if err != nil{
		fmt.Println("打开文件失败")
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n') // 字符
		if err == io.EOF {
			fmt.Println(str)
			return
		}
		if err != nil {
			fmt.Println("读取文件内容失败")
			return
		}
		fmt.Print(str)
	}
}

// ioutil 读取文件
func readFile(f string){
	content, err := ioutil.ReadFile(f)
	if err != nil{
		fmt.Println("文件打开失败")
		return
	}
	fmt.Println(string(content))
}

func main(){
	
	// 利用缓冲区从文件读取数据
	readBufline()
	// ioutil
	readFile("../01file_opens/xx.txt")

}