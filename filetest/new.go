package main

import "fmt"

func main() {
	i := new(int)
	//a := 1
	//i = &a
	fmt.Println(*i)

	c := make(chan int)
	go func() {
		fmt.Println(<-c)
	}()
	c <- 133

	d := make(chan int)
	fmt.Println(d)

}
