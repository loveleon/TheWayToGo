package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main(){
	var conn net.Conn //Connect Struct
	var err error
	var inputReader *bufio.Reader //Buffered input
	var input string
	var clientname string

	//Dial to server
	conn,err = net.Dial("tcp","localhost:50001")
	if nil != err {
		return
	}

	//construct the buffer io
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("what's the client named:")
	clientname,_ = inputReader.ReadString('\n')
	//trim the client name
	trimmedClientName := strings.Trim(clientname,"\r\n")  //\r\n ---windows \n--linux


	//send loop
	for {
		fmt.Println("what's msg to be sent to server,TYPE Q for quit.")
		input,err = inputReader.ReadString('\n')
		//trim
		trimmedInput := strings.Trim(input,"\r\n")
		if "Q" == trimmedInput {
			return
		}
		_,err = conn.Write([]byte(trimmedClientName + " says " + trimmedInput))
		checkError(err)
	}
}

//checkError :check the function return err value.
func checkError(err error){
	if nil != err {
		panic("Error:" + err.Error())
	}
}
