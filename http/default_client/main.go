package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	//go httpDirectGet()
	go httpGetWithContext(ctx)
	<-ctx.Done()
	//time.Sleep(1000e9)
}

func httpDirectGet() {
	resp, err := http.Get("http://localhost:8083/moreJSON")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func httpGetWithContext(ctx context.Context) {
	ctx, _ = context.WithTimeout(ctx, 11*time.Second)
	//defer cancel()
	req, err := http.NewRequest("get", "http://localhost:8000", nil)
	if err != nil {
		log.Fatal("无法生成请求：", err)
	}
	ctx = context.WithValue(ctx, "p", "q")
	//req = req.WithContext(ctx)
	req = req.WithContext(ctx)
	fmt.Println("第一：", req.Context().Value("p"))
	//deadline, ok := req.Context().Deadline()
	//fmt.Println(deadline, "===", ok)
	fmt.Println("ctx 第二：", ctx)
	fmt.Println("req的值：", req)
	fmt.Println("req的值：", req.Context())
	resp, err := http.DefaultClient.Do(req)
	fmt.Println("resp====:", resp)
	fmt.Println("err=======:", err)
	if err != nil {
		log.Println("无法发送请求：", err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("无法读取返回内容：", err)
	}
	fmt.Println("结果：", string(data))
}
