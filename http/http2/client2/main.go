package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptrace"

	"golang.org/x/net/http2"
)

func main() {
	client := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	req, err := http.NewRequest("HEAD", "http://localhost:8080/", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Ping", "HelloServer")
	// 定义ClientTrace钩子，在获取连接时触发
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			localAddr := connInfo.Conn.LocalAddr().String()
			remoteAddr := connInfo.Conn.RemoteAddr().String()
			fmt.Println("===================")
			fmt.Printf("本地地址 (源IP): %s\n", localAddr)
			fmt.Printf("远程地址 (目标IP): %s\n", remoteAddr)
		},
	}
	req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
	fmt.Println("Response:")
	fmt.Println("--------------------")
	fmt.Println(resp.Body)
	buf := make([]byte, 1024)
	n, err := resp.Body.Read(buf)
	if err != nil {
		fmt.Println("lhb")
		log.Fatal(err)
	}

	fmt.Println("Response Body:", string(buf[:n]))
}
