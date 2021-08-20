package main

import "fmt"

//pointer
func pointer1() {
	var a int
	fmt.Println(a)
	b := &a
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	c := "zhangsan"
	d := &c
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println(*d)
}

//array1
func modifyArray(a1 [3]int) {
	a1[0] = 100
}

//array2
func modifyArray2(a2 *[3]int) {
	a2[0] = 100
}

// make
func makeTest() {
	var a = new(int) // get a pointer of type int
	*a = 10
	fmt.Println(*a)
	var c = new([3]int)
	c[0] = 1
	fmt.Println(c)
	fmt.Println(*c)
}

// panic
func panicRecover() {
	//panic err
	defer func() {
		// recover
		err := recover()
		fmt.Println(err)
	}()
	var a []int
	a[0] = 100
}

// struct
type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

//struct func
func structTest() {
	// method 1
	var ls = student{
		name:   "lisi",
		age:    20,
		gender: "man",
		hobby:  []string{"basketball", "football"},
	}
	fmt.Println(ls)
	fmt.Println(ls.name)
	var wz = student{}
	fmt.Println(wz)
	// method 2
	var ww = new(student)
	ww.name = "王武"
	ww.age = 29
	ww.gender = "female"
	ww.hobby = []string{"male"}
	fmt.Println(ww)
	// method 3
	var nz = &student{}
	nz.name = "zznnn"
	fmt.Println(nz)
	// initialization without key
	var stu1 = student{
		"stu1",
		29,
		"male",
		[]string{"female"},
	}
	var stu2 = student{
		name: "stu2",
	}
	fmt.Println(stu1)
	fmt.Println(stu2)
}

// struct 2
func structTest2() {
	type test struct {
		a int8
		b int8
		c int8
	}
	var t = test{
		a: 1,
		b: 2,
		c: 3,
	}
	fmt.Println(&(t.a))
	fmt.Println(&(t.b))
	fmt.Println(&(t.c))
}

// structure is a value type
func struct3() {
	type stu struct {
		name string
		age  int8
	}
	var stu1 = student{
		name: "zs",
		age:  18,
	}
	stu2 := stu1
	stu2.name = "ww"
	fmt.Println(stu1)
	fmt.Println(stu2)
	stu3 := &stu1
	stu3.name = "hh"
	fmt.Println(stu1.name, stu2.name, stu3.name)
}

// constructed function
func newStudent(n string, age int, g string, h []string) *student {
	return &student{
		name:   n,
		age:    age,
		gender: g,
		hobby:  h,
	}
}

func test() {
	hs := newStudent("张三", 18, "male", []string{"female"})
	fmt.Println(hs)
	fmt.Println(hs.name)
}


// main
func main() {
	test()
	// struct3()
	// structTest2()
	// structTest()
	// panicRecover()
	// fmt.Println("in main")
	// makeTest()
	// pointer1()
	// a := [3]int{2, 3, 4}
	// modifyArray(a)
	// fmt.Println(a)
	// modifyArray2(&a)
	// fmt.Println(a)
}
