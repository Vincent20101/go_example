package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	url := "http://localhost:8000"
	payload := []byte("your request body here")
	reader := bytes.NewReader(payload)

	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		fmt.Println("创建请求时出错:", err)
		return
	}

	req.Header.Set("lhb", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("3gpp-Sbi-Target-ApiRoot", "http://<HOSTPORT>")

	// 在发送之前打印或记录请求以检查头部
	fmt.Printf("%+v\n", req)

	// 发送请求
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求时出错:", err)
		return
	}
	defer resp.Body.Close()

	// 根据需要处理响应
	// ...
}
