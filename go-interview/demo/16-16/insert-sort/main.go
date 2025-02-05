package main

import "log"

func InsertSort(list []int) {
	n := len(list)
	// 进行 N-1 轮迭代
	for i := 1; i <= n-1; i++ {
		current := list[i] // 待排序的数
		j := i - 1         // 待排序的数左边的第一个数的位置

		// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
		if current < list[j] {
			// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && current < list[j]; j-- {
				list[j+1] = list[j] // 某数后移，给待排序留空位
			}
			list[j+1] = current // 结束了，待排序的数插入空位
		}
	}
}

func main() {
	list := []int{5}
	InsertSort(list)
	log.Println(list)

	list1 := []int{5, 3}
	InsertSort(list1)
	log.Println(list1)

	list2 := []int{5, 3, 8, 2, 1}
	InsertSort(list2)
	log.Println(list2)
}
