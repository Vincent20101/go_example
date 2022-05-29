package main

import (
	"fmt"
)

func main() {
	dict := map[string]int{"王五": 60, "张三": 43}
	modify(dict)
	fmt.Println(dict["张三"])

	tMap := make([]int,0)
	fmt.Println(tMap)
}

//10

func modify(dict map[string]int) {
	dict["张三"] = 10
}
