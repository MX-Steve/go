package main

/*
book manage platform
*/
type Book struct {
	Id    int     `db:"id"`
	Title string  `db:"title"`
	Price float64 `db:"price"`
	Publisher
}

type Word struct {
	Id      int    `db:"id"`
	Ename   string `db:"ename"`
	Cname   string `db:"cname"`
	Splits  string `db:"splits"`
	Memory  string `db:"memory"`
	Example string `db:"example"`
}

type Publisher struct {
	Id       int64  `db:"id"`
	Province string `db:"province"`
	City     string `db:"city"`
	Name     string `db:"name"`
}

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
