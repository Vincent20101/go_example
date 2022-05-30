// Program to illustrate the usage of
// StorePointer function in Golang

// Including main package
package main

// importing fmt,
// sync/atomic and unsafe
import (
	"fmt"
	"reflect"
	"sync/atomic"
	"unsafe"
)

// Defining a struct type L
type L struct{ x, y, z int }

// Declaring pointer
// to L struct type
var PL *L

// Calling main
func main() {

	// Defining *addr unsafe.Pointer
	var unsafepL = (*unsafe.Pointer)(unsafe.Pointer(&PL))
	fmt.Println(*unsafepL)
	fmt.Println(&unsafepL)
	fmt.Println(unsafepL)

	// Defining value
	// of unsafe.Pointer
	var px L
	px = L{
		x: 1,
		y: 2,
		z: 3,
	}

	// Calling StorePointer and
	// storing unsafe.Pointer
	// value to *addr
	atomic.StorePointer(
		unsafepL, unsafe.Pointer(&px))

	pointer := (*L)(atomic.LoadPointer(unsafepL))
	fmt.Println("-------------")
	fmt.Println(pointer)
	//fmt.Println(*pointer) // 报错
	fmt.Println(&pointer)

	// Printed if value is stored
	fmt.Println("Val Stored!")

	type age *int
	typeOf := reflect.TypeOf((*age)(nil))
	value := reflect.New(typeOf)
	fmt.Println(value.Elem().Interface())
	fmt.Println(&typeOf)

}




