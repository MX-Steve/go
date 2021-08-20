package main

import (
	"fmt"
)

type Student struct {
	ID		int 			`json:"id"`
	Name	string 			`json:"name"`
	Age		int 			`json:"age"`
	Class	string			`json:"class"`
}

// 构造函数
func NewStudent(id int, name string, age int, class string) *Student{
	return &Student{
		ID: id,
		Name: name,
		Age: age,
		Class: class,
	}
}

var AllStudents []*Student

// 菜单展示
func showMenus(){
	msg := "Operations List\n 1. list students\n 2. add student\n 3. modify student\n 4. delete student"
	var ops int 
	for {
		fmt.Println(msg)
		fmt.Scanln(&ops)
		switch ops {
		case 1: 
			listStudents()
		case 2:
			addStudents()
		case 3:
			modStudent()
		case 4:
			delStudent()
		default:
			continue 
		}
	}
}

// 学员信息展示
func listStudents(){
	for i,stu := range AllStudents{
		fmt.Printf("%d: id: %d, name: %s, age: %d, class: %s\n",i,stu.ID,stu.Name,stu.Age,stu.Class)
	}
}

// 学员信息添加
func addStudents(){
	var (
		ID int 
		Name string 
		Age int 
		Class string
	)
	fmt.Println("请输入ID:")
	fmt.Scanln(&ID)
	fmt.Println("请输入Name:")
	fmt.Scanln(&Name)
	fmt.Println("请输入Age:")
	fmt.Scanln(&Age)
	fmt.Println("请输入Class:")
	fmt.Scanln(&Class)
	// var stu1 = &Student{
	// 	ID: ID,
	// 	Name: Name,
	// 	Age: Age,
	// 	Class: Class,
	// }
	var stu2 = NewStudent(ID,Name,Age,Class)
	AllStudents = append(AllStudents,stu2)
}

// 学员信息修改
func modStudent(){
	var (
		ID int 
		Name string 
		Age int 
		Class string
	)
	var ops string 
	fmt.Println("请输入要修改的学生变化ID:")
	fmt.Scanln(&ID)
	for _,stu := range AllStudents{
		if stu.ID == ID {
			fmt.Println("请输入要修改的列：Name/Age/Class")
			fmt.Scanln(&ops)
			switch ops {
			case "Name":
				fmt.Println("请输入新用户名")
				fmt.Scanln(&Name)
				stu.Name = Name
				return 
			case "Age":
				fmt.Println("请输入新年龄")
				fmt.Scanln(&Age)
				stu.Age = Age
				return 
			case "Class":
				fmt.Println("请输入新班级")
				fmt.Scanln(&Class)
				stu.Class = Class
				return 
			}
		}
	}
	fmt.Println("当前ID不存在，无法修改")
}

// 学员信息删除
func delStudent(){
	var ID int 
	fmt.Println("请输入要移除的学生ID")
	fmt.Scanln(&ID)
	for i, stu := range AllStudents {
		if stu.ID == ID {
			AllStudents = append(AllStudents[:i],AllStudents[i+1:]...)
		}
	}
	fmt.Println("当前ID不存在，无法删除")
}

func main(){
	showMenus()
}