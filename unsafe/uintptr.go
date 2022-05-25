// Program to illustrate the usage of
// StoreUintptr function in Golang

// Including main package
package main

// importing fmt and sync/atomic
import (
	"fmt"
	"sync/atomic"
)

// Main function
func maino01() {

	// Defining variables for
	// the address to store the val
	var (
		x uintptr
		y uintptr
	)

	// Using StoreUintptr method
	// with its parameters
	atomic.StoreUintptr(&x, 444443)
	atomic.StoreUintptr(&y, 223)

	// Displays the value stored in addr
	fmt.Println(atomic.LoadUintptr(&x))
	fmt.Println(atomic.LoadUintptr(&y))
	fmt.Println(x,y)
	fmt.Println(&x,&y)
}





// Program to illustrate the usage of
// StoreUintptr function in Golang

// Main function
func main02() {

	// Defining variables for
	// the address to store the val
	var (
		x uintptr
	)

	// Using StoreUintptr method
	// with its parameters
	atomic.StoreUintptr(&x, 5255151111)

	// Loading the stored val
	z:= atomic.LoadUintptr(&x)

	// Prints true if values
	// are same else false
	fmt.Println(z == x)

	// Prints true if addresses
	// are same else false
	fmt.Println(&z == &x)
}