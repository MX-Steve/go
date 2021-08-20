package main 

import (
	"fmt"
	"os"
	"bufio"
)

// 打开文件支持文件写入
func main(){
	file, err := os.OpenFile("xx.txt",os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0755)
	if err != nil {
		fmt.Println("open file failed, err: ",err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i :=0;i<10;i++{
		writer.WriteString("hello世界\n")
		// 将数据先写入缓存
	}
	writer.Flush() // 将缓存中的内容写入文件
}