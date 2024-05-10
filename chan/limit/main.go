package main

import (
	"fmt"
	"time"
)

var datas []int

func main() {
	var a, b, c, d, e, f = 1, 2, 3, 4, 5, 6

	if a <= b || a > b && b > c {
		fmt.Println(d, e, f)
	}
}

func do() {
	ch := make(chan int, 500)
	for {
		select {
		case data := <-ch:
			datas = append(datas, data)
			func() {
				for {
					select {
					case data := <-ch:
						if datas = append(datas, data); len(datas) >= 21 /*config.Size*/ {
							//next
							return
						}
					case <-time.After(time.Second * 5):
						//next
						return
					}
				}
			}()
			if len(datas) > 0 {
				send(datas)
			}
		}
	}
}

func send(datas []int) {

}
