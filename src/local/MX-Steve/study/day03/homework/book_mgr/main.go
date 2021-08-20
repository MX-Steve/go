package main 

import (
	"fmt"
	"os"
)

// 使用函数实现一个简单的图书管理系统
// 每本书有书名，作者，价格，上架信息
// 用户可以在控制台添加数据，修改书籍信息，打印所有的书籍列表

type Book struct {
	name string
	author string 
	price float32
	update bool 
}

var (
	AllBooks []*Book
)
// var AllBooks = make([]*Book,0,200)
// 展示所有书籍
func showBooks(){
	fmt.Println("show book list")
	for _,book :=range AllBooks{
		fmt.Printf("name:%s ,author:%s ,price:%f ,update:%v;\n",book.name,book.author,book.price,book.update)
	}
}

// 实例化一个对象函数
func newBook(n string,a string,price float32,update bool) *Book{
	return &Book{
		name: n,
		author: a,
		price: price,
		update: update,
	}
}

// 添加书籍
func addBooks(){
	fmt.Println("add book info")
	var (
		name string 
		author string 
		price float32
		update bool
	)
	fmt.Println("请输入书名 name")
	fmt.Scanln(&name)
	fmt.Println("请输入作者 author")
	fmt.Scanln(&author)
	fmt.Println("请输入价格 price")
	fmt.Scanln(&price)
	fmt.Println("请输入是否上架 update")
	fmt.Scanln(&update)
	var book = newBook(name,author,price,update)
	AllBooks = append(AllBooks,book)
}

// 修改书籍
func modBooks(){
	fmt.Println("modify book info")
	var (
		name string 
		author string 
		price float32
		update bool
	)
	fmt.Println("请输入书名 name")
	fmt.Scanln(&name)
	var ex = false
	var choice = ""
	for _,book := range AllBooks{
		if name == book.name {
			fmt.Println("图书存在，可以修改")
			fmt.Printf("name:%s ,author:%s ,price:%f ,update:%v;\n",book.name,book.author,book.price,book.update)
			ex = true
			for {
				fmt.Println("请输入要修改的列,exit 退出 ")
				fmt.Scanln(&choice)
				switch choice {
				case "name":
					fmt.Println("请输入新值：")
					fmt.Scanln(&name)
					book.name = name
				case "author":
					fmt.Println("请输入新值：")
					fmt.Scanln(&author)
					book.author = author
				case "price":
					fmt.Println("请输入新值：")
					fmt.Scanln(&price)
					book.price = price
				case "update":
					fmt.Println("请输入新值：")
					fmt.Scanln(&update)
					book.update = update
				case "exit":
					goto loop			
				default:
					fmt.Println("列名不存在")
				}
			}
		}
	}
	loop:
		fmt.Println("修改当前图书完成")
	if !ex {
		fmt.Println("书名不存在，无法修改")
	}
}

func main(){
	var mes = "1 展示书籍\n2 添加书籍\n3 修改数据\n4 退出\n"
	var opt int 
	for {
		fmt.Println(mes)
		fmt.Scanln(&opt)
		switch opt{
		case 1:
			showBooks()
		case 2:
			addBooks()
		case 3:
			modBooks()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("输入错误，请输入")
		}
	}
}