package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(10 % 4)
	fmt.Println(14 % 4)
	return
	u, err := url.Parse("http://www.baidu.com:8080/test")
	fmt.Println(u.Host, u.Path, err)
	addr, err := net.ResolveIPAddr("ip4", "www.baidu.com")
	fmt.Println(addr, err)
	ip, err := net.LookupIP("www.baidu.com")
	fmt.Println(ip, err)
	port, p, err := net.SplitHostPort("smfsm.com:80")
	fmt.Println(port, p, err)
	v4, v6 := LookUpIPAddrFromFqdn("www.baidu.com")
	fmt.Println(v4, v6)
	fmt.Println(net.ParseIP("127.0.0.1:8888"))
	var sli []byte
	a := &sli
	fmt.Println(a)
	fmt.Println(net.ParseIP("::1").String())
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())
	var body []byte
	fmt.Println(len(body))
	parse, err := time.Parse(time.RFC3339, "2023-04-23T00:00:00Z")
	fmt.Println(parse, err)
	fmt.Println(time.Now().Before(parse))

	fmt.Println(ParseHost("172.29.95.4:8081"))
	fmt.Println(net.LookupIP("www.baidu.com"))
	//_ = make([]byte, -100)
	fmt.Println(-1 < 0)
}

func ParseHost(hostPort string) (ipv4, ipv6 net.IP, port uint16) {
	if hostStr, portStr, err := net.SplitHostPort(hostPort); err == nil {
		v4, v6 := FqdnToIP(hostStr)
		if v4 != nil {
			ipv4 = v4
		} else if v6 != nil {
			ipv6 = v6
		} else {
			ipv4 = net.ParseIP("127.0.0.1")
		}
		if p, err := strconv.Atoi(portStr); err == nil {
			port = uint16(p)
		}
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
		} else if ipv6 = addr.To16(); ipv6 != nil {
			v6 = ipv6
		}
	}
	return
}

func RemoteIp(req *http.Request) string {
	remoteIp := req.RemoteAddr

	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteIp = ip
	} else if ips := req.Header.Get("X-Forwarded-For"); ips != "" {
		for _, ip = range strings.Split(ips, ",") {
			if net.ParseIP(ip) != nil {
				remoteIp = ip
				break
			}
		}
	} else {
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}

	if remoteIp == "::1" || remoteIp == "" {
		remoteIp = net.ParseIP("127.0.0.1").String()
	}
	return remoteIp
}

func LookUpIPAddrFromFqdn(fqdn string) (v4, v6 string) {
	fmt.Println("====LookUpIPAddrFromFqdn")
	// Get the ip address from fqdn.
	addrs, err := net.LookupIP(fqdn)
	if err != nil {
		fmt.Printf("can't find ip address from fqdn:%s, err:%v", fqdn, err)
		return
	}
	var ipv4, ipv6 net.IP
	for _, addr := range addrs {
		if len(v4) != 0 && len(v6) != 0 {
			break
		}
		if ipv4 = addr.To4(); ipv4 != nil {
			fmt.Printf("local IPv4: %s", ipv4)
			v4 = ipv4.String()
		} else if ipv6 = addr.To16(); ipv6 != nil {
			fmt.Printf("local IPv6: %s", ipv6)
			v6 = ipv6.String()
		}
	}

	if len(v4) == 0 && len(v6) == 0 {
		fmt.Printf("no ip address found with this fqdn: %s", fqdn)
	}

	return
}
