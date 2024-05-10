package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

func main() {
	//time.AfterFunc(0*time.Millisecond, func() {
	//	fmt.Println("test")
	//})
	//time.Sleep(1e9 * time.Second)

	var nearest time.Time
	endTime := time.Now()
	//endTime := time.Now().Add(3000)
	b := nearest.After(endTime)
	c := nearest.Before(endTime)
	fmt.Println(b, c)

	fmt.Println(os.Getwd())
	fmt.Println(path.Join("/", "configs", "template"))

	fmt.Println([]byte(`tcpdump -ddd "(port 8805 or port 2123) and udp or (ip[6:2] & 0x3fff != 0 and udp) or (ip6[6] == 44 and udp)"`))
}
