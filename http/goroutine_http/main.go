package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func handleJobSave(resp http.ResponseWriter, req *http.Request) {
	go func() {
		for {
			fmt.Println("goroutine background")
		}
	}()
	return
}
func main() {

	// 配置路由
	mux := http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)

	// 启动TCP监听
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		return
	}

	// 创建一个HTTP服务
	httpServer := &http.Server{
		ReadTimeout:  time.Duration(3) * time.Millisecond,
		WriteTimeout: time.Duration(3) * time.Millisecond,
		Handler:      mux,
	}

	err = httpServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
		return
	}
}
