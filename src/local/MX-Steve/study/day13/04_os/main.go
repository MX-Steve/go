package main

import (
	"fmt"
	"os"
	"time"
)

func demo01() {
	// 1. os.StartProcess
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	Pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", Pid)
}

func signalListen() {
	c := make(chan os.Signal)
	// signal.Notify(c, syscall.SIGUSR2)
	for {
		s := <-c
		fmt.Println("get signal: ", s)
	}
}

func main() {
	// demo01()
	go signalListen()
	for {
		time.Sleep(10 * time.Second)
	}
}
