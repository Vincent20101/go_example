package main

import (
	"fmt"
	"net"
)

func main() {
	// 获取默认路由的网络接口
	defaultRoute, err := net.InterfaceByName("enp1s0")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 获取网络接口的MTU值
	mtu := defaultRoute.MTU
	fmt.Println("MTU:", mtu)
}
