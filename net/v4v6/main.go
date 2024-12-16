package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// IPv4 地址
	//ipv4Addr := net.ParseIP("fd01:0:0:1::5fb")
	ipv4Addr := net.ParseIP("::ffff:192.0.2.1")
	fmt.Println(ipv4Addr.String())
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

	fmt.Println("===di==")
	fmt.Println(ParseHost("www.baidu.com"))
	fmt.Println(ParseHost("192.168.10.1:8080"))
	fmt.Println(ParseHost("[fd01:0:0:1::5fb]:8080"))
	fmt.Println(strconv.Atoi(""))
}

func ParseHost(hostPort string) (ipv4, ipv6 net.IP, port uint16) {
	hostStr, portStr, err := net.SplitHostPort(hostPort)
	fmt.Println(hostStr, portStr, err)
	if err != nil {
		if strings.Contains(err.Error(), "missing port in address") && hostStr == "" {
			fmt.Println("lhb test failed:", err, "hostStr:", hostStr)
		}
		return
	}
	if p, err := strconv.Atoi(portStr); err == nil {
		port = uint16(p)
	}
	v4, v6 := FqdnToIP(hostStr)
	if v4 != nil {
		ipv4 = v4
	}
	if v6 != nil {
		ipv6 = v6
	}
	if v4 == nil && v6 == nil {
		ipv4 = net.ParseIP("127.0.0.1")
	}
	return
}

func FqdnToIP(fqdn string) (v4, v6 net.IP) {
	addrs, err := net.LookupIP(fqdn)
	if err != nil {
		return
	}
	var ipv4, ipv6 net.IP
	for _, addr := range addrs {
		if len(v4) != 0 && len(v6) != 0 {
			break
		}
		if ipv4 = addr.To4(); ipv4 != nil {
			v4 = ipv4
		}
		if ipv6 = addr.To16(); ipv6 != nil {
			v6 = ipv6
		}
	}
	return
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
			fmt.Println("lhb: ", ipObj.String())
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
