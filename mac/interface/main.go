package main

import (
	"fmt"
	"net"
)

func main() {
	infterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("get Interfaces failure:%v", err)
		return
	}
	for _, inft := range infterfaces {
		srcMac := inft.HardwareAddr
		fmt.Println(srcMac)
	}
}
