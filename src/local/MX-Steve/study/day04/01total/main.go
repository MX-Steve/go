package main

import (
	"fmt"
	"time"
)

var today = "sunday"

const Week = 7

type Student struct {
	Name string
}

func init() {
	fmt.Println("init in main")
}

// timestamp
func timestampDemo(timestamp int64) {
	timeObj := time.Unix(timestamp, 360) // parameter 2: Nanosecond
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minite := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minite, second)
}

// tick
func tickDemo() {
	ticker := time.Tick(time.Second * 2)
	for i := range ticker {
		fmt.Println(i)
	}
}

// time format
func formatDemo() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("15:04 2006/01/02"))
}

func printTime(t time.Time) {
	timeStr := t.Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)
}

// record function time
func calcTime() {
	start := time.Now().UnixNano() / 1000
	fmt.Println("Start Time: ", start)
	time.Sleep(time.Microsecond * 30)
	end := time.Now().UnixNano() / 1000
	fmt.Println("End Time: ", end)
	fmt.Printf("spend time:%d ws\n ", end-start)
	start2 := time.Now()
	fmt.Println(time.Since(start2))
}

// Interface
type Cat struct{}

func (c Cat) Say() string { return "miao" }

type Dog struct{}

func (d Dog) Say() string { return "wang" }

type Sayer interface {
	Say() string
}

// interface: dry
type dryer interface {
	wash()
	dry()
}

type Haier struct {
	name  string
	price float64
	mode  string
}

func (h Haier) wash() {
	fmt.Println("haier wash")
}
func (h Haier) dry() {
	fmt.Println("haier dry")
}

func main() {
	var a dryer
	h1 := Haier{
		name:  "xxx",
		price: 28.99,
	}
	fmt.Println(h1)
	a = h1
	fmt.Println(a)
	a.dry()

	// var haier dryer
	// fmt.Println(haier)

	// var animalList []Sayer
	// c := Cat{}
	// d := Dog{}
	// animalList = append(animalList, c, d)
	// fmt.Println(animalList)

	// c := Cat{}
	// fmt.Println(c.Say())

	// fmt.Println(today)
	// time package
	// time.Time
	// now := time.Now()
	// fmt.Printf("%#v", now)
	// fmt.Println(now.Year())
	// fmt.Println(now.Month())
	// fmt.Println(now.Date())
	// fmt.Println(now.Hour())
	// fmt.Println(now.Minute())
	// fmt.Println(now.Second())
	// fmt.Println(now.Unix())
	// fmt.Println(now.UnixNano())

	// timestampDemo(now.Unix() + 3600)

	// tickDemo()

	// formatDemo()

	// now := time.Now()
	// printTime(now)

	// calcTime()
}
