package main

import (
	"fmt"
	"time"
)

func main() {
	through()
	//swith()
	//delete(nil, 1)
	//forSwitch()
}

func forSwitch() {
	var doneCh chan struct{}
	heartbeatCh := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-doneCh:
			return
		case <-heartbeatCh.C:
			fmt.Println("heartbeats")
			//return
		default:
			time.Sleep(time.Second)
			fmt.Println("default sleep")
		}
	}
}

func through() {
	var a = 1
	switch a {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("sdfsdf")
		fallthrough
	case 11:
		fmt.Println("1.5")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("third")
		fmt.Println("3")
	}
}

func swith() {
	var a = 1
	switch a {
	case 1:
		fmt.Println("coming...")
		break
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}
}
