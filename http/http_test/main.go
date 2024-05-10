package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//func main() {
//	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
//		_ = r.ParseForm()
//		fmt.Println("path:", r.URL.Path)
//		a, _ := strconv.Atoi(r.Form["a"][0])
//		b, _ := strconv.Atoi(r.Form["b"][0])
//		w.Header().Set("Content-Type", "application/json")
//		jData, _ := json.Marshal(map[string]int{
//			"data": a + b,
//		})
//		w.Write(jData)
//	})
//	log.Fatal(http.ListenAndServe(":8010", nil))
//}

func main() {
	// 启动 HTTP 服务器
	go startServer()

	time.Sleep(5 * time.Second)
	// 创建 HTTP 客户端并发起请求
	resp, err := http.Get("http://localhost:8089")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response from server:", string(body))
}

func startServer() {
	// 创建一个 HTTP 处理函数
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 写入响应头
		w.WriteHeader(http.StatusOK)
		// 再次写入响应头，导致冗余调用
		w.WriteHeader(http.StatusOK)
	})

	// 注册处理函数并启动 HTTP 服务器
	http.Handle("/", handler)
	fmt.Println("Server started on port 8089")
	http.ListenAndServe(":8089", nil)
}
