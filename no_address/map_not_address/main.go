package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}

func main() {

	//建立结构体map
	s := make(map[int]person)
	//给map赋值
	s[1] = person{"tony", 20, "man"}
	fmt.Println(s[1])
	//修改map里结构体的成员属性的值

	//s[1].name = "tom"
	//会报 main\main.go:17:12: cannot assign to struct field s[1].name in map 的错误
	//x = y 这种赋值的方式，你必须知道 x的地址，然后才能把值 y 赋给 x。
	//但 go 中的 map 的 value 本身是不可寻址的，因为 map 的扩容的时候，可能要做 key/val pair迁移
	//value 本身地址是会改变的
	//不支持寻址的话又怎么能赋值呢

	fmt.Println(s[1].name)

	// modify
	s1 := make(map[int]*person)
	s1[1] = &person{"tony", 20, "man"}
	fmt.Println(s[1])
	s1[1].name = "tom"
	fmt.Println(s[1].name)

	//由刚刚得推断我们可以发现，只要知道了被修改值的地址，我们就可以修改它了
	//所以我们使用指针和引用保证每次赋值都可以找到地址
	//就可以实现 map 的结构体赋值了
}
