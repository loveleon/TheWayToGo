package main

import "fmt"

func main() {
	ch := make(chan string)
	go SendData(ch)
	GetData(ch)
}

func SendData(ch chan string) {
	ch <- "hello world."
	ch <- "how are you."
	ch <- "fine."
	ch <- "tkx,and you?"
	ch <- "stop the channel."
	close(ch)
}

func GetData(ch chan string) {
	//for loop get data
	for {
		input, ok := <-ch
		if !ok {
			fmt.Println("input is not ok.")
			break
		}
		fmt.Printf("input : %s\n", input)
	}
}
