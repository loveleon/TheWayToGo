package main

import (
	"fmt"
)

func tel(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	//fmt.Println("before close ch.")
	close(ch)
}

func main() {
	ch := make(chan int)
	var ok = true
	go tel(ch)
	for ok {
		if i, ok := <-ch; ok {
			fmt.Printf("ok is %t and the i is %d\n", ok, i)
		}
	}
}

