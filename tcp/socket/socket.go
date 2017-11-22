package main

import (
	"fmt"
	"net"
)

func main() {
	var (
		host   = "www.apache.com"
		port   = "80"
		remote = host + ":" + port
		msg  string  = "GET /\n"
		data   = make([]uint8, 4096)
		read   = true
		count  = 0
	)

	//Dial
	conn,err := net.Dial("tcp",remote)
	if nil != err {
		fmt.Println("Dial error.")
	}
	conn.Write([]byte(msg))
	for read {
		count,err = conn.Read(data)
		read = (err == nil)
		fmt.Println(string(data[0:count]))
	}
	conn.Close()
}
