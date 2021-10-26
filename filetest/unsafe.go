package main

import "unsafe"

func main() {
	// string
	str1 := "abc"
	println("string1:", unsafe.Sizeof(str1)) // 16
	str2 := "abcdef"
	println("string1:", unsafe.Sizeof(str2)) // 16

	// arr
	arr1 := [...]int{1, 2, 3, 4}
	println("arr1:", unsafe.Sizeof(arr1)) // 32 = 8  * 4

	arr2 := [...]int{1, 2, 3, 4, 5}
	println("arr2:", unsafe.Sizeof(arr2)) // 32 = 8  * 5

	arr3 := [7]int{}
	println("arr3:", unsafe.Sizeof(arr3)) // 56 = 8  * 7

	// slice
	slice1 := []int{1, 2, 3, 4}
	println("slice1:", unsafe.Sizeof(slice1)) // 24

	slice2 := []int{1, 2, 3, 4, 5}
	println("slice2:", unsafe.Sizeof(slice2)) // 24

	slice3 := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
	println("slice3:", unsafe.Sizeof(slice3)) // 24

}
