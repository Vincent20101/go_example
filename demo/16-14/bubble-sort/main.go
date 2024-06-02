package main

import "log"

func BubbleSort(list []int) {
	n := len(list)

	// 进行 N-1 轮迭代
	for i := n - 1; i > 0; i-- {
		// 在一轮中有没有交换过
		hasSwapped := false

		// 每次从第一位开始比较，比较到第 i 位就不比较了，因为前一轮该位已经有序了
		for j := 0; j < i; j++ {
			// 如果前面的数比后面的大，那么交换
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				hasSwapped = true
			}
		}

		// 如果在一轮中没有交换过，那么已经排好序了，直接返回
		if !hasSwapped {
			return
		}
	}
}

func main() {
	list := []int{5, 3, 8, 2, 1}
	BubbleSort(list)
	log.Println(list)
}
