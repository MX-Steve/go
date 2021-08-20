package mtime

import (
	"fmt"
	"time"
)

// 命令执行耗时计时器
func CalcTime(f func()){
	start := time.Now()
	f()
	fmt.Println("总耗时：",time.Since(start))
}