package main

import (
	"fmt"
	"net"
)

// CheckIPAddress checks if the given FQDN is an IP address and returns its type.
func CheckIPAddress(fqdn string) (string, bool) {
	ip := net.ParseIP(fqdn)
	if ip == nil {
		return "", false
	}
	if ip.To4() != nil {
		return "IPv4", true
	}
	return "IPv6", true
}

func main() {
	tests := []string{
		"ab12:2211:0:1::55:8000",
		"192.168.1.1",                    // IPv4
		"2001:0db8:85a3::8a2e:0370:7334", // IPv6
		"example.com",                    // Not an IP
	}

	for _, test := range tests {
		ipType, isIP := CheckIPAddress(test)
		if isIP {
			fmt.Printf("%s is an %s address\n", test, ipType)
		} else {
			fmt.Printf("%s is not an IP address\n", test)
		}
	}
}
