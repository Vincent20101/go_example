package main

import "fmt"

type Mymap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

type Man struct {
}
type Woman struct {
}

type Company[T Man | Woman] struct {
	Name string
	CEO  T
}

type MyChannel[T int | string] chan T

//类型嵌套

type WowStruct[T string | int, S []T] struct {
	A T
	B S
}

//错误用法1, 类型参数不能单独使用
//type CommonType[T int | string] T

// 错误用法2
type CommonType[T interface{ *int } | string] []T

//匿名结构体不支持泛型
//泛型不支持switch判断

//匿名函数不支持泛型

type GenericSlice[T comparable] []T

func (s GenericSlice[T]) Contains(val T) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	//company := Company[Man]{
	//	Name: "bobby",
	//	CEO: Man{},
	//}

	//company := Company[Woman]{
	//	Name: "bobby",
	//	CEO:  Woman{},
	//}

	//var c MyChannel[string]

	//几种常见的错误
	intSlice := GenericSlice[int]{1, 2, 3}
	fmt.Println(intSlice.Contains(2)) // 输出 true

	strSlice := GenericSlice[string]{"a", "b", "c"}
	fmt.Println(strSlice.Contains("b")) // 输出 true
}
