package main 

import (
	"fmt"
	"encoding/json"
)

// 大写外部才能访问
// 结构体中大写开头表示可公开访问，小写只能当前文件夹访问
/*
type Student struct {
 	ID 			int	
 	Gender 		string
 	Name 		string
}
*/
// Student 是一个结构体
// 定义元信息： json tag
type Student struct {
	ID 			int				`json:"id"`
	Gender 		string			`json:"gender"`
	Name 		string			`json:"name"`
}


// json 序列化
func main(){
	var stu1 = Student{
		ID: 1,
		Gender: "男",
		Name: "李四",
	}
	v,err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("Json Format Error")
		fmt.Println(v)
	}
	fmt.Println(v)  // []byte
	fmt.Println(string(v)) // 把 []byte转换成 string
	fmt.Printf("%#v",string(v))

	str := string(v)
	var stu2 = &Student{}
	json.Unmarshal([]byte(str),stu2)
	fmt.Println(stu2)
	fmt.Printf("%T\n",stu2)
}