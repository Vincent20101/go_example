package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	str := "1030::C9B4:FF12:48AA:1A2B"
	ip := net.ParseIP(str)
	fmt.Println([]byte(ip))
	fmt.Println(ip[8:].String())
	toString := hex.EncodeToString(ip[8:])
	fmt.Println(toString)
	decodeString, _ := hex.DecodeString(toString)
	fmt.Println(string(decodeString))
	//fmt.Println(IP6toInt(ip))

	//prefix := "2001:DB8::/64"
	//slice := strings.Split(prefix, "/")
	//netIP := net.IPNet{
	//	IP:   net.ParseIP(slice[0]),
	//	Mask: net.IPv4Mask(),
	//}
	//fmt.Println(netIP)
	var cidr net.IP
	var ipNet *net.IPNet
	var error error
	cidr, ipNet, error = net.ParseCIDR("5:6:7:8::/64")
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(cidr)
	fmt.Println(ipNet.IP, len(ipNet.IP), ipNet.Mask, len(ipNet.Mask), ipNet.Mask[:])
	fmt.Println(net.CIDRMask(12, 32).Size())

}

// ############## IPV6 ######################
func NetIpv6ToInt(ip net.IP) (*big.Int, error) {
	if ip == nil {
		return nil, errors.New("invalid ipv6")
	}
	IPv6Int := big.NewInt(0)
	IPv6Int.SetBytes(ip)
	return IPv6Int, nil
}

//NumToIPv6 converts a big integer represented by a string into an IPv6 address string
func NumToIPv6(numasstr string) (string, error) {
	bi, ok := new(big.Int).SetString(numasstr, 10)
	if !ok {
		return "", errors.New("fail to convert string to big.Int")
	}

	b255 := new(big.Int).SetBytes([]byte{255})
	var buf = make([]byte, 2)
	p := make([]string, 8)
	j := 0
	var i uint
	tmpint := new(big.Int)
	for i = 0; i < 16; i += 2 {
		tmpint.Rsh(bi, 120-i*8).And(tmpint, b255)
		bytes := tmpint.Bytes()
		if len(bytes) > 0 {
			buf[0] = bytes[0]
		} else {
			buf[0] = 0
		}
		tmpint.Rsh(bi, 120-(i+1)*8).And(tmpint, b255)
		bytes = tmpint.Bytes()
		if len(bytes) > 0 {
			buf[1] = bytes[0]
		} else {
			buf[1] = 0
		}
		p[j] = hex.EncodeToString(buf)
		j++
	}

	return strings.Join(p, ":"), nil
}

func convertAndPrint(numasstr string) {
	ipstr, _ := NumToIPv6(numasstr)
	fmt.Printf("%s %v\n", ipstr, numasstr)
}

// ############## IPV4 start ######################

func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

// ############## IPV4 end ######################

// 0: invalid ip
// 4: IPv4
// 6: IPv6
// 检查IP 用于 net.ParseIP(ipv4)
func ParseIP(s string) (net.IP, int) {
	ip := net.ParseIP(s)
	if ip == nil {
		return nil, 0
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return ip, 4
		case ':':
			return ip, 6
		}
	}
	return nil, 0
}

// IPv4ByLong
// 将 uint32 长整型转换成IPV4 地址
// converts a uint32 represented by a string into an ipv4 address string
// 168427779 => "10.10.1.2"
func IPv4ByLong(ipv4long string) (string, error) {

	ipv4Int, err := strconv.ParseInt(ipv4long, 10, 64)
	if err != nil {
		return "", errors.New(fmt.Sprintf("fail to convert string to Int64 :%s", err.Error()))
	}

	ipv4 := uint32(ipv4Int)

	return fmt.Sprintf("%d.%d.%d.%d", ipv4>>24, ipv4<<8>>24, ipv4<<16>>24, ipv4<<24>>24), nil
}

// IPv6ByLong
// 将 big.Int 长整型转换成IPV6 地址
// converts a big integer represented by a string into an IPv6 address string
// 53174336768213711679990085974688268287=> "2801:0137:0000:0000:0000:ffff:ffff:ffff"
func IPv6ByLong(ipv6long string) (string, error) {
	bi, ok := new(big.Int).SetString(ipv6long, 10)
	if !ok {
		return "", errors.New("fail to convert string to big.Int")
	}

	b255 := new(big.Int).SetBytes([]byte{255})
	var buf = make([]byte, 2)
	p := make([]string, 8)
	j := 0
	var i uint
	tmpint := new(big.Int)
	for i = 0; i < 16; i += 2 {
		tmpint.Rsh(bi, 120-i*8).And(tmpint, b255)
		bytes := tmpint.Bytes()
		if len(bytes) > 0 {
			buf[0] = bytes[0]
		} else {
			buf[0] = 0
		}
		tmpint.Rsh(bi, 120-(i+1)*8).And(tmpint, b255)
		bytes = tmpint.Bytes()
		if len(bytes) > 0 {
			buf[1] = bytes[0]
		} else {
			buf[1] = 0
		}
		p[j] = hex.EncodeToString(buf)
		j++
	}

	return strings.Join(p, ":"), nil
}

// 将IPV4 转换成 uint32 长整型
func IPv4ToInt(ipv4 string) (ip uint32) {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return 0
	}
	ips := reg.FindStringSubmatch(ipv4)
	if ips == nil {
		return 0
	}

	//上面正则做了判断，这里就不报错了
	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])

	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return 0
	}

	ip += uint32(ip1 * 0x1000000) // 左移24位
	ip += uint32(ip2 * 0x10000)   // 左移16位
	ip += uint32(ip3 * 0x100)     // 左移8位
	ip += uint32(ip4)             // 左移0位

	return ip
}

// 将IPV6 转换成 big.Int 长整型
func IPv6ToInt(ipv6 string) (*big.Int, error) {
	ip := net.ParseIP(ipv6)
	return NetIpv6ToInt(ip)
}

// 将net.IP 类型 转换成 big.Int 长整型
func DupNetIpv6ToInt(ip net.IP) (*big.Int, error) {
	if ip == nil {
		return nil, errors.New("invalid ipv6")
	}
	IPv6Int := big.NewInt(0)
	IPv6Int.SetBytes(ip)
	return IPv6Int, nil
}
