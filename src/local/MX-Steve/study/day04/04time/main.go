package main

import (
	"fmt"
	"time"

	"local/MX-Steve/tools/mtime"
)

// 内置 time 包

// 时间戳
func timestampDemo() {
	now := time.Now()
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

// 使用 time.Unix() 函数将时间戳转为时间格式
func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	// fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间格式化
func formatDemo() {
	now := time.Now()
	// 格式化的模板为GO的出生时间2006年1月2日15点04分05秒
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("15:04:05 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

// 定时器: 使用 time.Tick(时间间隔)来设置定时器
func tickDemo() {
	ticker := time.Tick(time.Second * 2)
	for i := range ticker {
		fmt.Println(i)
	}
}

// 计时器1
func calcTime() {
	start := time.Now().UnixNano() / 1000
	fmt.Println("这是一个测试内容")
	time.Sleep(time.Millisecond * 30)
	end := time.Now().UnixNano() / 1000
	fmt.Printf("总耗时 %#v 微秒\n", end-start)
}

// 计时器2
func calcTime2() {
	start := time.Now()
	fmt.Println("这是一个测试内容")
	time.Sleep(time.Millisecond * 30)
	fmt.Println("总耗时：", time.Since(start))
}

func f1() {
	fmt.Println("开始 f1 函数")
	time.Sleep(time.Millisecond * 30)
	fmt.Println("结束 f1 函数")
}

func main() {
	// now := time.Now()
	// fmt.Printf("%#v\n",now)
	// fmt.Println(now.Year())
	// fmt.Println(now.Month())
	// fmt.Println(now.Day())
	// fmt.Println(now.Hour())
	// fmt.Println(now.Minute())
	// fmt.Println(now.Second())
	// fmt.Println(now.Nanosecond())

	// timestampDemo()

	// now := time.Now().Unix()
	// timestampDemo2(now)

	// formatDemo()

	// tickDemo()

	// calcTime()
	// calcTime2()

	mtime.CalcTime(f1)
}
