package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
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

type Student struct {
	id   string
	name string
}

func httpGetWithContext(ctx context.Context) {
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	//ctx = context.Background()
	//defer cancel()
	stu := Student{
		id:   "2ed4tg5fe35fgty3yy6uh",
		name: "amber",
	}
	stus, err := json.Marshal(&stu)
	reader := bytes.NewReader(stus)
	req, err := http.NewRequest("GET", "http://localhost:8000", reader)
	dump, errs := httputil.DumpRequestOut(req, true)
	fmt.Printf("\n%s, DumpRequestOut error:%v\n", string(dump), errs)
	fmt.Println("ssss==req==", req)
	fmt.Println("ssss====", req.Context())
	if err != nil {
		log.Fatal("无法生成请求：", err)
	}
	req.Header.Add("User-Agent", "cli")
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
	dump, errs = httputil.DumpResponse(resp, true)
	fmt.Printf("\n%s, DumpResponse error:%v\n", string(dump), errs)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("无法读取返回内容：", err)
	}
	fmt.Println("结果：", string(data))
}
