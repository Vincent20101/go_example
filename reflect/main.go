package main

import (
	"fmt"
	"math/big"
	"net/url"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	//type age *int
	//age := new(int)

	type GetSmDataReq struct {
		UeId  string     `json:"ueId"`
		Query url.Values `json:"query"`
	}
	typeOf := reflect.TypeOf((*GetSmDataReq)(nil))
	fmt.Println("typeOf:", spew.Sdump(typeOf))
	value := reflect.New(typeOf)
	fmt.Println("elem:", spew.Sdump(value.Elem()))
	fmt.Println("interface:1", spew.Sdump(value.Interface()))
	fmt.Println("interface2:", spew.Sdump(value.Elem().Interface()))

	fmt.Println("========")
	fmt.Println("typeOf:", spew.Sdump(typeOf))
	fmt.Println("&typeOf", spew.Sdump(&typeOf))
	of := reflect.TypeOf(new(big.Int))
	fmt.Println("of:", spew.Sdump(of))
	fmt.Println(1 << 23)
}
