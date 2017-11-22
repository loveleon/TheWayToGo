package main

import (
	"net"
	"fmt"
	"strings"
	"os"
	"bufio"
)

var mapUserName map[string]int
// Map of the clients: contains: clientname - 1 (active) / 0 - (inactive)

func main(){
	var listener net.Listener
	var err error
	var conn net.Conn

	mapUserName = make(map[string]int)
	fmt.Println("Start server ...")

	listener,err = net.Listen("tcp","localhost:50001")
	checkError(err)

	//accept loop
	for {
		conn,err = listener.Accept()
		checkError(err)
		go doServerStuff(conn)
	}

}

func doServerStuff(conn net.Conn){
	//read data from conn's buffer
	var buf []byte
	var err error

	//loop receive data from client
	for {
		//read the client data type
		buf = make([]byte,512)
		_,err = conn.Read(buf)
		checkError(err)

		input := string(buf)
		if strings.Contains(input,": SH") {
			fmt.Println("Server shutting down.")
			os.Exit(0)
		}

		//write out users
		if strings.Contains(input, ": WHO"){
			DisplayList()
		}

		index := strings.Index(input, "says")
		clintName :=input[0:index-1]
		mapUserName[string(clintName)] = 1
		fmt.Println("receive data :",string(buf))
	}
}

func DisplayList(){
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key,value := range mapUserName{
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("--------------------------------------------")
}

func checkError(err error){
	if nil != err {
		fmt.Println("error panic.")
		fmt.Println(err.Error())
		inputReader := bufio.NewReader(os.Stdin)
		inputReader.ReadString('\n')
		panic ("Error:" + err.Error())
	}
}