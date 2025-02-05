package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {
	// 假设这是您的应用层数据
	data := []byte("GET / HTTP/1.1\r\nHost: example.com\r\n\r\n")

	// 创建数据包源
	packet := gopacket.NewPacket(data, layers.LayerTypeIPv4, gopacket.Default)

	// 获取应用层
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("应用层数据:")
		fmt.Println(string(applicationLayer.LayerContents()))
	} else {
		fmt.Println("未找到应用层数据")
	}
}
