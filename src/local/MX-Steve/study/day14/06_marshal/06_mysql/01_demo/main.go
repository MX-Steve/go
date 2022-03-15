package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type DbWorker struct {
	Dsn string
	Db  *sql.DB
}

type Roles struct {
	id         int
	name       string
	desc       string
	page_perms string
	created_at string
	del        int
}

func main() {
	dbw := DbWorker{Dsn: "root:YHhg9-5.*@tcp(121.43.41.139:3306)/zh_hotel?charset=utf8mb4"}
	dbTemp, err := sql.Open("mysql", dbw.Dsn)
	dbw.Db = dbTemp
	if err != nil {
		panic(err)
		return
	}
	defer dbw.Db.Close()
	// dbw.insertData()
	// dbw.deleteData()
	// dbw.editData()
	dbw.queryData()
	dbw.transaction()
}

func (dbw *DbWorker) insertData() {
	stmt, _ := dbw.Db.Prepare(`INSERT INTO manages_role VALUES(?,?,?,?,?,?)`)
	defer stmt.Close()
	ret, err := stmt.Exec(1, "sss", "desc1", "[1,2]", "2022-03-14 09:00", 1)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); err == nil {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); err == nil {
		fmt.Println("RowsAffected: ", RowsAffected)
	}
}

func (dbw *DbWorker) deleteData() {
	stmt, _ := dbw.Db.Prepare(`DELETE FROM manages_role WHERE id=?`)
	ret, err := stmt.Exec(1)
	if err != nil {
		fmt.Printf("delete data error: %v\n", err)
		return
	}
	if RowsAffected, err := ret.RowsAffected(); err == nil {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

func (dbw *DbWorker) editData() {
	stmt, _ := dbw.Db.Prepare(`UPDATE manages_role SET del=? WHERE id=?`)
	ret, err := stmt.Exec(0, 1)
	if err != nil {
		fmt.Printf("edit data error: %v\n", err)
		return
	}
	if RowsAffected, err := ret.RowsAffected(); err == nil {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}
func (dbw *DbWorker) queryData() {
	stmt, _ := dbw.Db.Prepare(`SELECT * from manages_role where id=?`)
	defer stmt.Close()
	rows, err := stmt.Query(1)
	if err != nil {
		fmt.Printf("query data error: %v\n", err)
		return
	}
	columns, _ := rows.Columns()
	fmt.Println(columns)
	rowMaps := make([]map[string]string, 9)
	values := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))
	for i := range values {
		scans[i] = &values[i]
		scans[i] = &values[i]
	}
	i := 0
	for rows.Next() {
		_ = rows.Scan(scans...)
		each := make(map[string]string, 4)
		for i, col := range values {
			each[columns[i]] = string(col)
		}
		rowMaps = append(rowMaps[:i], each)
		fmt.Println(each)
		i++
	}
	fmt.Println(rowMaps)
	for i, col := range rowMaps {
		fmt.Println(i, col)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err.Error())
	}
}
func (dbw *DbWorker) transaction() {
	tx, err := dbw.Db.Begin()
	if err != nil {
		fmt.Printf("transaction data error: %v\n", err)
		return
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(`INSERT INTO manages_role VALUES(?,?,?,?,?,?)`)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	for i := 100; i < 110; i++ {
		cname := strings.Join([]string{"栏目-", string(i)}, "-")
		_, err := stmt.Exec(i, cname, "xxx", "[2,]", "2022-03-14 09:00:00", 0)
		if err != nil {
			fmt.Printf("insert data error: %v\n", err)
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	stmt.Close()
}
