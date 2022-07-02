//客户端发送封包
package main

import (
	"fmt"
	"net"
	"os"

	"github.com/vincent20101/go-example/nian_bao/protocol"
)

func sender(conn net.Conn) {
	for i := 0; i < 200; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write(protocol.Packet([]byte(words)))
	}
	fmt.Println()
	fmt.Println("send over")
}

func main() {
	server := "127.0.0.1:9988"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	fmt.Println("connect success")
	go sender(conn)
	fmt.Println("df end")
	//for {
	//	//time.Sleep(1 * 1e9)
	//	time.Sleep(1 * time.Second)
	//}
	select {}
}
