package main

import (
	"fmt"
	"os"
)

func tel(ch chan int,quit chan bool) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- false
}

func main() {
	ch := make(chan int)
	chquit := make(chan bool)

	var ok = true
	go tel(ch,chquit)
	for ok {
		select {
		case i := <-ch:
			{
				fmt.Println(i)
			}
			case ok = <-chquit:
				{
					fmt.Println("ok is ",ok)
					os.Exit(0)
				}
		}
	}
}
