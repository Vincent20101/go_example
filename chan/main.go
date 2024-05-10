package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//test1()
	//stopCh := make(chan bool, 1)
	//go test2(stopCh)
	//time.Sleep(time.Second * 100)
	//fmt.Println("done")
	var wg sync.WaitGroup
	var rwlock sync.RWMutex
	go func() {
		defer wg.Done()
		for {
			rwlock.RLock()
			time.Sleep(time.Second)
			fmt.Println("dsf")
			defer rwlock.RUnlock()
		}
	}()
	type nchf_intf struct {
		chf_cfg     *int
		chf_client  *int
		chf_stats   *int
		chf_ip_addr string
	}
	n := nchf_intf{}
	go func() {
		time.Sleep(time.Second * 3)
		n.chf_ip_addr = "127.0.0.1"
	}()
	n.chf_ip_addr = "sdf"
	time.Sleep(time.Second * 4)
	fmt.Println(n)
}

func test2(stopCh chan bool) {
	var i int

	for {
		select {
		case <-stopCh:
			fmt.Println("stop loop")
			goto LOOP
		default:
			//do something
			fmt.Println("doing something...")
			time.Sleep(time.Second)
			i++
			if i == 5 {
				stopCh <- true
			}
		}
	}
LOOP:
	fmt.Println("111111111")
	return
}

func test1() {
	c1 := make(chan interface{})
	c2 := make(chan interface{})

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- nil
		//close(c1)
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- nil
		//close(c2)
	}()

	for i := 0; i < 20; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received1", msg1)
			fmt.Println(c1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			//default:
			//	fmt.Println("default")
		}
	}
}
