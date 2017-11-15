package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("create server listen at localhost:50000")
	//create listen port
	listener, err := net.Listen("tcp", "localhost:50000")
	if nil != err {
		fmt.Println("listen at localhost:50000 error.")
		return
	}

	//loop accept the connection.
	for {
		conn, err := listener.Accept()
		if nil != err {
			fmt.Println("accept connection error.")
			return
		}
		go doServer(conn)
	}

}

func doServer(conn net.Conn) {
	//loop read data
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if nil != err {
			fmt.Println("read error.")
			return
		}
		fmt.Println("read ", len, ",the content is :", string(buf))
	}
}
