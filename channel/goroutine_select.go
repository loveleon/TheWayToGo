package main

import (
	"fmt"

	"time"
)

func main() {
	//create channel
	ch1 := make(chan int)
	ch2 := make(chan int)



	go pump1(ch1)
	go pump2(ch2)
	go suck1(ch1, ch2)

	//time.Sleep(1e9)
	time.Sleep(1)

}

func pump1(ch chan int) {
	for i := 0; ; i-- {
		ch <- i
		//fmt.Println("i:",i)
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
		//fmt.Println("i:",i)
	}
}

func suck1(ch1 chan int, ch2 chan int) {
	for {
		select {
		case i := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", i)
		case i := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", i)
		//default:
		//	fmt.Printf("default.\n")
		}
	}
}
