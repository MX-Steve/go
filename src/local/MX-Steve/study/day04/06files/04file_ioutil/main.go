package main 

import (
	"fmt"
	"io/ioutil"
)

func main(){
	str := "hello 世界"
	err := ioutil.WriteFile("./xx.txt",[]byte(str),0666)
	if err != nil {
		fmt.Println("write file failed, err:",err)
		return
	}
}