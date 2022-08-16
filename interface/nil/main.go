package main

import "fmt"

type iStruct struct {
	age int
	b   []byte
}

func main() {
	//var i *iStruct
	i := new(iStruct)
	fmt.Println(i)
	test()
}

// return 后 defer occur to panic
func test() (result *iStruct) {
	defer func() {
		result.age = 1
	}()
	return nil
}
