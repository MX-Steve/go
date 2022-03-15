package main

import (
	"log"
	"os"
	"text/template"
)

var report = template.Must(template.ParseFiles("tmp.txt"))

type Book struct {
	Name  string
	Price float64
}

func main() {
	Data := []Book{
		{"《三国演义》", 19.82},
		{"《红楼梦》", 9.82},
		{"《西游记》", 29.82},
		{"《世纪》", 39.82},
	}
	if err := report.Execute(os.Stdout, Data); err != nil {
		log.Fatal(err)
	}
}
