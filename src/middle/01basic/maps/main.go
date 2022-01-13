package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func m1() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	fmt.Println(scoreMap)
	fmt.Printf("type of a:%T\n", scoreMap)
}

func m2() {
	userInfo := map[string]string{
		"username": "zhangsan",
		"password": "1234",
	}
	fmt.Println(userInfo)
}

func m3() {
	scoreMap := make(map[string]int, 8)
	scoreMap["zhangsan"] = 80
	v, ok := scoreMap["zhangsan"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
func m4() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}
func m5() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	delete(scoreMap, "王五")
	fmt.Println(scoreMap)
}
func m6() {
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

}
func m7() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d, value:%v\n", index, value)
	}
	fmt.Println("after init")
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d, value:%v\n", index, value)
	}
}
func m8() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
func m9() {
	var mapInit = map[string]string{"xiaoli": "湖南", "xiaoliu": "天津"}
	fmt.Println(mapInit)
	var mapTemp map[string]string
	mapTemp = make(map[string]string, 10)
	mapTemp["xiaoming"] = "北京"
	mapTemp["xiaowang"] = "河北"
	v1, ok := mapTemp["xiaoming"]
	fmt.Println(ok, v1)
	if v2, ok := mapTemp["xiaowang"]; ok {
		fmt.Println(v2)
	}
	for k, v := range mapTemp {
		fmt.Println(k, v)
	}
	delete(mapTemp, "xiaowang")
	fmt.Println(mapTemp)
	fmt.Println(len(mapTemp))
}
func main() {
	m1()
	m2()
	m3()
	m4()
	m5()
	m6()
	m7()
	m8()
	m9()
}
