package main

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDb() (err error) {
	dns := "root:kaituozhe520@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Connect("mysql", dns)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return nil
}

func Login(username, password string) (err error) {
	sqlStr := "select id from user where username=? and password=?"
	rows, err := db.Query(sqlStr, username, password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if !rows.Next() {
		return errors.New("username or password error")
	}
	fmt.Println(rows.Next())
	return
}
