package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Human struct{
	Name string `json:"name"`
}
func unmarshal2Struct(humanStr string) Human {
	h := Human{}
	err := json.Unmarshal([]byte(humanStr), &h)
	if err != nil {
		println(err)
	}
	return h
}

func main(){
	h := unmarshal2Struct(`{"name":"test12","strs":["a","b"]}`)
	fmt.Println(h.Name)

	whatAmI := func(i interface{}) {
		switch t:=i.(type) {
		case interface{}:
			fmt.Println("interface{}")
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("dont know type %T\n",t)
		}
	}
	whatAmI((true))
	whatAmI((1))
	whatAmI("hey")
	whatAmI(nil)

	var c *sync.Cond
	c = sync.NewCond(&sync.Mutex{})
	c.L.Lock()
	defer c.L.Unlock()

	i := []int{1,2,3,4,5,6,}
	i = i[1:]
	fmt.Println(i)

}