package main

import "fmt"

type Book struct {
	id     int
	title  string
	price  float32
	update bool
}

var AllBooks []Book

func newBook(id int, title string, price float32, update bool) *Book {
	return &Book{
		id:     id,
		title:  title,
		price:  price,
		update: update,
	}
}

//addBook
func addBook() {
	var (
		id     int
		title  string
		price  float32
		update bool
	)
	fmt.Printf("id:")
	fmt.Scanln(&id)
	fmt.Printf("title:")
	fmt.Scanln(&title)
	fmt.Printf("price:")
	fmt.Scanln(&price)
	fmt.Printf("update:")
	fmt.Scanln(&update)
	var book = newBook(id, title, price, update)
	AllBooks = append(AllBooks, *book)
}

//modifyBook
func modifyBook() {
	var (
		id     int
		title  string
		price  float32
		update bool
	)
	fmt.Printf("id:")
	fmt.Scanln(&id)
	fmt.Printf("title:")
	fmt.Scanln(&title)
	fmt.Printf("price:")
	fmt.Scanln(&price)
	fmt.Printf("update:")
	fmt.Scanln(&update)
	for k, item := range AllBooks {
		if item.id == id {
			AllBooks[k].title = title
			AllBooks[k].price = price
			AllBooks[k].update = update
		} else {
			fmt.Println("not exists,can not modify")
		}
	}
}

//listBook
func listBook() {
	fmt.Println("list")
	for _, item := range AllBooks {
		fmt.Println("id: ", item.id)
		fmt.Println("title: ", item.title)
		fmt.Println("price: ", item.price)
		fmt.Println("update: ", item.update)
		fmt.Println("++++++++++")
	}
}

//deleteBook
func deleteBook() {
	fmt.Println("delete")
	var id int
	fmt.Printf("id:")
	fmt.Scanln(&id)
	for k, item := range AllBooks {
		if item.id == id {
			AllBooks = append(AllBooks[:k], AllBooks[k+1:]...)
		}
	}
}

// menus
func menus() {
	exit := false
	for {
		fmt.Printf("1: 添加书籍\n2: 更新书籍\n3: 展示书籍\n4: 删除书籍\n5: 退出\n")
		var op int
		fmt.Scanln(&op)
		switch op {
		case 1:
			addBook()
		case 2:
			modifyBook()
		case 3:
			listBook()
		case 4:
			deleteBook()
		case 5:
			exit = true
		}
		if exit {
			break
		}

	}
}

// main
func main() {
	menus()
}
