package mylog 

import (
	"os"
	"fmt"
	"time"
	"io"
)

// copyFile 拷贝文件函数
func copyFile(dstName, srcName string)(written int64,err error){
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil{
		fmt.Printf("open %s failed, err:%v.\n",srcName,err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n",dstName,err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) // 调用 io.Copy() 拷贝内容
}

// cutLog 日志大于 10M 则切割日志
func cutLog(fileName string){
	topSize := 10 * 1024 * 1024
	file, err := os.Stat(fileName)
	if err != nil {
		panic(fmt.Sprintf("open file:%s failed", fileName))
	}
	fileSize := file.Size()
	if int(fileSize) >= int(topSize) {
		nowStr := time.Now().Format("20060102_150405")
		newFileName := fmt.Sprintf("%s-%s",fileName,nowStr)
		_,err := copyFile(newFileName,fileName)
		if err != nil{
			panic(fmt.Sprintf("error: %s",err))
		}
		os.Remove(fileName)
	}
}