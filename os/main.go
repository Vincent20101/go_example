package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)

	open, err := os.Open("./ad.txt")
	fmt.Println(open)
	fmt.Println(err)
	fmt.Println(errors.Is(err, os.ErrNotExist))
	//if err != nil {
	//	fmt.Println(err)
	//	fmt.Println(errors.Is(err, os.ErrNotExist))
	//} else {
	//	fmt.Println(open)
	//}

}
