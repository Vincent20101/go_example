package main

import (
	"encoding/json"
	"fmt"

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
}
