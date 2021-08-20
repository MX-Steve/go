package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// open and close file
func file1() {
	file, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Println("open file failed,err:", err)
	} else {
		file.Close()
	}
}

// open and read file
func file2() {
	file, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	for {
		var tmp [128]byte
		// var s = make([]byte,0,128)
		_, err := file.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("read file finished.")
			return
		}
		if err != nil {
			fmt.Println("read from file failed, err: ", err)
			return
		}
		// fmt.Printf("read %d bytes\n", n)
		fmt.Println(string(tmp[:]))
	}
}

// bufio read data
func file3() {
	file, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		// use bufio to read data from file
		if err == io.EOF {
			fmt.Print(str)
			return
		}
		if err != nil {
			fmt.Println("read data failed.")
			return
		}
		fmt.Print(str)
	}

}

// ioutil read file
func file4() {
	content, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		fmt.Println("read file failed,err: ", err)
		return
	}
	fmt.Println(string(content))
}

// open file and write something
func file5() {
	file, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("open file failed,err: ", err)
	}
	defer file.Close()
	str := "Hello world. Hello, China.\n"
	file.Write([]byte(str))
	file.WriteString(str)
}

func main() {
	testIt := true
	if testIt {
		fmt.Println("starting.")
		file5()
	} else {
		fmt.Println("end.")
		file1()
		file2()
		file3()
		file4()

	}
}
