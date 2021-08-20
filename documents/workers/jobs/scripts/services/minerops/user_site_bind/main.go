package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type operatorConf struct {
	DatabaseURI string
	Sitekey     string
	UserName    string
}

const binS = "0123456789ABCDEF"
const senS = "0123456789QWERTYUIOPASDFGHJKLZXCVBNM"

func genUuid() string {
	total := ""
	len := 42
	for i := 0; i < len; i++ {
		ind := rand.Intn(16)
		total = fmt.Sprintf("%s%v", total, string(binS[ind]))
	}
	return total
}

func genSend() string {
	total := ""
	len := 36
	for i := 0; i < len; i++ {
		ind := rand.Intn(34)
		total = fmt.Sprintf("%s%v", total, string(senS[ind]))
	}
	return total
}

type machineOperator struct {
	conf *operatorConf
	db   *sql.DB
}

func create_site(DatabaseURI string, SiteKey string) {
	fmt.Println("Begin create site!")
	db, err := sql.Open("mysql", DatabaseURI)
	if err != nil {
		log.Fatal(err)
	}
	sql1 := "select name from site where name=?"
	rows, err := db.Query(sql1, SiteKey)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() == true {
		fmt.Println("site already exists.")
	} else {
		uid := genUuid()
		sql2 := "INSERT INTO site(site_key,name,description,encoded_key) VALUES(UNHEX(?),?,?,?)"
		_, err = db.Exec(sql2, uid, SiteKey, "", genSend())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func bind_host(DatabaseURI string, Sitekey string, mac string) {
	fmt.Println("Begin bind host and site")
	if mac == "" {
		log.Fatal("mac is required")
	}
	// sitekey: 1F7B169C846F218AB552FA82FBF86758BF5C97D2D2
	// uuid: 7943C74F6F944EF3A1E88A49C5D912FA
	// mac: 7C2BE108E645
	mac = strings.ReplaceAll(mac, ":", "")
	mac = strings.ToUpper(mac)
	db, err := sql.Open("mysql", DatabaseURI)
	if err != nil {
		log.Fatal(err)
	}
	sql0 := "select hex(site_key) from site where name=?"
	rows, err := db.Query(sql0, Sitekey)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	s0 := ""
	m0 := ""
	if rows.Next() {
		rows.Scan(&s0)
	}
	// SELECT hex(mac),hex(site_key),worker,card_number,card_type FROM machine_card mc WHERE mac = unhex('7C2BE108E645')
	sql1 := "select hex(mac) from machine where mac=unhex(?) and site_key=unhex(?)"
	rows, err = db.Query(sql1, mac, s0)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() {
		fmt.Println("mac already exists.")
		rows.Scan(&m0)
		fmt.Println(m0)
	} else {
		m := ""
		s := ""
		sql2 := "select hex(mac),hex(site_key) from machine where mac=unhex(?)"
		rows, err := db.Query(sql2, mac)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		if rows.Next() {
			rows.Scan(&m, &s)
			fmt.Println("mac in machine: ", m)
			fmt.Println("site_key in machine: ", s)
			sql3 := "update machine set site_key=unhex(?) where mac=unhex(?)"
			_, err := db.Exec(sql3, s0, m)
			if err != nil {
				fmt.Printf("update failed, err:%v\n", err)
			}
			fmt.Println("add mac funished")
		} else {
			fmt.Println("mac does not exists in machine.")
		}
	}
	sql10 := "update machine_card set site_key=unhex(?) where mac=unhex(?)"
	fmt.Println(s0)
	fmt.Println(m0)
	_, err = db.Exec(sql10, s0, m0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("machine_card modify ok")
}

func bind_site(DatabaseURI string, SiteKey string, UserName string) {
	fmt.Println("Begin bind site and user")
	db, err := sql.Open("mysql", DatabaseURI)
	if err != nil {
		log.Fatal(err)
	}
	sql1 := "select site_key from site where name=?"
	site_key := ""
	rows, err := db.Query(sql1, SiteKey)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		rows.Scan(&site_key)
	} else {
		log.Fatal("site_key does not exists.")
	}
	sql2 := "select uuid from user where name=?"
	uuid := ""
	rows, err = db.Query(sql2, UserName)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		rows.Scan(&uuid)
	} else {
		log.Fatal("UserId does not exists.")
	}
	sql3 := "select id from user_site where site_key=? and user_uuid=?"
	rows, err = db.Query(sql3, site_key, uuid)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		fmt.Println("already bind between user and site")
	} else {
		// SELECT id,hex(site_key),hex(user_uuid) FROM user_site us order by id desc limit 10;
		sql4 := "insert into user_site(id,site_key,user_uuid) values(null,?,?)"
		_, err := db.Exec(sql4, site_key, uuid)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("bind between user and site ok.")
	}
}

func main() {
	var user string
	var site string
	var ops string
	var mac string
	flag.StringVar(&user, "u", "", "minerops UserName")
	flag.StringVar(&site, "s", "", "minerops site name")
	flag.StringVar(&mac, "m", "", "minerops machine mac address")
	flag.StringVar(&ops, "t", "", "minerops operation type:create|bind_user|bind_host")
	flag.Parse()
	var conf = &operatorConf{
		DatabaseURI: "bminer:2WHvybnS5vHsrMJ2@tcp(54.187.56.18:3306)/bminer",
		Sitekey:     site,
		UserName:    user,
	}
	if ops == "create" {
		create_site(conf.DatabaseURI, conf.Sitekey)
	} else if ops == "bind_user" {
		bind_site(conf.DatabaseURI, conf.Sitekey, conf.UserName)
	} else if ops == "bind_host" {
		bind_host(conf.DatabaseURI, conf.Sitekey, mac)
	} else {
		fmt.Println("this option does not exists.")
	}
}
