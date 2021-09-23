package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// init database
func initDB() (err error) {
	dsn := "root:kaituozhe520@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return
}

func createUser(username, password string) (err error) {
	sql := "insert into userinfo(username,password) values(?,?)"
	_, err = db.Exec(sql, username, password)
	if err != nil {
		fmt.Println("insert user error, ", err)
		return
	}
	return nil
}

func queryUser(username, password string) error {
	sqlStr := "SELECT id from userinfo WHERE username = ? and password = ?;"
	var id int64
	err := db.Get(&id, sqlStr, username, password)
	if err != nil {
		return err
	}
	return nil
}
