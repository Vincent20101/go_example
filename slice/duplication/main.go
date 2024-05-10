package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Work []*string
}

func main() {
	s1 := []int{1, 2, 3, 4, 5, 4, 4, 4, 5, 6, 7, 8, 9, 5, 4, 10, 11}
	dMap := map[int]struct{}{
		1:  {},
		13: {},
		4:  {},
		5:  {},
	}

	for i, v := range s1 {
		fmt.Println("lhb :", i)
		if _, ok := dMap[v]; ok {
			if i > len(s1) {
				//continue
			}
			// Bug
			s1[i] = s1[len(s1)-1]
			fmt.Println(len(s1))
			fmt.Println(s1)
			s1 = s1[:len(s1)-1]
			fmt.Println("s1", s1, "len", len(s1))
		}
	}

	fmt.Println(s1)

}

func stu() {
	s := Student{
		Work: make([]*string, 0),
	}
	for i := 0; i < 10; i++ {
		w := "work" + strconv.Itoa(i)
		s.Work = append(s.Work, &w)
		s.Work = append(s.Work, &w)
	}
	//for i := 5; i < 10; i++ {
	//	w := "work" + strconv.Itoa(i)
	//	s.Work = append(s.Work, &w)
	//}
	//fmt.Println(spew.Sdump(s))
	fmt.Println(s)
	// 调用去重函数
	removeDuplicates(&s)

	// 输出去重后的切片
	fmt.Println(s)
	for _, workPtr := range s.Work {
		fmt.Println(*workPtr)
	}
}

func removeDuplicates(s *Student) {
	// 使用 map 来跟踪唯一的字符串
	uniqueWork := make(map[*string]bool)

	// 创建新的切片
	var uniqueSlice []*string

	// 遍历原始切片
	for _, workPtr := range s.Work {
		// 如果字符串内容不存在于映射中，添加到新切片并标记为已见过
		if _, exists := uniqueWork[workPtr]; !exists {
			uniqueSlice = append(uniqueSlice, workPtr)
			uniqueWork[workPtr] = true
		}
	}

	// 将新切片的地址分配给原始切片
	s.Work = uniqueSlice
}
