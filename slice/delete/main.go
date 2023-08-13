package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ch = make(chan struct{})
	close(ch)
	type st struct {
		name string
		age  int
	}
	var stSlice = make([]map[int]*st, 0)
	for i := 0; i < 10; i++ {
		m := make(map[int]*st)
		m[i] = &st{
			name: strconv.Itoa(i),
			age:  i,
		}
		stSlice = append(stSlice, m)
	}
	fmt.Println(stSlice)
	for k, v := range stSlice {
		if v != nil && k == 0 {
			stSlice[k] = stSlice[len(stSlice)-1]
			stSlice = stSlice[:len(stSlice)-1]
		}
	}

	fmt.Println(stSlice)
}
