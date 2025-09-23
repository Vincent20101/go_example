package main

import (
	"encoding/json"
	"fmt"
	"bytes"

	"github.com/davecgh/go-spew/spew"
)

type net struct {
	IsTure bool `json:"isTrue,omitempty"`
}

func main() {
	var str = `
{
  "isTrue": true
}
`

	var n net
	if err := json.Unmarshal([]byte(str), &n); err == nil {
		fmt.Println("result:", spew.Sdump(n))
	} else {
		fmt.Println(err)
	}

	var a *string
	fmt.Println(a)
	js, err := json.MarshalIndent("", "", "   ")
	fmt.Println(js, err)
	fmt.Println("lhb", string(js))
	test(a)
	test(nil)

	var i interface{} = a
	fmt.Println(i)
	fmt.Println(i == nil)
	body := bytes.NewBuffer(nil)
	body.Write([]byte{})
	fmt.Println(body.Len())
}

func test(t any) {
	fmt.Println(t == nil)
	fmt.Println(t)

}
