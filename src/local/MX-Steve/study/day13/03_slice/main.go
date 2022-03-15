package main

import (
	"fmt"
	"os"
	"sort"
)

type Peak struct {
	Name      string `json:"name"`
	Elevation int    `json:"elevation"`
}

func main() {
	peaks := []Peak{
		{"Aconcagua", 22838},
		{"Denali", 20322},
		{"Kilimanjaro", 19341},
		{"Mount Elbrus", 18510},
		{"Mount Everest", 29029},
		{"Mount Kosciuszko", 7310},
		{"Mount Vinson", 16050},
		{"Puncak Jaya", 16024},
	}
	sort.Slice(peaks, func(i, j int) bool { return peaks[i].Elevation > peaks[j].Elevation })
	fmt.Println(peaks)
	fmt.Println(os.Args)
}
