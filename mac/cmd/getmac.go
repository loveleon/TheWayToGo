package main

import (
	"net"
	"fmt"
)

func main(){
	interfaces,err := net.Interfaces()
	if nil != err {
		panic(err.Error())
	}
	for _,inter := range interfaces{
		//fmt.Println(inter.Name,inter.HardwareAddr,inter.Index)
		fmt.Println(inter.HardwareAddr)
	}
}
