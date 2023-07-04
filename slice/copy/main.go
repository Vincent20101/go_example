package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	type HasLabels []string
	s := HasLabels{"s6", "s4"}
	fmt.Println(s)
	var lv4 net.IP
	LocalIPv4 := net.IPv4(127, 0, 0, 1)
	//copy(lv4, LocalIPv4) // unused
	//lv4 = append(lv4, LocalIPv4...)
	lv4 = LocalIPv4
	fmt.Printf("%p, %p\n", &lv4, &LocalIPv4)

	fmt.Println(strings.Split("1 + 2 + 3 - 4 + 5 - 6", " "))

	var a, b []int
	b = append(b, 1, 2, 3, 4, 5)
	a = append(a, b...)
	b = nil
	fmt.Println(a)
}
