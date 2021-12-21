package main

import "fmt"

func main() {
	i := new(int)
	a := 1
	i = &a
	fmt.Println(*i)

	c := make(chan int)
	go func() {
		fmt.Println(<-c)
	}()
	c <- 133

	d := make(chan int)
	fmt.Println(d)

	//s := new([]int)
	//s = append(s,1)

	var s2 *[]int
	//s2 = new([]int)
	fmt.Println(s2)
}
