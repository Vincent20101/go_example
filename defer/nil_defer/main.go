package main

import (
	"fmt"
	"net"
)

type student struct {
	name string
	age  int
}

func main() {
	var s *student
	defer func() {
		s.name = "new name"
		fmt.Println(s.name)
	}()

	// panic
	// defer func(s string){
	// 	fmt.Println(s)
	// }(s.name)

	s = new(student)
	s.name = "test"
	s.age = 10
	fmt.Println(s)

	var v4 net.IP
	fmt.Println(v4)
	fmt.Println(v4.Equal([]byte{}))

	var v6 net.IP
	v6 = net.IP{}
	fmt.Println(v6)

	var conn *net.UDPConn
	// fmt.Println(conn.Read([]byte{}))
	rAddr := new(net.UDPAddr)
	conn.WriteToUDP([]byte{}, rAddr)
}
