package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	sendPing()
	//time.Sleep(time.Hour)
}

func sendPing() {
	// 建立 TCP 连接
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	//defer cancel()
	//d := net.Dialer{}
	//conn, err := d.DialContext(ctx, "tcp", "localhost:8082")
	conn, err := net.DialTimeout("tcp", "localhost:8080", time.Second*3)
	//conn, err := net.Dial("tcp", "localhost:8085")
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer conn.Close()

	// 创建 http2.ClientConn
	//cc := http2.ClientConn{}
	transport := &http2.Transport{
		//PingTimeout: 2,
		AllowHTTP: true,
		//DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
		//	return net.Dial(network, addr)
		//},
		DialTLSContext: func(_ context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
			//set connection timeout as 2 seconds
			return net.DialTimeout(network, addr, time.Second*time.Duration(2))
		},
	}
	newClientConn, err := transport.NewClientConn(conn)
	//cc.Ping()
	//clientConn, err := http2.NewClientConn(conn, &http2.ClientConnOpts{})
	if err != nil {
		fmt.Println("Failed to create client connection:", err)
		return
	}
	defer newClientConn.Close()

	// 发送 Ping 帧
	err = newClientConn.Ping(context.Background())
	if err != nil {
		fmt.Println("Failed to send Ping frame:", err)
		return
	}

	//fmt.Println("Ping frame sent with ID:", pingID)
	//
	//// 等待 Ping 响应
	//select {
	//case <-pingID.Ack:
	//	fmt.Println("Ping acknowledged:", string(pingData))
	//	break
	//case <-clientConn.GoAway():
	//	fmt.Println("Connection closed before Ping response")
	//	break
	//}
	return
}
