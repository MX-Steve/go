package main

import (
	"errors"

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
	// sqlStr := "select id,title,price from book;"
	sqlStr := "select book.id,book.title,book.price,publisher.province,publisher.city,publisher.name from book join publisher  on book.publisher_id = publisher.id;"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		log.Warn("search all books err, ", err)
		return
	}
	return
}

// query data
func queryAllPublisher() (publisher []*Publisher, err error) {
	sqlStr := "select id,province,city,name from publisher;"
	err = db.Select(&publisher, sqlStr)
	if err != nil {
		log.Error("search all publisher err, ", err)
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

// insert data
func insertPublisher(province, city, name string) (err error) {
	sqlStr := "insert into publisher(province,city,name) values(?,?,?)"
	_, err = db.Exec(sqlStr, province, city, name)
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
func deletePublisher(id int64) (err error) {
	sqlStr := "delete from publisher where id=?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return
}

// modify data
func modifyBook(id int64, title string, price float64) (err error) {
	sqlStr := "update book set title=? , price=? where id=?"
	_, err = db.Exec(sqlStr, title, price, id)
	if err != nil {
		return err
	}
	return
}

// modify data
func modifyPublisher(id int64, province, city, name string) (err error) {
	sqlStr := "update publisher set province=? , city=?, name=? where id=?"
	_, err = db.Exec(sqlStr, province, city, name, id)
	if err != nil {
		return err
	}
	return
}
func getBookById(id int64) (book Book, err error) {
	sqlStr := "select id,title,price from book where id=?"
	err = db.Get(&book, sqlStr, id)
	if err != nil {
		return
	}
	return
}
func getPublisherById(id int64) (publisher Publisher, err error) {
	sqlStr := "select id,province,city,name from publisher where id=?"
	err = db.Get(&publisher, sqlStr, id)
	if err != nil {
		return
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
		log.Warn(err)
		return err
	}
	if !rows.Next() {
		return errors.New("username or password error")
	}
	log.Debug(username, password)
	return
}
