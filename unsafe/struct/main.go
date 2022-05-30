package main

import (
	"fmt"
	"go-example/unsafe/struct/person"
	"unsafe"
)

func main() {
	p := person.Programmer{"stefno", 18, "go"}
	fmt.Println(p)
	fmt.Println(unsafe.Sizeof(p))

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string(""))))
	*lang = "Golang"

	fmt.Println(p)
}
