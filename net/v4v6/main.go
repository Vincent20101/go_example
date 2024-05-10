package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// IPv4 地址
	ipv4Addr := net.ParseIP("192.0.2.1")

	// IPv4 转 IPv6 过渡地址
	ipv6Addr := ipv4Addr.To16()
	if ipv6Addr == nil {
		fmt.Println("IPv4 地址转换为 IPv6 失败")
		return
	}

	fmt.Println("IPv4 转 IPv6 过渡地址:", ipv6Addr)

	// 从 IPv6 过渡地址提取 IPv4 地址
	ipv4FromIPv6 := ipv6Addr.To4()
	if ipv4FromIPv6 == nil {
		fmt.Println("IPv6 过渡地址提取 IPv4 失败")
		return
	}

	fmt.Println("IPv6 过渡地址提取 IPv4 地址:", ipv4FromIPv6)
	getPodIPsFromEnv()
}

func getPodIPsFromEnv() (localV4 net.IP, localV6 net.IP) {
	// Example: POD_IPS=10.42.234.84,2001:cafe:42:bd:5fcf:8c4f:3e4b:fb0b
	// Please pay attention to the env value if present or not!
	podIPs := os.Getenv("POD_IPS")
	ips := strings.Split(podIPs, ",")
	//ips = append(ips, "10.42.234.84", "2001:cafe:42:bd:5fcf:8c4f:3e4b:fb0b", "::ffff:192.0.2.1")
	ips = append(ips, "::ffff:192.0.2.1")
	fmt.Println(ips, len(ips))
	for _, ip := range ips {
		fmt.Println("=========11111======:", ip)
		if ipObj := net.ParseIP(ip); ipObj != nil {
			fmt.Println("to v6: ", ipObj.To16())
			fmt.Println("to v4: ", ipObj.To4())
			if v4 := ipObj.To4(); v4 != nil && localV4 == nil {
				localV4 = v4
				fmt.Printf("Init pcap trace with localIPv4Addr from env POD_IPS: %v\n", localV4)
			} else if v6 := ipObj.To16(); v6 != nil {
				fmt.Println("==", v6.To4())
				// Avoid the v4v6 IP address
				if v4 := v6.To4(); v4 != nil && localV4 == nil {
					localV4 = v4
					fmt.Printf("Init pcap trace with localIPv4Addr from env POD_IPS: %v, v4v6 format: %v\n", localV4, v6)
				} else if localV6 == nil {
					localV6 = v6
					fmt.Printf("Init pcap trace with localIPv6Addr from env POD_IPS: %v\n", localV6)
				}
			} else {
				fmt.Println("end:", ipObj)
			}
		} else {
			fmt.Println("lhb endpoint", ip, ipObj)
		}
		//if localV4 != nil && localV6 != nil {
		//	break
		//}
	}
	if localV6 == nil { // v4v6 transition address
		localV6 = localV4.To16()
	}
	return
}
