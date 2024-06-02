package main

import (
	"log"
	"quick-sort/stack"
)

// QuickSort 普通快速排序
func QuickSort(array []int, begin, end int) {
	if begin < end {
		// 进行切分
		loc := partition(array, begin, end)
		// 对左部分进行快排
		QuickSort(array, begin, loc-1)
		// 对右部分进行快排
		QuickSort(array, loc+1, end)
	}
}

// 切分函数，并返回切分元素的下标
func partition(array []int, begin, end int) int {
	i := begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
	j := end       // array[end]是数组的最后一位

	// 没重合之前
	for i < j {
		if array[i] > array[begin] {
			array[i], array[j] = array[j], array[i] // 交换
			j--
		} else {
			i++
		}
	}

	/* 跳出 for 循环后，i = j。
	 * 此时数组被分割成两个部分   -->  array[begin+1] ~ array[i-1] < array[begin]
	 *                        -->  array[i+1] ~ array[end] > array[begin]
	 * 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置。
	 * 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到最后i = j不满足条件就退出！
	 */
	if array[i] >= array[begin] { // 这里必须要取等“>=”，否则数组元素由相同的值组成时，会出现错误！
		i--
	}

	array[begin], array[i] = array[i], array[begin]
	return i
}

// 非递归快速排序
func QuickSort2(array []int) {
	
	st := stack.ArrayStack{}

	// 第一次初始化栈，推入下标0，len(array)-1，表示第一次对全数组范围切分
	st.Push(len(array) - 1)
	st.Push(0)

	// 栈非空证明存在未排序的部分
	for !st.IsEmpty() {
		// 出栈，对begin-end范围进行切分排序
		begin := st.Pop() // 范围区间左边
		end := st.Pop()   // 范围

		// 进行切分
		loc := partition(array, begin, end)

		// 右边范围入栈
		if loc+1 < end {
			st.Push(end)
			st.Push(loc + 1)
		}

		// 左边返回入栈
		if begin < loc-1 {
			st.Push(loc - 1)
			st.Push(begin)
		}
	}
}

func main() {
	list := []int{3}
	QuickSort(list, 0, len(list)-1)
	log.Println(list)

	list1 := []int{3, 8}
	QuickSort(list1, 0, len(list1)-1)
	log.Println(list1)

	list2 := []int{3, 8, 5}
	QuickSort(list2, 0, len(list2)-1)
	log.Println(list2)

	list3 := []int{3, 8, 5, 1, 9, 6}
	QuickSort(list3, 0, len(list3)-1)
	log.Println(list3)

	list4 := []int{3, 8, 5, 1, 9, 6}
	QuickSort2(list4)
	log.Println(list4)
}
