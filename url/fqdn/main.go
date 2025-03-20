package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {
	// 创建一个带有超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	// 创建一个自定义的 Resolver
	resolver := &net.Resolver{
		PreferGo: true, // 使用 Go 的 DNS 解析器
	}

	// 使用 Resolver 查询 IP 地址
	ips, err := resolver.LookupIPAddr(ctx, "192.168.0.1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印查询结果
	for _, ip := range ips {
		fmt.Println(ip.IP)
	}
}
