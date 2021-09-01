package main

import (
	"fmt"
	"reflect"
)

// interface contains type and value
func interface01() {
	var x interface{}
	var a int64 = 100
	var b int32 = 10
	var c int8 = 1
	x = a
	fmt.Println(x, b, c)
	x = false
	fmt.Println(x)
	// type assert
	// value, ok := x.(int)
	// if ok {
	// 	fmt.Printf("%v", value)
	// } else {
	// 	fmt.Println("x is not type of int")
	// }
	x = 12.34
	value, ok := x.(string)
	fmt.Printf("ok:%t  value:%#v  value type:%T\n", ok, value, value)
}

// reflect: the type of interface value
func reflectType(x interface{}) {
	// TypeOf
	t := reflect.TypeOf(x) // Get dynamic type information
	// fmt.Printf("type:%v\n", v)
	// fmt.Printf("%T\n", x)
	fmt.Printf("name:%v; kind:%v\n", t.Name(), t.Kind())
}

// reflect: the value of reflect.ValueOf
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Float32:
		fmt.Printf("type is float32, value is %.2f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %.2f\n", float64(v.Float()))
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	}
}

// modifyValue: modify value by reflect
func modifyValue(x interface{}) {
	v := reflect.ValueOf(x)
	kind := v.Kind()
	if kind == reflect.Ptr {
		// v.Elem().SetFloat(12.34)
		v.Elem().SetInt(36)
	}
}

// Cat struct
type Cat struct {
	name string
}

// Person struct
type Person struct {
	name string
	age  uint8
}

// Student struct
type Student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func structReflect() {
	stu1 := Student{
		Name:  "little prince",
		Score: 90,
	}
	t := reflect.TypeOf(stu1)
	// fmt.Println(t.Name(), t.Kind())
	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		fmt.Printf("name:%s;index:%d;type:%v;json tag:%v\n", filed.Name, filed.Index, filed.Type, filed.Tag.Get("json"))
	}
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s;type:%v;json tag:%v\n", scoreField.Name, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

// Study: method
func (s Student) Study() string {
	msg := "good good study , day day up"
	return msg
}

// Sleep: method
func (s Student) Sleep() string {
	msg := "sleep...."
	return msg
}

// printMethod
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	var test = true
	if test {
		fmt.Println("start...")

	} else {
		fmt.Println("end...")
		interface01()
		reflectType(100)
		reflectType(3.14)
		reflectType(false)
		reflectType("helloworld")
		reflectType([3]int{1, 2, 3})
		reflectType(map[string]int{})
		reflectType(Cat{name: "huahua"})
		reflectType(&Person{name: "lisi", age: 18})
		var i int32 = 100
		reflectType(i)
		var j float64 = 2.34
		reflectType(&j)
		reflectType([]int{11, 22, 33})
		// reference type : t.Name() is none
		// value type : t.Name() is value type
		var x int64 = 66
		reflectValue(x)
		var y float32 = 2.36
		reflectValue(y)
		c := reflect.ValueOf(10)
		fmt.Printf("type c:%T\n", c)
		var a int64 = 100
		modifyValue(&a)
		fmt.Println(a)

		structReflect()

		var s = Student{
			Name:  "lisi",
			Score: 98,
		}
		printMethod(s)
	}
}
