package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	client := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLSContext: func(_ context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
				//set connection timeout as 2 seconds
				return net.DialTimeout(network, addr, time.Second*time.Duration(2))
			},
		},
	}

	request, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println("Response Status:", response.Status)

	// 读取响应内容
	buf := make([]byte, 1024)
	n, err := response.Body.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response Body:", string(buf[:n]))
}
