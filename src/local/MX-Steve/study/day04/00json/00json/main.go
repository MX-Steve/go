package main 

import (
	"fmt"
	"encoding/json"
	"strings"
	"io"
	"log"
)
/*
使用 stream 数据流的方式处理json字符串，适用于数据量比较大的情况
*/
func main(){
	const jsonStream = `
					{"Name":"Steve", "Text":"This is Steve!"}
					{"Name":"Lisi", "Text":"This is Lisi!"}
					{"Name":"Zhangsan", "Text":"This is Zhangsan!"}
					{"Name":"Wangwu", "Text":"This is Wangwu!"}
	`
	type Message struct {
		Name,Text string 
	}
	// func NewDecoder(r io.Reader) *Decoder
	// fmt.Println(strings.NewReader(jsonStream))
	// NewReader创建一个从s读取数据的Reader。
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	var allMes []*Message
	for {
		var m Message
		// Decode从输入流读取下一个json编码值并保存在v指向的值里
		if err := dec.Decode(&m); err == io.EOF {
			break
		}else if err != nil {
			log.Fatal(err)
		}
		allMes = append(allMes,&m)
		fmt.Printf("%s: %s\n",m.Name,m.Text)
	}
	for _,v := range allMes {
		fmt.Println(v.Name)
	}
}