package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 创建一个 HTTP 处理函数
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 写入响应头
		w.WriteHeader(http.StatusBadGateway)
		// 再次写入响应头，导致冗余调用
		w.WriteHeader(http.StatusBadGateway)
	})

	// 注册处理函数并启动 HTTP 服务器
	http.Handle("/", handler)
	fmt.Println("Server started on port 8089")
	http.ListenAndServe(":8089", nil)
}
