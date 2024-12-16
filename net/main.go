package main

import (
	"fmt"
	"math/big"
	"net"
	"time"
	"unsafe"
)

func main() {
	fmt.Println(net.ParseIP("12::/64"))

	fmt.Println(unsafe.Sizeof("aaa"))

	var t time.Time
	fmt.Println(t)
	fmt.Println(t.IsZero())

	var s *string
	s1 := "test"
	s = &s1
	fmt.Println(*s)

	ss := "12::/64"
	byte := []byte(ss)
	fmt.Println(byte)
	fmt.Println(byte[len(byte)-1])
	fmt.Println(string(byte[len(byte)-1]))

	fmt.Println(generateIPRanges("19.40.2.11", "19.40.2.30"))
}

func generateIPRanges(start, end string) ([]string, bool) {
	if start == "" || end == "" {
		return nil, true
	}

	startIP := net.ParseIP(start)
	endIP := net.ParseIP(end)
	diff, isV4 := IpDiff(startIP, endIP)
	ipRanges := make([]string, 0)
	for i := 0; i < diff+1; i++ {
		if isV4 {
			hostIpv4 := startIP.To4()
			hostIpModified := make(net.IP, len(hostIpv4))
			copy(hostIpModified, hostIpv4)
			hostIpModified[3] = hostIpModified[3] + byte(i)
			ipRanges = append(ipRanges, hostIpModified.String())
		} else {
			hostIpv6 := startIP.To16()
			hostIpModified := make(net.IP, len(hostIpv6))
			copy(hostIpModified, hostIpv6)
			hostIpModified[15] = hostIpModified[15] + byte(i)
			ipRanges = append(ipRanges, hostIpModified.String())
		}
	}
	return ipRanges, isV4
}

func IpDiff(ip1, ip2 net.IP) (int, bool) {
	isV4 := true
	if ip1.To4() != nil && ip2.To4() != nil {
		ip1Int := IpToInt(ip1)
		ip2Int := IpToInt(ip2)
		return ip2Int - ip1Int, isV4
	} else {
		return Ip6Diff(ip1, ip2), !isV4
	}
}

func IpToInt(ip net.IP) int {
	ip = ip.To4()

	var ipInt int
	for _, b := range ip {
		ipInt = ipInt<<8 + int(b)
	}

	return ipInt
}

func Ip6Diff(ip1, ip2 net.IP) int {
	ip1Int := Ip6ToInt(ip1)
	ip2Int := Ip6ToInt(ip2)

	return int(ip2Int.Sub(ip2Int, ip1Int).Int64())
}

func Ip6ToInt(ip net.IP) *big.Int {
	ipInt := new(big.Int)
	ipInt.SetBytes(ip.To16())

	return ipInt
}
