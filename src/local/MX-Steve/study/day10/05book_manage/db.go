package main

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// operate mysql functions

var db *sqlx.DB

func initDB() (err error) {
	dns := "root:kaituozhe520@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Connect("mysql", dns)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return nil
}

// query data
func queryAllBook() (bookList []*Book, err error) {
	sqlStr := "select id,title,price from book;"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("search all books err, ", err)
		return
	}
	return
}

// insert data
func insertBook(title string, price float64) (err error) {
	sqlStr := "insert into book(title,price) values(?,?)"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		return err
	}
	return
}

// delete data
func deleteBook(id int64) (err error) {
	sqlStr := "delete from book where id=?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return
}

// delete data
func modifyBook(id int64, title string, price float64) (err error) {
	sqlStr := "update book set title=? , price=? where id=?"
	_, err = db.Exec(sqlStr, title, price, id)
	if err != nil {
		return err
	}
	return
}

func getUser(username string) (ok bool, err error) {
	sqlStr := "select id from user where username=?"
	row, err := db.Query(sqlStr, username)
	if err != nil {
		return false, err
	}
	if !row.Next() {
		return false, err
	}
	return true, nil
}

// user register
func userRegister(username, password string) (err error) {
	ok, _ := getUser(username)
	fmt.Println(ok)
	if ok {
		return errors.New("error")
	}
	sqlStr := "insert into user(username,password) values(?,?)"
	_, err = db.Exec(sqlStr, username, password)
	if err != nil {
		return err
	}
	return
}

// user login
func userLogin(username, password string) (err error) {
	sqlStr := "select id from user where username=? and password=?"
	rows, err := db.Query(sqlStr, username, password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if !rows.Next() {
		return errors.New("username or password error")
	}
	fmt.Println(username, password)
	fmt.Println(rows.Next())
	return
}
