package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		var b bool
		for y := 0; y < 50; y++ {
			fmt.Println(i, "===", y)
			if i == y {
				b = true
				break
			}
		}
		if !b {
			fmt.Println(i)
		}
	}
}
