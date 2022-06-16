package main

import (
	"fmt"
	"sort"
)

type SortSliceTest struct {
	Num  int
	Name string
}

func main() {
	arrs := InitArrs()
	sort.Slice(arrs, func(i, j int) bool {
		return arrs[i].Num > arrs[j].Num
	})

	for k, v := range arrs {
		fmt.Println("index", k, "value", v)
	}

	var arr [8]int = [8]int{8, 3, 2, 9, 4, 6, 10, 0}

	// bubble_sort
	//for i := 0; i < len(arr); i++ {
	//	for j := 0; j < len(arr)-i-1; j++ {
	//		if arr[j] > arr[j+1] {
	//			arr[j], arr[j+1] = arr[j+1], arr[j]
	//		}
	//	}
	//}

	// select_sort
	//for i := 0; i < len(arr); i++ {
	//	for j := i + 1; j < len(arr); j++ {
	//		if arr[j] < arr[i] {
	//			arr[j], arr[i] = arr[i], arr[j]
	//		}
	//
	//	}
	//}

	// insert_sort
	//for i := 1; i < len(arr); i++ {
	//	for j := i; j > 0; j-- {
	//		if arr[j] < arr[j-1] {
	//			arr[j], arr[j-1] = arr[j-1], arr[j]
	//		} else {
	//			break
	//		}
	//		fmt.Println(arr)
	//	}
	//	fmt.Println("=========")
	//}

	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}

func InitArrs() (arrs []SortSliceTest) {
	arr1 := SortSliceTest{
		Num:  3,
		Name: "arr1",
	}

	arr2 := SortSliceTest{
		Num:  1,
		Name: "arr2",
	}

	arr3 := SortSliceTest{
		Num:  5,
		Name: "arr3",
	}

	arr4 := SortSliceTest{
		Num:  2,
		Name: "arr4",
	}

	arrs = append(arrs, arr1, arr2, arr3, arr4)
	return
}
