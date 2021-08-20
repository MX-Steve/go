package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// 解析日志库的配置文件

// Config 是一个日志配置项
type Config struct {
	filePath string `conf:"file_path"`
	fileName string `conf:"file_name"`
	maxSize  int64  `conf:"max_size"`
}

// 从 conf 文件中读取内容赋值给结构体指针
func parseConf(fileName string, result interface{}) (err error) {
	// 0. 前提条件，result 必须是一个 ptr
	t := reflect.TypeOf(result)
	// v := reflect.ValueOf(result)
	if t.Kind() != reflect.Ptr {
		err = errors.New("result 必须是一个指针")
		return
	}
	// 1. 打开配置文件
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = fmt.Errorf("打开配置文件%s失败", fileName)
		return err
	}
	// 2. 将读取的文件数据按行分割, 得到行切片
	lineSlices := strings.Split(string(data), "\n")
	// 3. 一行一行解析数据
	for index, line := range lineSlices {
		line = strings.TrimSpace(line) // 去除字符串收尾空白
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			// 是空行和注释
			continue
		}
		// 解析是不是正确配置项[行内有等号]
		equalIndex := strings.Index(line, "=")
		if equalIndex == -1 {
			err = fmt.Errorf("第%d行语法错误", index+1)
			return
		}
		// 按照 = 分割每一行，左边是key，右边是value
		key := line[:equalIndex]
		value := line[equalIndex+1:]
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if len(key) == 0 {
			err = fmt.Errorf("第%d行语法错误", index+1)
			return
		}
		// 利用反射，给 result 赋值
		// 遍历结构体的每一个字段和 key 比较，匹配上了就把 value 赋值
		//for i:=0;i<tE
	}
	return

}

func main() {

	// 2. 读取内容
	// 3. 一行一行读取内容，根据Tag找结构体对应字段
	// 4. 找到了要赋值

	var c = &Config{} // 用来存储读取的数据
	err := parseConf("xx.conf", c)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c)
}
