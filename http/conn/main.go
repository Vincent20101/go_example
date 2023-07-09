package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type connectionCounter struct {
	count int64
}

func (cc *connectionCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&cc.count, 1)
	defer atomic.AddInt64(&cc.count, -1)

	// 处理请求的逻辑
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	counter := &connectionCounter{}
	handler := counter

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Server started on port 8080")

	// 获取当前连接数
	fmt.Println("Current connections:", atomic.LoadInt64(&counter.count))

	// 阻塞主线程
	select {}
}
