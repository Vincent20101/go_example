package main

import (
	"fmt"
	"sync"
	"time"
)

func init() {
	fmt.Println("initial")
}

func main() {
	dict := map[string]int{"王五": 60, "张三": 43}
	modify(dict)
	fmt.Println(dict["张三"])
	fmt.Println(dict)

	tMap := make([]int, 0)
	fmt.Println(tMap)

	t := Teacher{}
	t.ShowA()
	fmt.Println()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}

//var wg sync.WaitGroup
//
//func f1() {
//	time.Sleep(1 * time.Second)
//	wg.Done()
//}
//
//func main() {
//	var i int
//	for i = 0; i < 3; i++ {
//		wg.Add(1)
//		go f1()
//	}
//	wg.Wait()
//	fmt.Println("end...")
//}

//10

func modify(dict map[string]int) {
	dict["张三"] = 10
	dict["aa"] = 10

	//var x = nil
	//var x interface{} = nil
	//var x string = nil
	var x error = nil
	fmt.Println(x)
}

type People struct{}

func (p *People) ShowA() {
	fmt.Print("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Print("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Print("teacher showB")
}
