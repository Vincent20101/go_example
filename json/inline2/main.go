package main

import (
	"encoding/json"
	"fmt"
)

type T1 struct {
	FieldInt     int    `json:"field_int"`
	FieldIgnore  int    `json:"-"`                       //忽略
	FieldBooleab bool   `json:"field_boolean,string"`    //不同类型
	FieldString1 string `json:"field_string1,omitempty"` //忽略空值，当时复合结构时为要为指针类型
	FieldString2 string `json:"field_string2,omitempty"`
}
type T2 struct {
	T1 `json:",inline"`
	//T1
}
type T3 struct {
	T1 `json:"t1"`
}

func main() {
	val1 := T1{
		FieldInt:     11,
		FieldIgnore:  11,
		FieldBooleab: true,
		FieldString2: "no empty",
	}
	bty1, _ := json.Marshal(val1)
	fmt.Printf("%v\r\n", string(bty1))

	val2 := T2{
		val1,
	}
	bty2, _ := json.Marshal(val2)
	fmt.Printf("%v\r\n", string(bty2))

	val3 := T3{
		val1,
	}
	bty3, _ := json.Marshal(val3)
	fmt.Printf("%v\r\n", string(bty3))
}

// #程序输出
//{"field_int":11,"field_boolean":"true","field_string2":"no empty"}
//{"field_int":11,"field_boolean":"true","field_string2":"no empty"}
//{"t1":{"field_int":11,"field_boolean":"true","field_string2":"no empty"}}
