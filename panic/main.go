package main

import (
	"encoding/json"
	"sync"
)

type Test struct {
	Str string `json:"str"`
}

func main() {
	//test1()
	test2()
}

func test2() {
	s := "hello"
	p1 := &s
	*p1 = ""
	wg := sync.WaitGroup{}
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if i%2 == 0 {
				*p1 = "a"
			} else {
				*p1 = ""
			}
			json.Marshal(p1)
		}()
	}
	wg.Wait()
}
func test1() {
	t := &Test{
		Str: "123",
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if i%2 == 0 {
				t.Str = "99"
			} else {
				t.Str = ""
			}
			json.Marshal(t)
		}()
	}
	wg.Wait()
}
