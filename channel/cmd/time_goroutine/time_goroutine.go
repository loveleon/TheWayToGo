package main

import (
	time2 "time"
	"fmt"
)

func main(){
	time :=time2.Tick(1e8)//100ms
	boom :=time2.After(5e8)//500ms

	for {
		select {
		case <-time:
			fmt.Println("tick.")
		case <-boom:
			{
				fmt.Println("boom!")
				return
			}
		default:
			fmt.Println(".")
			time2.Sleep(1e7)//10ms
		}

	}
}
