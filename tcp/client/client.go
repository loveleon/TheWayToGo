package main

import (
	"fmt"
	"net"
	"time"
	"bufio"
	"os"
)

//func main() {
//	fmt.Println("connect to server .")
//	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.01:50000")
//	if nil != err {
//		fmt.Println("ResolveTcpAddr error.")
//		return
//	}
//	conn, err := net.DialTCP("tcp", nil, addr)
//	if nil != err {
//		fmt.Println("DialTCP error.")
//		return
//	}
//	go doSendData(conn)
//
//}
//
//func doSendData(conn net.Conn) {
//	if conn == nil {
//		fmt.Println("client::doSendData parameter error.conn is equal to nil. ")
//		return
//	}
//
//	for i := 0; i < 3; i++ {
//		str := "hello"
//		conn.Write([]byte(str))
//		time.Sleep(1e8)
//	}
//}

func main(){
	//dial
	conn,err :=net.Dial("tcp","localhost:50000")
	if nil != err{
		fmt.Println("Dial error.")
		return
	}

	inputReader :=bufio.NewReader(os.Stdin)
	fmt.Println("what's your(client) name :")
	inputReader.ReadLine("\n")
}
