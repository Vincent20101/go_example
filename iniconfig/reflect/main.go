package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	type Student struct {
		Age  int    `json:"age"`
		Name string `json:"name"`
		Sex  string `json:"sex"`
	}
	s := reflect.TypeOf((*Student)(nil))
	fmt.Println(s)
	fmt.Println(reflect.New(s))
	fmt.Println(reflect.New(s).Interface())

	str := `{"age": 1,"name": "test","sex": "男"}`

	i := reflect.New(s)
	if err := json.Unmarshal([]byte(str), i.Interface()); err != nil {
		fmt.Println(err)
	}
	fmt.Println(i.Elem().Interface())
	fmt.Println(i.Elem())
	fmt.Println(i.Elem().Addr())
	fmt.Println(i.Elem().Addr().Interface())
	fmt.Println("=======", i.Elem().Elem().Interface())
	i.Elem().Elem().FieldByName("Sex").SetString("女")
	student := i.Elem().Interface().(*Student)
	fmt.Println(student)
	fmt.Println(i.Elem().Interface().(*Student).Sex)

	s2 := i.Elem().Interface().(*Student)
	fmt.Printf("==== %#v\n", s2)
	//i.Elem().FieldByName("sex").SetString("女")
	fmt.Println(s2.Age)
}
