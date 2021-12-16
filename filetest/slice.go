package main

import (
	"fmt"
	"time"
	"unsafe"
)

type MySlice []int

func (m MySlice) append(i int) {
	m = append(m, i)
	fmt.Println(m) // [0 1]

}

// receiver 是值不是指针时
//func main() {
//	var m MySlice = make(MySlice, 1, 1)
//	fmt.Println(m) // [0]
//	m.append(1)
//	fmt.Println(m) // [0]
//}

func Test(slice []int) []int {
	slice = append(slice, 1, 2, 3, 4, 5)

	//fmt.Println("demo函数：", slice)
	//slice[0] = 666
	//slice[1] = 777
	//slice[2] = 888
	//slice[3] = 999
	//slice[4] = 000
	return slice
}

//结果：
//demo函数： [1 2 3 4 5 1 2 3 4 5]
//[666 777 888 999 0 1 2 3 4 5]

func main0301() {
	slice := []int{1, 2, 3, 4, 5}

	slice = Test(slice)

	fmt.Println(slice)
}

func Demo(slice *[]int) {
	*slice = append(*slice, 1, 2, 3, 4, 5)
	fmt.Println("sdfsdf", &slice)
}
func main0302() {
	/*
	   //所有内存地址 是一个无符号十六进制整型数据表示的
	   type slice struct {
	      array unsafe.Pointer   //指针 8字节
	      len   int           //长度 8字节
	      cap   int           //容量 8字节
	   }
	*/
	time.Sleep(1e9)
	fmt.Println("1212121212")
	str := "sdfdsf"
	fmt.Println(&str)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice), "---", cap(slice))
	fmt.Println(unsafe.Sizeof(slice))
	Demo(&slice)
	fmt.Println(slice)
	fmt.Println(len(slice), "---", cap(slice))
	fmt.Println(unsafe.Sizeof(slice))

	fmt.Printf("%p\n", &slice)
	fmt.Println(slice)
}

//1212121212
//0xc000010200
//24
//sdfsdf 0xc00000e038
//24
//0xc00000c060
//[1 2 3 4 5 1 2 3 4 5]

func Test2(slice []int) {
	slice[0] = 1
	slice[1] = 1324
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(slice)

	fmt.Printf("the len is %d and cap is %d \n", len(slice), cap(slice))

	fmt.Printf("%p\n", &slice)
}

func main() {
	slice := make([]int, 5, 20)
	fmt.Printf("%p\n", &slice)

	Test2(slice)

	fmt.Println("我想看看：", slice)
	fmt.Printf("the len is %d and cap is %d \n", len(slice), cap(slice))
	fmt.Printf("%p\n", &slice)

}

func arrayPointer() {
	arrA := []int{100, 200}
	testARR(arrA)
	arrB := arrA[:]
	testARR(arrB)
	fmt.Printf("fun arr : %p, %v\n", arrA, arrA)

}

func testARR(i []int) {
	fmt.Printf("fun arr : %p, %v\n", i, i)
	i[1] = 100
}

func main0304() {
	arrayPointer()
}
