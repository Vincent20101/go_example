package main

import (
	"fmt"
	"reflect"
)

func main() {
	type age *int
	typeOf := reflect.TypeOf((*age)(nil))
	value := reflect.New(typeOf)
	fmt.Println(value.Elem().Interface())
	fmt.Println(&typeOf)
}
