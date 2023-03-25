package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	path := "a.pcapng"
	// 使用pcap子包读取xxx.pcap数据包
	handle, err := pcap.OpenOffline(path)
	if err != nil {
		panic(err)
	}

	// 使用gopacket创建数据包源对象
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println(packetSource.Packets())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
		// 获取网络接口层
		ethLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethLayer != nil {
			fmt.Println("****** Ethernet layer detected ******")
			// 通过断言强制把Layer类型转换为Ethernet类型
			if eth, ok := ethLayer.(*layers.Ethernet); ok {
				// 解码相关信息,源和目的MAC地址
				fmt.Println("Src MAC: ", eth.SrcMAC)
				fmt.Println("Dst MAC: ", eth.DstMAC)
			}
		}

		// 获取IP层
		ip4Layer := packet.Layer(layers.LayerTypeIPv4)
		if ip4Layer != nil {
			fmt.Println("****** IPv4 layer detected ******")
			if ipv4, ok := ip4Layer.(*layers.IPv4); ok {
				// 解码相关信息,源和目的IP地址
				fmt.Println("Src IP: ", ipv4.SrcIP)
				fmt.Println("Dst IP: ", ipv4.DstIP)
				// 其他信息
				fmt.Println("Protocol: ", ipv4.Protocol)
				fmt.Println(ipv4.IHL)
				fmt.Printf("%x\n", ipv4.Checksum)
			}
		}

		// 获取传输层
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer != nil {
			fmt.Println("****** TCP layer detected ******")
			if tcp, ok := tcpLayer.(*layers.TCP); ok {
				fmt.Println("Src port: ", tcp.SrcPort)
				fmt.Println("Dst port: ", tcp.DstPort)
				fmt.Printf("%x\n", tcp.Checksum)
				fmt.Println(tcp.Seq)
			}
		}
	}
}
