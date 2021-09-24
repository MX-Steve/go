package main

/*
book manage platform
*/
type Book struct {
	Id    int     `db:"id"`
	Title string  `db:"title"`
	Price float64 `db:"price"`
}

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
