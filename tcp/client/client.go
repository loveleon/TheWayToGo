package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
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
	clientName,_ := inputReader.ReadString('\n')
	//fmt.Println("client name ",clientName)
	trimmedClient := strings.Trim(clientName,"\r\n")

	//send to server until quit
	for {
		fmt.Println("What send to server ? Q for quit.")
		input,_ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input,"\r\n")
		if trimmedInput == "Q" {
			return
		}
		_,err := conn.Write([]byte(trimmedClient+"says:"+trimmedInput))
		if nil != err {
			fmt.Println("conn write error.")
		}

	}



}
