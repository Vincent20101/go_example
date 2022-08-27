package main

import "fmt"

type OpenId struct {
	id int `json:"id"`
}

type Trip struct {
	name string `json:"name"`
	time string `json:"time"`
}
type Open struct {
	OpenId
	Trip
}

func main() {
	o := &Open{
		// 这里并没有name time id 字段
		OpenId{
			id: 1,
		},
		Trip{
			name: "test",
			time: "2022-12-1",
		},
	}
	o.name = "test2"
	fmt.Println(o)
	var a int
	b := make(map[string]string, 10)
	b["a"] = "b"
	defer func() {}()
	var f func()
	//f = nil
	//defer f()
	fmt.Println(f)
	if a == 0 {
		if s, ok := b["a"]; ok && a != 0 {
			fmt.Println(s)
		}
	}
}
