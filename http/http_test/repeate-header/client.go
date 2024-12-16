package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 创建 HTTP 客户端并发起请求
	resp, err := http.Get("http://localhost:8089")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response from server:", string(body))

}
