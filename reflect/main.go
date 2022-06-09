package main

import (
	"fmt"
	"math/big"
	"reflect"
)

func main() {
	type age *int
	typeOf := reflect.TypeOf((*age)(nil))
	value := reflect.New(typeOf)
	fmt.Println(value.Elem().Interface())
	fmt.Println(&typeOf)
	of := reflect.TypeOf(new(big.Int))
	fmt.Println(of)
	fmt.Println(1 << 23)
}
