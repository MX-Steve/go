package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
	Age  int `json:"years"`
}

func main() {
	// var x float64 = 3.45
	// fmt.Println(reflect.TypeOf(x))
	// fmt.Println(reflect.ValueOf(x))
	// var a int = 1000
	// v := reflect.ValueOf(a)
	// t := reflect.TypeOf(a)
	// fmt.Println(v, t, v.Type(), t.Kind())
	// var b [5]int = [5]int{11, 22, 33, 44}
	// fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind(), reflect.TypeOf(b).Elem())
	// var Pupil Student
	// p := reflect.ValueOf(Pupil)
	// fmt.Println(p.Type(), p.Kind())
	var a int = 50
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	fmt.Println(v, t, v.Type(), t.Kind(), reflect.TypeOf(&a).Elem())
	seta := reflect.ValueOf(&a).Elem()
	fmt.Println(seta, seta.CanSet())
	seta.SetInt(1000)
	fmt.Println(seta)

	var b [5]int = [5]int{5, 6, 7, 8}
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind(), reflect.TypeOf(b).Elem())

	var Pupil Student = Student{"joke", 50}
	p := reflect.ValueOf(Pupil)
	fmt.Println(p.Type())
	fmt.Println(p.Kind())
	setStudent := reflect.ValueOf(&Pupil).Elem()
	// setStudent.Field(0).SetString("Mike") 未导出字段，不能修改，panic 会发生
	setStudent.Field(1).SetInt(19)
	fmt.Println(setStudent)
	sSAge, _ := setStudent.Type().FieldByName("Age")
	fmt.Println(sSAge.Tag.Get("json"))
}
