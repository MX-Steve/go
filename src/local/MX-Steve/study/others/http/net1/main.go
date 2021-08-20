/*
package main 

import (
	"net"
	"fmt"
	"bufio"
)

func main(){
	conn, err := net.Dial("tcp", "baidu.com:80")
	if err != nil {
		fmt.Println("handle error")
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("handle error")
	}
	fmt.Println(status)
}
*/
/*
package main 

import (
	"fmt"
	"net"
	"os"
)

func main(){
	addr, err := net.ResolveIPAddr("ip", "www.baidu.com")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    // fmt.Println(addr.IP)
    fmt.Println(addr)
    fmt.Printf("%T",addr)
}
*/
package main 

import (
	"fmt"
	"net"
)

func Read(con net.Conn){
	data := make([]byte, 1000)
	for{
		n,err := con.Read(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(data[0:n]))
	}
}

func main(){
	listen, err := net.Listen("tcp","192.168.226.130:3569")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		con, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go Read(con)
	}
}
// ]# go run main.go 挂起
// shell 发送TCP请求测试
// ]# exec 8<>/dev/tcp/192.168.226.130/3569
// ]# echo -e "stats" >&8
// ]# go run main.go
// stats
