package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// read log configeration file
type Config struct {
	FilePath string `conf:"file_path"`
	FileName string `conf:"file_name"`
	MaxSize  int64  `conf:"max_size"`
}

func parseConf(fileName string, result interface{}) (err error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = fmt.Errorf("open configeration file err:%s ", fileName)
		return
	}
	lineSlice := strings.Split(string(data), "\n")
	for index, line := range lineSlice {
		line = strings.TrimSpace(line)
		// ignore null and comment
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		equalIndex := strings.Index(line, "=")
		if equalIndex == -1 {
			err = fmt.Errorf("paragram error: %d line ", index+1)
			return
		}
		key := strings.TrimSpace(line[:equalIndex])
		value := strings.TrimSpace(line[equalIndex+1:])
		if len(key) == 0 {
			err = fmt.Errorf("paragram error: %d line ", index+1)
			return
		}
		// result:-> ptr, use reflect to write
		t := reflect.TypeOf(result)
		v := reflect.ValueOf(result)
		if t.Kind() != reflect.Ptr {
			err = errors.New("result must be a Ptr")
			return
		}
		tElem := t.Elem()
		if tElem.Kind() != reflect.Struct {
			err = errors.New("result must be a Struct Ptr")
			return
		}
		for i := 0; i < tElem.NumField(); i++ {
			field := tElem.Field(i)
			tag := field.Tag.Get("conf")
			if key == tag {
				fieldType := field.Type
				switch fieldType.Kind() {
				case reflect.String:
					fieldValue := v.Elem().FieldByName(field.Name)
					fieldValue.SetString(value)
				case reflect.Int64, reflect.Int, reflect.Int32:
					value64, _ := strconv.ParseInt(value, 10, 64)
					v.Elem().Field(i).SetInt(value64)

				}
			}
		}
	}
	return
}

func main() {
	// 1. open configeration file

	// 2. read data
	// 3. read data line by line,get data from tag
	var c = &Config{}
	err := parseConf("xxx.conf", c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
	fmt.Println(c.FileName)
}
