package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	fmt.Println("hello module")
	dsn := "root:kaituozhe520@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
}
