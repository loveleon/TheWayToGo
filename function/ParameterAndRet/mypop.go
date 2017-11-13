package main

import "fmt"

type Stack []int

func (st Stack) Pop() int {
	v := 0
	for ix := len(st) - 1; ix >= 0; ix-- {
		if v = st[ix]; v != 0 {
			fmt.Println(st[ix])
			st[ix] = 0
			return v
		}
	}
	return 0
}

func main() {
	st := Stack{1, 2, 3, 4, 5}
	st.Pop()
}
