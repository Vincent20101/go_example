package main

import (
	"fmt"
	"unsafe"
)

type Human struct {
	sex  bool
	age  uint8
	min  int
	name string
}
type Student struct {
	//结构体成员存储在内存中 会按照数据类型中最大的数据的整数倍数进行存储
	name string //16字节
	sex  byte   //1字节  7
	age  int    //8字节
	addr string //16字节
	aa   int32 //4字节
	bb   int32 //4字节
}

func main() {
	h := Human{
		true,
		30,
		1,
		"hello",
	}
	var b bool = true
	var u uint8 = 3
	var a int =1
	var s string = "test"
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(u))
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(s))
	stu := Student{
		name: "test",
		sex:  1,
		age:  22,
		addr: "test addr",
		aa:   33,
		bb:   44,
	}
	fmt.Println(unsafe.Sizeof(stu))
	i := unsafe.Sizeof(h)
	j := unsafe.Alignof(h.age)
	k := unsafe.Offsetof(h.name)
	fmt.Println(i, j, k)
	fmt.Printf("%p\n", &h)
	var p unsafe.Pointer
	p = unsafe.Pointer(&h)
	fmt.Println(p)
}

//输出
//32 1 16
//0xc00000a080
//0xc00000a080
// 32：string 占16字节，所以16+16 =32；1 是因为age前是bool，占用1个字节；8是name的偏移是int 占8个字节