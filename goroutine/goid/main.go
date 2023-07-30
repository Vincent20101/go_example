package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func main() {
	//fmt.Println(runtime.Version())
	fmt.Println(GetGoid())
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(GetGoid())
		}()
	}
	//panic("gold")
	wg.Wait()
}

func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)
	//fmt.Println(stk)
	//fmt.Println("===")
	idField := strings.Fields(stk)[0]
	//fmt.Println(idField)
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}
	return int64(id)

}
