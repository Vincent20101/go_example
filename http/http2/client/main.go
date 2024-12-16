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
			//DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			//	return net.Dial(network, addr)
			//},
		},
	}

	request, err := http.NewRequest("GET", "http://172.0.3.64", nil)
	//request.Host = "localhost:8082"
	//request.URL.Host = "localhost:8080"
	fmt.Println("request.Host:", request.Host)
	fmt.Println("request.URL.Host:", request.URL.Host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("request context: %+v\n", request)
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
		fmt.Println("lhb")
		log.Fatal(err)
	}

	fmt.Println("Response Body:", string(buf[:n]))
}
