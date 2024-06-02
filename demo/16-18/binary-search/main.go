package main

import "fmt"

func binarySearch(arr []int, target, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearch(arr, target, mid+1, high)
	} else {
		return binarySearch(arr, target, low, mid-1)
	}
}

func binarySearch2(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 5, 6, 8, 9}
	target := 5

	index := binarySearch(arr, target, 0, len(arr)-1)
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}

	index = binarySearch2(arr, target)
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
