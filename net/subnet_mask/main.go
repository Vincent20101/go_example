package main

import (
	"fmt"
	"net"
)

// findMinSubnetMask finds the minimum subnet mask that can cover both IPs.
func findMinSubnetMask(ip1, ip2 string) string {
	ip1Binary := net.ParseIP(ip1).To4()
	ip2Binary := net.ParseIP(ip2).To4()

	if ip1Binary == nil || ip2Binary == nil {
		return "Invalid IP address"
	}

	// Convert IPs to binary strings
	ip1BinaryStr := fmt.Sprintf("%08b", ip1Binary)
	ip2BinaryStr := fmt.Sprintf("%08b", ip2Binary)

	// Find the common prefix length
	commonPrefixLen := 0
	for i := 0; i < 8; i++ {
		if ip1BinaryStr[i] == ip2BinaryStr[i] {
			commonPrefixLen++
		} else {
			break
		}
	}

	// Calculate the subnet mask
	subnetMask := net.IPv4Mask(0xFFFFFFFF << (32 - commonPrefixLen))

	return subnetMask.String()
}

func main() {
	ip1 := "192.168.11.10"
	ip2 := "192.168.11.120"
	fmt.Printf("Minimum subnet mask for %s and %s is: %s\n", ip1, ip2, findMinSubnetMask(ip1, ip2))

	ip1 = "192.168.11.10"
	ip2 = "192.168.12.120"
	fmt.Printf("Minimum subnet mask for %s and %s is: %s\n", ip1, ip2, findMinSubnetMask(ip1, ip2))
}
