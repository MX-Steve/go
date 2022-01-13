package main

import "fmt"

type NewInt int
type MyInt = int
type person struct {
	name string
	city string
	age  int8
}
type student struct {
	name string
	age  int
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

type Person struct {
	name string
	age  int8
}

func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}
func main() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)
	var p1 person
	p1.name = "zhangsan"
	p1.city = "北京"
	p1.age = 28
	fmt.Println(p1)
	var user struct {
		Name string
		Age  int
	}
	user.Name = "pprof.cn"
	user.Age = 18
	fmt.Printf("%#v\n", user)
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "博客"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}
	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}
	p8 := &person{
		"pprof.cn",
		"北京",
		18,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"pprof.cn", city:"北京", age:18}
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	p9 := newPerson("pprof.cn", "测试", 90)
	fmt.Printf("%#v\n", p9)
	p10 := NewPerson("测试", 25)
	p10.Dream()
	p10.SetAge(30)
	fmt.Println(p10)
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}
